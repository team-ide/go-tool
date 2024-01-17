package datamove

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"io"
	"os"
	"strings"
)

func NewDataSourceTxt() *DataSourceTxt {
	return &DataSourceTxt{
		DataSourceBase: &DataSourceBase{},
		DataSourceFile: &DataSourceFile{},
	}
}

type DataSourceTxt struct {
	*DataSourceBase
	*DataSourceFile
	headerRead   bool
	headerWrite  bool
	ColSeparator string `json:"colSeparator"` // 列 分隔符 默认 `,`
	ReplaceCol   string `json:"replaceCol"`   //
	ReplaceLine  string `json:"replaceLine"`  //
}

func (this_ *DataSourceTxt) GetColSeparator() string {
	if this_.ColSeparator == "" {
		return ","
	}
	return this_.ColSeparator
}

func (this_ *DataSourceTxt) StringsToValues(progress *Progress, cols []string) (res []interface{}, err error) {
	vSize := len(cols)
	for index, _ := range this_.ColumnList {
		var v string
		if vSize > index {
			v = cols[index]
		}
		if this_.ReplaceCol != "" {
			v = strings.ReplaceAll(v, this_.ReplaceCol, this_.GetColSeparator())
		}
		if this_.ReplaceLine != "" {
			v = strings.ReplaceAll(v, this_.ReplaceLine, "\n")
		}
		if this_.ShouldTrimSpace {
			v = strings.TrimSpace(v)
		}
		res = append(res, v)
	}
	return
}

func (this_ *DataSourceTxt) ValuesToStrings(progress *Progress, cols []interface{}) (res []string, err error) {

	vSize := len(cols)
	for index, c := range this_.ColumnList {
		if c.ColumnName == "" {
			continue
		}
		var v string
		if vSize > index {
			v = util.GetStringValue(cols[index])
		}
		if this_.ReplaceCol != "" {
			v = strings.ReplaceAll(v, this_.GetColSeparator(), this_.ReplaceCol)
		}
		if this_.ReplaceLine != "" {
			v = strings.ReplaceAll(v, "\n", this_.ReplaceLine)
		}
		if this_.ShouldTrimSpace {
			v = strings.TrimSpace(v)
		}
		res = append(res, v)
	}

	return
}

func (this_ *DataSourceTxt) ReadTitles(progress *Progress) (titles []string, err error) {
	file, err := os.Open(this_.GetFilePath())
	if err != nil {
		err = errors.New("open file [" + this_.GetFilePath() + "] error:" + err.Error())
		return
	}
	defer func() { _ = file.Close() }()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		cols := strings.Split(line, this_.GetColSeparator())
		for _, c := range cols {
			if this_.ShouldTrimSpace {
				c = strings.TrimSpace(c)
			}
			titles = append(titles, c)
		}
		break
	}
	util.Logger.Info("file data source read titles", zap.Any("titles", titles))
	return
}

func (this_ *DataSourceTxt) ReadStart(progress *Progress) (err error) {
	err = this_.DataSourceFile.ReadStart(progress)
	if err != nil {
		return
	}

	titles, err := this_.ReadTitles(progress)
	if err != nil {
		return
	}
	if len(this_.ColumnList) == 0 {
		for _, title := range titles {
			column := &Column{
				ColumnModel: &dialect.ColumnModel{},
			}
			column.ColumnName = title
			this_.ColumnList = append(this_.ColumnList, column)
		}
	}

	lineCount, err := this_.ReadLineCount()
	if err != nil {
		return
	}
	progress.DataTotal += lineCount - 1 // 第一行为头信息
	return
}

func (this_ *DataSourceTxt) Read(progress *Progress, dataChan chan *Data) (err error) {

	file, err := this_.GetReadFile()
	if err != nil {
		return
	}
	buf := bufio.NewReader(file)
	var line string
	colSeparator := this_.GetColSeparator()

	var lastData = &Data{
		DataType: DataTypeCols,
	}
	pageSize := progress.BatchNumber
	for {
		if progress.ShouldStop() {
			return
		}
		line, err = buf.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		if line != "" {
			cols := strings.Split(line, colSeparator)
			if !this_.headerRead {
				this_.headerRead = true
				continue
			}

			values, e := this_.StringsToValues(progress, cols)
			if e != nil {
				progress.ReadCount.AddError(1, e)
				if !progress.ErrorContinue {
					err = e
					return
				}
			} else {
				lastData.ColsList = append(lastData.ColsList, values)
				lastData.Total++
				progress.ReadCount.AddSuccess(1)
				if lastData.Total >= pageSize {
					dataChan <- lastData
					lastData = &Data{
						DataType: DataTypeCols,
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
		dataChan <- lastData
	}
	return
}

func (this_ *DataSourceTxt) Write(progress *Progress, data *Data) (err error) {

	if data.columnList != nil {
		this_.ColumnList = *data.columnList
	}

	colSeparator := this_.GetColSeparator()

	buf, err := this_.GetWriteBuf()
	if err != nil {
		return
	}
	if !this_.headerWrite {
		if len(this_.ColumnList) == 0 {
			err = errors.New("字段信息不存在，无法写入头部信息")
			return
		}
		line := strings.Join(this_.GetColumnNames(), colSeparator)
		_, _ = buf.WriteString(line)
		_ = buf.Flush()
		this_.headerWrite = true
	}

	switch data.DataType {
	case DataTypeCols:
		data.Total = int64(len(data.ColsList))
		if data.Total > 0 {
			for _, cols := range data.ColsList {
				stringValues, e := this_.ValuesToStrings(progress, cols)
				if e != nil {
					progress.WriteCount.AddError(1, e)
					if !progress.ErrorContinue {
						err = e
						return
					}
				} else {
					line := strings.Join(stringValues, colSeparator)
					_, _ = buf.WriteString("\n" + line)
					progress.WriteCount.AddSuccess(1)
				}
			}
		}
		break
	default:
		err = errors.New(fmt.Sprintf("当前数据类型[%d]，不支持写入", data.DataType))
		return
	}
	_ = buf.Flush()
	return
}
