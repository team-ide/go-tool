package datamove

import "github.com/team-ide/go-tool/util"

func (this_ *Executor) kafkaToDb() (err error) {
	util.Logger.Info("kafka to db start")
	err = this_.onKafkaSourceData(this_.datasourceToDb)
	util.Logger.Info("kafka to db end")
	return
}

func (this_ *Executor) kafkaToEs() (err error) {
	util.Logger.Info("kafka to es start")
	err = this_.onKafkaSourceData(this_.datasourceToEs)
	util.Logger.Info("kafka to es end")
	return
}

func (this_ *Executor) kafkaToSql() (err error) {
	util.Logger.Info("kafka to sql start")
	err = this_.onKafkaSourceData(this_.datasourceToSql)
	util.Logger.Info("kafka to sql end")
	return
}

func (this_ *Executor) kafkaToTxt() (err error) {
	util.Logger.Info("kafka to txt start")
	err = this_.onKafkaSourceData(this_.datasourceToTxt)
	util.Logger.Info("kafka to txt end")
	return
}

func (this_ *Executor) kafkaToExcel() (err error) {
	util.Logger.Info("kafka to excel start")
	err = this_.onKafkaSourceData(this_.datasourceToExcel)
	util.Logger.Info("kafka to excel end")
	return
}

func (this_ *Executor) onKafkaSourceData(on func(datasource DataSource) (err error)) (err error) {
	datasource := NewDataSourceKafka()
	err = on(datasource)
	return
}
