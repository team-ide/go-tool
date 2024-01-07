package datamove

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

type DataSourceTxt struct {
	DataSourceFile
	headerRead  bool
	headerWrite bool
}

func (this_ *DataSourceTxt) ReadStart(progress *DateMoveProgress) (err error) {
	err = this_.DataSourceFile.ReadStart(progress)
	if err != nil {
		return
	}

	titles, err := this_.ReadTitles()
	if err != nil {
		return
	}
	if len(this_.ColumnList) == 0 {
		for _, title := range titles {
			column := &Column{}
			column.ColumnName = title
			this_.ColumnList = append(this_.ColumnList, column)
		}
	}

	lineCount, err := this_.ReadLineCount()
	if err != nil {
		return
	}
	progress.Total = lineCount - 1 // 第一行为头信息
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
		DataType: DataTypeCols,
	}
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
				this_.headerRead = true
				continue
			}

			values, e := this_.StringsToValues(progress, cols)
			if e != nil {
				progress.Read.Errors = append(progress.Read.Errors, e.Error())
				progress.Read.AddError(1)
				progress.callback(progress)
				if !progress.ErrorContinue {
					err = e
					return
				}
			} else {
				lastData.ColsList = append(lastData.ColsList, values)
				lastData.Total++
				progress.Read.AddSuccess(1)
				if lastData.Total >= pageSize {
					progress.callback(progress)
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
		progress.callback(progress)
		dataChan <- lastData
	}
	return
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
		if len(this_.ColumnList) == 0 {
			err = errors.New("字段信息不存在，无法写入头部信息")
			return
		}
		line := strings.Join(this_.GetColumnNames(), colSeparator)
		_, _ = file.WriteString(line)
		this_.headerWrite = true
	}

	switch data.DataType {
	case DataTypeCols:
		data.Total = int64(len(data.ColsList))
		if data.Total > 0 {
			for _, cols := range data.ColsList {
				stringValues, e := this_.ValuesToStrings(progress, cols)
				if e != nil {
					progress.Write.Errors = append(progress.Write.Errors, e.Error())
					progress.Write.AddError(1)
					progress.callback(progress)
					if !progress.ErrorContinue {
						err = e
						return
					}
				} else {
					line := strings.Join(stringValues, colSeparator)
					_, _ = file.WriteString(string(rowSeparator) + line)
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
