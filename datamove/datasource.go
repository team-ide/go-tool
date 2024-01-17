package datamove

import (
	"errors"
	"fmt"
)

type DataSource interface {
	Stop(progress *Progress)
	ReadStart(progress *Progress) (err error)
	Read(progress *Progress, dataChan chan *Data) (err error)
	ReadEnd(progress *Progress) (err error)
	WriteStart(progress *Progress) (err error)
	Write(progress *Progress, data *Data) (err error)
	WriteEnd(progress *Progress) (err error)
}

type Data struct {
	Total        int64           `json:"total"`
	DataType     DataType        `json:"dataType"`
	ColsList     [][]interface{} `json:"colsList"`
	SqlList      []string        `json:"sqlList"`
	SqlAndParams []*SqlAndParam  `json:"sqlAndParams"`
	columnList   *[]*Column
}

type SqlAndParam struct {
	Sql    string        `json:"sql"`
	Params []interface{} `json:"params"`
}

type DataType int8

const (
	DataTypeEmpty        = 0
	DataTypeCols         = 1 // 列数据
	DataTypeSql          = 2 // SQL 语句
	DataTypeSqlAndParams = 3 // 带占位符的SQL语句
)

func ValidateDataType(dataType DataType) (err error) {
	if dataType == DataTypeCols || dataType == DataTypeSql || dataType == DataTypeSqlAndParams {
		return
	}
	err = errors.New(fmt.Sprintf("不支持的数据类型[%d]", dataType))
	return
}

var (
	ErrorStopped = errors.New(`task stopped`)
)
