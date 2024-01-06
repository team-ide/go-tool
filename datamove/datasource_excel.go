package datamove

import (
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-dialect/worker"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"strings"
)

type DataSourceExcel struct {
	FilePath string `json:"filePath"`

	readFile  *xlsx.File
	writeFile *xlsx.File

	readSheet  *xlsx.Sheet
	writeSheet *xlsx.Sheet

	headerRead        bool
	headerWrite       bool
	readColumnList    []*dialect.ColumnModel
	readColumnLength  int
	ColumnNameMapping map[string]string `json:"columnNameMapping"`
	ShouldTrimSpace   bool              `json:"shouldTrimSpace"` // 是否需要去除空白字符
	SheetName         string            `json:"sheetName"`
}

func (this_ *DataSourceExcel) Stop(progress *DateMoveProgress) {
	this_.CloseReadFile()
	this_.CloseWriteFile()
}

func (this_ *DataSourceExcel) ReadStart(progress *DateMoveProgress) (err error) {
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
	return
}

func (this_ *DataSourceExcel) ReadEnd(progress *DateMoveProgress) (err error) {
	this_.CloseReadFile()
	return
}

func (this_ *DataSourceExcel) WriteStart(progress *DateMoveProgress) (err error) {
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

func (this_ *DataSourceExcel) WriteEnd(progress *DateMoveProgress) (err error) {

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

func (this_ *DataSourceExcel) ColsToData(progress *DateMoveProgress, cols []string, columnList []*dialect.ColumnModel) (data map[string]interface{}, err error) {
	data = map[string]interface{}{}
	for index, col := range cols {
		v := col
		if this_.ShouldTrimSpace {
			v = strings.TrimSpace(v)
		}
		name := columnList[index].ColumnName
		//if this_.ColumnNameMapping != nil && this_.ColumnNameMapping[name] != "" {
		//	name = this_.ColumnNameMapping[name]
		//}
		data[name] = v
	}
	return
}

func (this_ *DataSourceExcel) DataToCols(progress *DateMoveProgress, data map[string]interface{}, columnList []*dialect.ColumnModel) (cols []string, err error) {
	for _, column := range columnList {
		v := util.GetStringValue(data[column.ColumnName])
		if this_.ShouldTrimSpace {
			v = strings.TrimSpace(v)
		}
		cols = append(cols, v)
	}
	return
}

func (this_ *DataSourceExcel) ReadLineCount() (lineCount int64, err error) {
	lineCount = int64(len(this_.readSheet.Rows))
	util.Logger.Info("excel data source read line count", zap.Any("lineCount", lineCount))
	return
}

func (this_ *DataSourceExcel) Read(progress *DateMoveProgress, dataChan chan *Data) (err error) {

	var lastData = &Data{
		DataType: DataTypeData,
	}
	lineCount, err := this_.ReadLineCount()
	if err != nil {
		return
	}
	progress.Total = lineCount - 1 // 第一行为头信息
	pageSize := progress.BatchNumber
	for _, row := range this_.readSheet.Rows {
		if progress.ShouldStop() {
			return
		}
		var cols []string
		for _, c := range row.Cells {
			cols = append(cols, c.String())
		}

		if !this_.headerRead {
			for _, name := range cols {
				if this_.ShouldTrimSpace {
					name = strings.TrimSpace(name)
				}
				if this_.ColumnNameMapping != nil && this_.ColumnNameMapping[name] != "" {
					name = this_.ColumnNameMapping[name]
				}
				this_.readColumnList = append(this_.readColumnList, &dialect.ColumnModel{
					ColumnName: name,
				})
			}
			this_.readColumnLength = len(this_.readColumnList)
			this_.headerRead = true
			continue
		}
		if len(cols) != this_.readColumnLength {
			err = errors.New("cols [" + strings.Join(cols, ",") + "] can not to column names [" + strings.Join(worker.GetColumnNames(this_.readColumnList), ",") + "]")
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
	if lastData.Total > 0 {
		progress.callback(progress)
		dataChan <- lastData
	}
	return
}

func (this_ *DataSourceExcel) Write(progress *DateMoveProgress, data *Data) (err error) {

	if err = ValidateDataType(data.DataType); err != nil {
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

		row := this_.writeSheet.AddRow()
		for _, v := range names {
			c := row.AddCell()
			c.SetString(v)
		}
		//_ = this_.writeFile.Write(this_.writeFile_)
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
					row := this_.writeSheet.AddRow()
					for _, v := range cols {
						c := row.AddCell()
						c.SetString(v)
					}
					//_ = this_.writeFile.Write(this_.writeFile_)
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
