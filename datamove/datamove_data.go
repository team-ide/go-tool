package datamove

import (
	"github.com/team-ide/go-tool/util"
)

func (this_ *Executor) dataToDb() (err error) {
	util.Logger.Info("data to db start")
	err = this_.onDataSourceData(this_.datasourceToDb)
	util.Logger.Info("data to db end")
	return
}

func (this_ *Executor) dataToEs() (err error) {
	util.Logger.Info("data to es start")
	err = this_.onDataSourceData(this_.datasourceToEs)
	util.Logger.Info("data to es end")
	return
}

func (this_ *Executor) dataToKafka() (err error) {
	util.Logger.Info("data to kafka start")
	err = this_.onDataSourceData(this_.datasourceToKafka)
	util.Logger.Info("data to kafka end")
	return
}

func (this_ *Executor) dataToRedis() (err error) {
	util.Logger.Info("data to redis start")
	err = this_.onDataSourceData(this_.datasourceToRedis)
	util.Logger.Info("data to redis end")
	return
}

func (this_ *Executor) dataToSql() (err error) {
	util.Logger.Info("data to sql start")
	err = this_.onDataSourceData(this_.datasourceToSql)
	util.Logger.Info("data to sql end")
	return
}

func (this_ *Executor) dataToTxt() (err error) {
	util.Logger.Info("data to txt start")
	err = this_.onDataSourceData(this_.datasourceToTxt)
	util.Logger.Info("data to txt end")
	return
}

func (this_ *Executor) dataToExcel() (err error) {
	util.Logger.Info("data to excel start")
	err = this_.onDataSourceData(this_.datasourceToExcel)
	util.Logger.Info("data to excel end")
	return
}

func (this_ *Executor) onDataSourceData(on func(datasource DataSource) (err error)) (err error) {
	datasource := NewDataSourceData()
	datasource.ColumnList = this_.From.ColumnList
	datasource.FillColumn = this_.From.FillColumn
	datasource.DataSourceDataParam = &this_.From.DataSourceDataParam
	err = on(datasource)
	return
}
