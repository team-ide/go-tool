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

	datasource.ColumnList = this_.ColumnList
	datasource.IndexName = this_.IndexName
	datasource.IdName = this_.IdName
	datasource.IdScript = this_.IdScript
	datasource.Service, err = elasticsearch.New(this_.From.EsConfig)
	if err != nil {
		util.Logger.Error("elasticsearch client new error", zap.Error(err))
		return
	}
	err = on(datasource)
	return
}
