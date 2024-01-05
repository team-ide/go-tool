package datamove

import (
	"errors"
	"fmt"
	"github.com/team-ide/go-dialect/dialect"
)

type DataSourceData struct {
	DataType       DataType
	SqlList        []string
	DataList       []map[string]interface{}
	ColumnList     []*dialect.ColumnModel
	Total          int64
	nextStartIndex int64
}

func (this_ *DataSourceData) Stop() {

}

func (this_ *DataSourceData) ReadStart() (err error) {
	if err = ValidateDataType(this_.DataType); err != nil {
		return
	}
	this_.nextStartIndex = 0

	return
}

func (this_ *DataSourceData) Read() (data *Data, err error) {
	pageSize := int64(100)
	total := this_.Total

	startIndex := this_.nextStartIndex
	endIndex := startIndex + pageSize
	if endIndex > total {
		endIndex = total
	}
	data = &Data{}
	data.DataType = this_.DataType
	data.ColumnList = this_.ColumnList
	switch this_.DataType {
	case DataTypeMap:
		data.DataList = this_.DataList[startIndex:endIndex]
		break
	case DataTypeSql:
		data.SqlList = this_.SqlList[startIndex:endIndex]
		break
	}
	data.HasNext = endIndex < total

	return
}

func (this_ *DataSourceData) ReadEnd() (err error) {
	return
}

func (this_ *DataSourceData) WriteStart() (err error) {
	this_.DataType = DataTypeEmpty
	this_.SqlList = []string{}
	this_.DataList = []map[string]interface{}{}
	this_.ColumnList = []*dialect.ColumnModel{}
	return
}

func ValidateDataType(dataType DataType) (err error) {
	if dataType == DataTypeSql || dataType == DataTypeMap {
		return
	}
	err = errors.New(fmt.Sprintf("不支持的数据类型[%d]", dataType))
	return
}

func (this_ *DataSourceData) Write(data *Data) (err error) {
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
	var size int

	switch this_.DataType {
	case DataTypeMap:
		size = len(data.DataList)
		if size > 0 {
			this_.DataList = append(this_.DataList, data.DataList...)
		}
		break
	case DataTypeSql:
		size = len(data.SqlList)
		if size > 0 {
			this_.SqlList = append(this_.SqlList, data.SqlList...)
		}
		break
	}
	this_.Total += int64(size)
	return
}

func (this_ *DataSourceData) WriteEnd() (err error) {
	return
}
