package datamove

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-dialect/worker"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"io"
	"os"
	"strings"
)

type DataSourceTxt struct {
	*DataSourceFile
	RowSeparator      byte              `json:"rowSeparator"`      // 行 分隔符 默认 `\n`
	ColSeparator      string            `json:"colSeparator"`      // 列 分隔符 默认 `,`
	ReplaceSeparators map[string]string `json:"replaceSeparators"` // 替换字符，如将：`\n` 替换为 `|:-n-:|`，`,` 替换为 `|:-，-:|`，写入时候 将 key 替换为 value，读取时候将 value 替换为 key
	ShouldTrimSpace   bool              `json:"shouldTrimSpace"`   // 是否需要去除空白字符
	headerRead        bool
	headerWrite       bool
	readColumnList    []*dialect.ColumnModel
	readColumnLength  int
	ColumnNameMapping map[string]string `json:"columnNameMapping"`
}

func (this_ *DataSourceTxt) Stop(progress *DateMoveProgress) {
	this_.CloseReadFile()
	this_.CloseWriteFile()
}

func (this_ *DataSourceTxt) ReadStart(progress *DateMoveProgress) (err error) {
	_, err = this_.GetReadFile()
	return
}

func (this_ *DataSourceTxt) ColsToData(progress *DateMoveProgress, cols []string, columnList []*dialect.ColumnModel) (data map[string]interface{}, err error) {
	data = map[string]interface{}{}
	for index, col := range cols {
		v := col
		if this_.ReplaceSeparators != nil {
			for rK, rV := range this_.ReplaceSeparators {
				v = strings.ReplaceAll(v, rV, rK)
			}
		}
		if this_.ShouldTrimSpace {
			v = strings.TrimSpace(v)
		}
		name := columnList[index].ColumnName
		if this_.ColumnNameMapping != nil && this_.ColumnNameMapping[name] != "" {
			name = this_.ColumnNameMapping[name]
		}
		data[name] = v
	}
	return
}

func (this_ *DataSourceTxt) DataToCols(progress *DateMoveProgress, data map[string]interface{}, columnList []*dialect.ColumnModel) (cols []string, err error) {
	for _, column := range columnList {
		v := util.GetStringValue(data[column.ColumnName])
		if this_.ReplaceSeparators != nil {
			for rK, rV := range this_.ReplaceSeparators {
				v = strings.ReplaceAll(v, rK, rV)
			}
		}
		if this_.ShouldTrimSpace {
			v = strings.TrimSpace(v)
		}
		cols = append(cols, v)
	}
	return
}

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

func (this_ *DataSourceTxt) ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	rowSeparator := this_.GetRowSeparator()
	if i := bytes.IndexByte(data, rowSeparator); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, dropCR(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

func (this_ *DataSourceTxt) ReadLineCount() (lineCount int64, err error) {
	file, err := os.Open(this_.FilePath)
	if err != nil {
		err = errors.New("open file [" + this_.FilePath + "] error:" + err.Error())
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(this_.ScanLines)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		lineCount++
	}
	util.Logger.Info("txt data source read line count", zap.Any("lineCount", lineCount))
	return
}

func (this_ *DataSourceTxt) Read(progress *DateMoveProgress, dataChan chan *Data) (err error) {

	file, err := this_.GetReadFile()
	if err != nil {
		return
	}
	buf := bufio.NewReader(file)
	var line string
	rowSeparator := this_.GetRowSeparator()
	colSeparator := this_.GetColSeparator()

	var lastData = &Data{
		DataType: DataTypeData,
	}
	lineCount, err := this_.ReadLineCount()
	if err != nil {
		return
	}
	progress.Total = lineCount - 1 // 第一行为头信息
	pageSize := progress.BatchNumber
	for {
		if progress.ShouldStop() {
			return
		}
		line, err = buf.ReadString(rowSeparator)
		line = strings.TrimSuffix(line, string(rowSeparator))
		if line != "" {
			cols := strings.Split(line, colSeparator)
			if !this_.headerRead {
				for _, col := range cols {
					this_.readColumnList = append(this_.readColumnList, &dialect.ColumnModel{
						ColumnName: col,
					})
				}
				this_.readColumnLength = len(this_.readColumnList)
				this_.headerRead = true
				continue
			}
			if len(cols) != this_.readColumnLength {
				err = errors.New("line [" + line + "] can not to column names [" + strings.Join(worker.GetColumnNames(this_.readColumnList), ",") + "]")
				return
			}

			rowData, e := this_.ColsToData(progress, cols, this_.readColumnList)
			if e != nil {
				progress.Read.Errors = append(progress.Read.Errors, e.Error())
				progress.Read.AddError(1)
				progress.callback(progress)
				if !progress.ErrorContinue {
					err = e
					return
				}
			} else {
				lastData.ColumnList = this_.readColumnList
				lastData.DataList = append(lastData.DataList, rowData)
				lastData.Total++
				progress.Read.AddSuccess(1)
				if lastData.Total >= pageSize {
					progress.callback(progress)
					dataChan <- lastData
					lastData = &Data{
						DataType: DataTypeData,
					}
				}
			}
		}
		if err != nil {
			if err == io.EOF { //读取结束，会报EOF
				err = nil
			}
			break
		}
	}
	if lastData.Total > 0 {
		progress.callback(progress)
		dataChan <- lastData
	}
	return
}

func (this_ *DataSourceTxt) ReadEnd(progress *DateMoveProgress) (err error) {
	this_.CloseReadFile()
	return
}

func (this_ *DataSourceTxt) WriteStart(progress *DateMoveProgress) (err error) {
	_, err = this_.GetWriteFile()
	return
}

func (this_ *DataSourceTxt) GetRowSeparator() byte {
	if this_.RowSeparator == 0 {
		return '\n'
	}
	return this_.RowSeparator
}

func (this_ *DataSourceTxt) GetColSeparator() string {
	if this_.ColSeparator == "" {
		return ","
	}
	return this_.ColSeparator
}

func (this_ *DataSourceTxt) Write(progress *DateMoveProgress, data *Data) (err error) {

	if err = ValidateDataType(data.DataType); err != nil {
		return
	}
	rowSeparator := this_.GetRowSeparator()
	colSeparator := this_.GetColSeparator()

	file, err := this_.GetWriteFile()
	if err != nil {
		return
	}
	if !this_.headerWrite {
		if len(data.ColumnList) == 0 {
			err = errors.New("字段信息不存在，无法写入头部信息")
			return
		}
		var names []string
		for _, c := range data.ColumnList {
			name := c.ColumnName
			if this_.ColumnNameMapping != nil && this_.ColumnNameMapping[name] != "" {
				name = this_.ColumnNameMapping[name]
			}
			names = append(names, name)
		}

		line := strings.Join(names, colSeparator)

		_, _ = file.WriteString(line + string(rowSeparator))
		this_.headerWrite = true
	}

	switch data.DataType {
	case DataTypeData:
		data.Total = int64(len(data.DataList))
		if data.Total > 0 {
			for _, one := range data.DataList {
				cols, e := this_.DataToCols(progress, one, data.ColumnList)
				if e != nil {
					progress.Write.Errors = append(progress.Write.Errors, e.Error())
					progress.Write.AddError(1)
					progress.callback(progress)
					if !progress.ErrorContinue {
						err = e
						return
					}
				} else {
					line := strings.Join(cols, colSeparator)
					_, _ = file.WriteString(line + string(rowSeparator))
					progress.Write.AddSuccess(1)
				}
			}
		}
		break
	default:
		err = errors.New(fmt.Sprintf("当前数据类型[%d]，不支持写入", data.DataType))
		return
	}
	progress.callback(progress)
	return
}

func (this_ *DataSourceTxt) WriteEnd(progress *DateMoveProgress) (err error) {
	this_.CloseWriteFile()
	return
}
