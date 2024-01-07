package datamove

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type DataSource interface {
	Stop(progress *DateMoveProgress)
	ReadStart(progress *DateMoveProgress) (err error)
	Read(progress *DateMoveProgress, dataChan chan *Data) (err error)
	ReadEnd(progress *DateMoveProgress) (err error)
	WriteStart(progress *DateMoveProgress) (err error)
	Write(progress *DateMoveProgress, data *Data) (err error)
	WriteEnd(progress *DateMoveProgress) (err error)
}

type DateMoveProgress struct {
	*Param
	Total int64 `json:"total"`

	Read  *DateMoveProgressInfo `json:"read"`
	Write *DateMoveProgressInfo `json:"write"`

	dataChan  chan *Data
	stopWait  sync.WaitGroup
	isEnd     bool
	isStopped bool
	callback  func(progress *DateMoveProgress)
}

type DateMoveProgressInfo struct {
	Total     int64     `json:"total"`
	Error     int64     `json:"error"`
	Success   int64     `json:"success"`
	Errors    []string  `json:"errors"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

func (this_ *DateMoveProgressInfo) AddSuccess(size int64) {
	this_.Total += size
	this_.Success += size
}

func (this_ *DateMoveProgressInfo) AddError(size int64) {
	this_.Total += size
	this_.Error += size
}

func (this_ *DateMoveProgress) ShouldStop() bool {
	return this_.isEnd || this_.isStopped
}

type Data struct {
	Total        int64           `json:"total"`
	DataType     DataType        `json:"dataType"`
	ColsList     [][]interface{} `json:"colsList"`
	SqlList      []string        `json:"sqlList"`
	SqlAndParams []*SqlAndParam  `json:"sqlAndParams"`
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

type Param struct {
	ErrorContinue bool  `json:"errorContinue"`
	BatchNumber   int64 `json:"batchNumber"`
}

func (this_ *Param) init() {
	if this_.BatchNumber <= 0 {
		this_.BatchNumber = 100
	}
}

var (
	ErrorStopped = errors.New(`task stopped`)
)
