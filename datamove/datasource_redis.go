package datamove

import (
	"github.com/team-ide/go-tool/redis"
)

func NewDataSourceRedis() *DataSourceRedis {
	return &DataSourceRedis{
		DataSourceBase: &DataSourceBase{},
	}
}

type DataSourceRedis struct {
	*DataSourceBase

	Service redis.IService
}

func (this_ *DataSourceRedis) Stop(progress *Progress) {

}

func (this_ *DataSourceRedis) ReadStart(progress *Progress) (err error) {

	return
}

func (this_ *DataSourceRedis) Read(progress *Progress, dataChan chan *Data) (err error) {

	return
}

func (this_ *DataSourceRedis) ReadEnd(progress *Progress) (err error) {
	return
}

func (this_ *DataSourceRedis) WriteStart(progress *Progress) (err error) {

	return
}

func (this_ *DataSourceRedis) Write(progress *Progress, data *Data) (err error) {

	if this_.FillColumn && data.columnList != nil {
		this_.fullColumnListByColumnList(progress, data.columnList)
	}

	return
}

func (this_ *DataSourceRedis) WriteEnd(progress *Progress) (err error) {
	return
}
