package datamove

import (
	"github.com/team-ide/go-tool/elasticsearch"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

func (this_ *Executor) esToDb() (err error) {
	util.Logger.Info("es to db start")
	err = this_.onEsSourceData(this_.datasourceToDb)
	util.Logger.Info("es to db end")
	return
}

func (this_ *Executor) esToEs() (err error) {
	util.Logger.Info("es to es start")
	err = this_.onEsSourceData(this_.datasourceToEs)
	util.Logger.Info("es to es end")
	return
}

func (this_ *Executor) esToKafka() (err error) {
	util.Logger.Info("es to kafka start")
	err = this_.onEsSourceData(this_.datasourceToKafka)
	util.Logger.Info("es to kafka end")
	return
}

func (this_ *Executor) esToRedis() (err error) {
	util.Logger.Info("es to redis start")
	err = this_.onEsSourceData(this_.datasourceToRedis)
	util.Logger.Info("es to redis end")
	return
}

func (this_ *Executor) esToSql() (err error) {
	util.Logger.Info("es to sql start")
	err = this_.onEsSourceData(this_.datasourceToSql)
	util.Logger.Info("es to sql end")
	return
}

func (this_ *Executor) esToTxt() (err error) {
	util.Logger.Info("es to txt start")
	err = this_.onEsSourceData(this_.datasourceToTxt)
	util.Logger.Info("es to txt end")
	return
}

func (this_ *Executor) esToExcel() (err error) {
	util.Logger.Info("es to excel start")
	err = this_.onEsSourceData(this_.datasourceToExcel)
	util.Logger.Info("es to excel end")
	return
}

func (this_ *Executor) onEsSourceData(on func(datasource DataSource) (err error)) (err error) {
	datasource := NewDataSourceEs()
	defer func() {
		if datasource.Service != nil {
			datasource.Service.Close()
		}
	}()
	datasource.ColumnList = this_.From.ColumnList
	datasource.FillColumn = this_.From.FillColumn
	if this_.From.DataSourceEsParam != nil {
		datasource.DataSourceEsParam = this_.From.DataSourceEsParam
	}
	datasource.Service, err = elasticsearch.New(this_.From.EsConfig)
	if err != nil {
		util.Logger.Error("elasticsearch client new error", zap.Error(err))
		return
	}
	err = on(datasource)
	return
}

func (this_ *Executor) datasourceToEs(from DataSource) (err error) {
	util.Logger.Info("datasource to es start")
	to := NewDataSourceEs()
	defer func() {
		if to.Service != nil {
			to.Service.Close()
		}
	}()
	to.ColumnList = this_.To.ColumnList
	to.FillColumn = this_.To.FillColumn
	if this_.To.DataSourceEsParam != nil {
		to.DataSourceEsParam = this_.To.DataSourceEsParam
	}
	to.Service, err = elasticsearch.New(this_.To.EsConfig)
	if err != nil {
		util.Logger.Error("elasticsearch client new error", zap.Error(err))
		return
	}
	err = DateMove(this_.Progress, from, to)
	if err != nil {
		util.Logger.Error("datasource to es error", zap.Error(err))
		return
	}
	util.Logger.Info("datasource to es end")
	return
}
