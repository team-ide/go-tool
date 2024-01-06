package datamove

import (
	"errors"
	"fmt"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
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

	dataChan       chan *Data
	stopWait       sync.WaitGroup
	nextStartIndex int64
	isEnd          bool
	isStopped      bool
	callback       func(progress *DateMoveProgress)
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

func (this_ *DateMoveProgress) NextIndies() (int64, int64) {
	pageSize := this_.BatchNumber
	total := this_.Total

	startIndex := this_.nextStartIndex
	endIndex := startIndex + pageSize - 1
	if endIndex >= total {
		endIndex = total - 1
	}
	this_.nextStartIndex = endIndex + 1
	util.Logger.Info("NextIndies", zap.Any("startIndex", startIndex), zap.Any("endIndex", endIndex))
	return startIndex, endIndex
}

func (this_ *DateMoveProgress) ShouldStop() bool {
	return this_.isEnd || this_.isStopped
}

type Data struct {
	Total        int64                    `json:"total"`
	DataType     DataType                 `json:"dataType"`
	SqlList      []string                 `json:"sqlList"`
	SqlAndParams []*SqlAndParam           `json:"sqlAndParams"`
	DataList     []map[string]interface{} `json:"dataList"`
	ColumnList   []*dialect.ColumnModel   `json:"columnList"`
}

type SqlAndParam struct {
	Sql    string        `json:"sql"`
	Params []interface{} `json:"params"`
}

type DataType int8

const (
	DataTypeEmpty        = 0
	DataTypeData         = 1
	DataTypeSql          = 2
	DataTypeSqlAndParams = 3
)

func ValidateDataType(dataType DataType) (err error) {
	if dataType == DataTypeData || dataType == DataTypeSql || dataType == DataTypeSqlAndParams {
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
