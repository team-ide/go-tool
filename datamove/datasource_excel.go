package datamove

import (
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"strings"
)

type DataSourceExcel struct {
	*DataSourceBase
	FilePath string `json:"filePath"`

	readFile  *xlsx.File
	writeFile *xlsx.File

	readSheet  *xlsx.Sheet
	writeSheet *xlsx.Sheet

	headerRead        bool
	headerWrite       bool
	ColumnNameMapping map[string]string `json:"columnNameMapping"`
	ShouldTrimSpace   bool              `json:"shouldTrimSpace"` // 是否需要去除空白字符
	SheetName         string            `json:"sheetName"`
}

func (this_ *DataSourceExcel) Stop(progress *Progress) {
	this_.CloseReadFile()
	this_.CloseWriteFile()
}

func (this_ *DataSourceExcel) ReadStart(progress *Progress) (err error) {
	file, err := this_.GetReadFile()
	if err != nil {
		return
	}

	if this_.SheetName != "" {
		this_.readSheet = file.Sheet[this_.SheetName]
	} else {
		if len(file.Sheets) > 0 {
			this_.readSheet = file.Sheets[0]
		}
	}
	if this_.readSheet == nil {
		err = errors.New("read file [" + this_.FilePath + "] sheet is not found")
		return
	}
	var titles []string
	if len(this_.readSheet.Cols) > 0 {
		for _, c := range this_.readSheet.Rows[0].Cells {
			s := c.String()
			if this_.ShouldTrimSpace {
				s = strings.TrimSpace(s)
			}
			titles = append(titles, s)
		}
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
	progress.DataTotal += lineCount - 1 // 第一行为头信息
	return
}

func (this_ *DataSourceExcel) ReadEnd(progress *Progress) (err error) {
	this_.CloseReadFile()
	return
}

func (this_ *DataSourceExcel) WriteStart(progress *Progress) (err error) {
	file, err := this_.GetWriteFile()
	if err != nil {
		return
	}
	if this_.SheetName == "" {
		this_.SheetName = "Sheet1"
	}

	this_.writeSheet, err = file.AddSheet(this_.SheetName)
	if err != nil {
		return
	}
	err = this_.writeFile.Save(this_.FilePath)
	if err != nil {
		err = errors.New("write file [" + this_.FilePath + "] add sheet error:" + err.Error())
		return
	}
	return
}

func (this_ *DataSourceExcel) WriteEnd(progress *Progress) (err error) {

	if this_.writeFile != nil {
		err = this_.writeFile.Save(this_.FilePath)
		if err != nil {
			err = errors.New("write file [" + this_.FilePath + "] save error:" + err.Error())
			return
		}
	}
	this_.CloseWriteFile()
	return
}

func (this_ *DataSourceExcel) GetReadFile() (file *xlsx.File, err error) {
	file = this_.readFile
	if file != nil {
		return
	}
	file, err = xlsx.OpenFile(this_.FilePath)
	if err != nil {
		err = errors.New("open file [" + this_.FilePath + "] error:" + err.Error())
		return
	}
	this_.readFile = file
	return
}

func (this_ *DataSourceExcel) CloseReadFile() {
	file := this_.readFile
	this_.readFile = nil
	if file != nil {
		//err := file.Close()
		//if err != nil {
		//	util.Logger.Error("close read file ["+this_.FilePath+"] error", zap.Error(err))
		//	return
		//}
	}
	return
}

func (this_ *DataSourceExcel) GetWriteFile() (file *xlsx.File, err error) {
	file = this_.writeFile
	if file != nil {
		return
	}
	file = xlsx.NewFile()
	this_.writeFile = file
	return
}

func (this_ *DataSourceExcel) CloseWriteFile() {
	file := this_.writeFile
	this_.writeFile = nil
	if file != nil {
		//err := file.Close()
		//if err != nil {
		//	util.Logger.Error("close write file ["+this_.FilePath+"] error", zap.Error(err))
		//	return
		//}
	}
	return
}

func (this_ *DataSourceExcel) ReadLineCount() (lineCount int64, err error) {
	lineCount = int64(len(this_.readSheet.Rows))
	util.Logger.Info("excel data source read line count", zap.Any("lineCount", lineCount))
	return
}

func (this_ *DataSourceExcel) Read(progress *Progress, dataChan chan *Data) (err error) {

	var lastData = &Data{
		DataType: DataTypeCols,
	}
	pageSize := progress.BatchNumber
	for _, row := range this_.readSheet.Rows {
		if progress.ShouldStop() {
			return
		}
		var cols []interface{}
		for _, c := range row.Cells {
			cols = append(cols, c.String())
		}

		if !this_.headerRead {
			this_.headerRead = true
			continue
		}

		values, e := this_.ValuesToValues(progress, cols)
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
	if lastData.Total > 0 {
		dataChan <- lastData
	}
	return
}

func (this_ *DataSourceExcel) Write(progress *Progress, data *Data) (err error) {

	if err = ValidateDataType(data.DataType); err != nil {
		return
	}

	if !this_.headerWrite {
		if len(this_.ColumnList) == 0 {
			err = errors.New("字段信息不存在，无法写入头部信息")
			return
		}
		var names = this_.GetColumnNames()

		row := this_.writeSheet.AddRow()
		for _, v := range names {
			c := row.AddCell()
			c.SetString(v)
		}
		//_ = this_.writeFile.Write(this_.writeFile_)
		this_.headerWrite = true
	}

	switch data.DataType {
	case DataTypeCols:
		data.Total = int64(len(data.ColsList))
		if data.Total > 0 {
			for _, one := range data.ColsList {
				values, e := this_.ValuesToValues(progress, one)
				if e != nil {
					progress.WriteCount.AddError(1, e)
					if !progress.ErrorContinue {
						err = e
						return
					}
				} else {
					row := this_.writeSheet.AddRow()
					for _, v := range values {
						c := row.AddCell()
						switch t := v.(type) {
						case int8:
							c.SetInt(int(t))
						case int16:
							c.SetInt(int(t))
						case int:
							c.SetInt(t)
						case string:
							c.SetString(t)
							break
						default:
							c.SetString(util.GetStringValue(v))
							break
						}
					}
					//_ = this_.writeFile.Write(this_.writeFile_)
					progress.WriteCount.AddSuccess(1)
				}
			}
		}
		break
	default:
		err = errors.New(fmt.Sprintf("当前数据类型[%d]，不支持写入", data.DataType))
		return
	}
	return
}
