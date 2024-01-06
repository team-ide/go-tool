package datamove

import (
	"errors"
	"fmt"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

type DataSourceData struct {
	*Data
}

func (this_ *DataSourceData) Stop(progress *DateMoveProgress) {

}

func (this_ *DataSourceData) ReadStart(progress *DateMoveProgress) (err error) {
	if this_.Data == nil {
		this_.Data = &Data{}
	}
	if err = ValidateDataType(this_.DataType); err != nil {
		return
	}

	switch this_.DataType {
	case DataTypeData:
		this_.Total = int64(len(this_.DataList))
		break
	case DataTypeSql:
		this_.Total = int64(len(this_.SqlList))
		break
	case DataTypeSqlAndParams:
		this_.Total = int64(len(this_.SqlAndParams))
		break
	default:
		err = errors.New(fmt.Sprintf("当前数据类型[%d]，不支持读取", this_.DataType))
		return
	}
	progress.Total = this_.Total
	return
}

func (this_ *DataSourceData) Read(progress *DateMoveProgress, dataChan chan *Data) (err error) {

	pageSize := progress.BatchNumber
	total := this_.Total
	var nextStartIndex int64

	for {
		if progress.ShouldStop() {
			return
		}

		startIndex := nextStartIndex
		endIndex := startIndex + pageSize - 1
		if endIndex >= total {
			endIndex = total - 1
		}
		if endIndex < 0 {
			return
		}
		nextStartIndex = endIndex + 1
		data := &Data{}
		data.DataType = this_.DataType
		data.ColumnList = this_.ColumnList
		// 获取包含结束索引的数据
		endIndex++
		var size int
		switch this_.DataType {
		case DataTypeData:
			data.DataList = this_.DataList[startIndex:endIndex]
			size = len(data.DataList)
			break
		case DataTypeSql:
			data.SqlList = this_.SqlList[startIndex:endIndex]
			size = len(data.SqlList)
			break
		case DataTypeSqlAndParams:
			data.SqlAndParams = this_.SqlAndParams[startIndex:endIndex]
			size = len(data.SqlAndParams)
			break
		default:
			err = errors.New(fmt.Sprintf("当前数据类型[%d]，不支持读取", data.DataType))
			return
		}
		data.Total = int64(size)
		progress.Read.AddSuccess(data.Total)
		util.Logger.Info("read data source", zap.Any("total", data.Total))
		progress.dataChan <- data
		progress.callback(progress)
		if endIndex >= progress.Total {
			break
		}
	}

	return
}

func (this_ *DataSourceData) ReadEnd(progress *DateMoveProgress) (err error) {
	return
}

func (this_ *DataSourceData) WriteStart(progress *DateMoveProgress) (err error) {
	if this_.Data == nil {
		this_.Data = &Data{}
	}
	this_.Total = 0
	this_.DataType = DataTypeEmpty
	this_.SqlList = nil
	this_.SqlAndParams = nil
	this_.DataList = nil
	this_.ColumnList = nil
	return
}

func (this_ *DataSourceData) Write(progress *DateMoveProgress, data *Data) (err error) {
	if this_.DataType == DataTypeEmpty {
		this_.DataType = data.DataType
	}
	if err = ValidateDataType(data.DataType); err != nil {
		return
	}
	if this_.DataType != data.DataType {
		err = errors.New(fmt.Sprintf("当前数据类型[%d]，写入数据类型[%d]，数据类型不一致", this_.DataType, data.DataType))
		return
	}

	switch this_.DataType {
	case DataTypeData:
		data.Total = int64(len(data.DataList))
		if data.Total > 0 {
			this_.DataList = append(this_.DataList, data.DataList...)
		}
		break
	case DataTypeSql:
		data.Total = int64(len(data.SqlList))
		if data.Total > 0 {
			this_.SqlList = append(this_.SqlList, data.SqlList...)
		}
	case DataTypeSqlAndParams:
		data.Total = int64(len(data.SqlAndParams))
		if data.Total > 0 {
			this_.SqlAndParams = append(this_.SqlAndParams, data.SqlAndParams...)
		}
		break
	default:
		err = errors.New(fmt.Sprintf("当前数据类型[%d]，不支持写入", data.DataType))
		return
	}
	progress.Write.AddSuccess(data.Total)
	this_.Total += data.Total
	return
}

func (this_ *DataSourceData) WriteEnd(progress *DateMoveProgress) (err error) {
	return
}
