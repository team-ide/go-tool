package datamove

import "github.com/team-ide/go-tool/util"

func (this_ *Executor) scriptToDb() (err error) {
	util.Logger.Info("script to db start")
	err = this_.onScriptSourceData(this_.datasourceToDb)
	util.Logger.Info("script to db end")
	return
}

func (this_ *Executor) scriptToEs() (err error) {
	util.Logger.Info("script to es start")
	err = this_.onScriptSourceData(this_.datasourceToEs)
	util.Logger.Info("script to es end")
	return
}

func (this_ *Executor) scriptToKafka() (err error) {
	util.Logger.Info("script to kafka start")
	err = this_.onScriptSourceData(this_.datasourceToKafka)
	util.Logger.Info("script to kafka end")
	return
}

func (this_ *Executor) scriptToRedis() (err error) {
	util.Logger.Info("script to redis start")
	err = this_.onScriptSourceData(this_.datasourceToRedis)
	util.Logger.Info("script to redis end")
	return
}

func (this_ *Executor) scriptToSql() (err error) {
	util.Logger.Info("script to sql start")
	err = this_.onScriptSourceData(this_.datasourceToSql)
	util.Logger.Info("script to sql end")
	return
}

func (this_ *Executor) scriptToTxt() (err error) {
	util.Logger.Info("script to txt start")
	err = this_.onScriptSourceData(this_.datasourceToTxt)
	util.Logger.Info("script to txt end")
	return
}

func (this_ *Executor) scriptToExcel() (err error) {
	util.Logger.Info("script to excel start")
	err = this_.onScriptSourceData(this_.datasourceToExcel)
	util.Logger.Info("script to excel end")
	return
}

func (this_ *Executor) onScriptSourceData(on func(datasource DataSource) (err error)) (err error) {
	datasource := NewDataSourceScript()
	datasource.ColumnList = this_.From.ColumnList
	datasource.FillColumn = this_.From.FillColumn
	datasource.DataSourceScriptParam = &this_.From.DataSourceScriptParam
	err = on(datasource)
	return
}
