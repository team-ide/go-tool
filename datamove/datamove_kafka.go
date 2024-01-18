package datamove

import (
	"github.com/team-ide/go-tool/kafka"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

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

func (this_ *Executor) kafkaToKafka() (err error) {
	util.Logger.Info("kafka to kafka start")
	err = this_.onKafkaSourceData(this_.datasourceToKafka)
	util.Logger.Info("kafka to kafka end")
	return
}

func (this_ *Executor) kafkaToRedis() (err error) {
	util.Logger.Info("kafka to redis start")
	err = this_.onKafkaSourceData(this_.datasourceToRedis)
	util.Logger.Info("kafka to redis end")
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

	datasource.ColumnList = this_.From.ColumnList
	datasource.TopicName = this_.From.TopicName
	datasource.TopicGroupName = this_.From.TopicGroupName
	datasource.TopicKey = this_.From.TopicKey
	datasource.TopicValue = this_.From.TopicValue
	datasource.TopicValueByData = this_.From.TopicValueByData
	datasource.FillColumn = this_.From.FillColumn
	datasource.Service, err = kafka.New(this_.From.KafkaConfig)
	if err != nil {
		util.Logger.Error("kafka client new error", zap.Error(err))
		return
	}
	err = on(datasource)
	return
}
