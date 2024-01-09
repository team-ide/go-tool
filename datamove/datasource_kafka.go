package datamove

import (
	"github.com/team-ide/go-tool/kafka"
)

func NewDataSourceKafka() *DataSourceKafka {
	return &DataSourceKafka{
		DataSourceBase: &DataSourceBase{},
	}
}

type DataSourceKafka struct {
	*DataSourceBase
	Topic   string `json:"topic"`
	GroupId string `json:"groupId"`

	Service kafka.IService
}

func (this_ *DataSourceKafka) Stop(progress *Progress) {

}

func (this_ *DataSourceKafka) ReadStart(progress *Progress) (err error) {

	return
}

func (this_ *DataSourceKafka) Read(progress *Progress, dataChan chan *Data) (err error) {

	return
}

func (this_ *DataSourceKafka) ReadEnd(progress *Progress) (err error) {
	return
}

func (this_ *DataSourceKafka) WriteStart(progress *Progress) (err error) {

	return
}

func (this_ *DataSourceKafka) Write(progress *Progress, data *Data) (err error) {

	return
}

func (this_ *DataSourceKafka) WriteEnd(progress *Progress) (err error) {
	return
}
