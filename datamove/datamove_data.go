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
	err = this_.onDataSourceData(this_.datasourceToDb)
	util.Logger.Info("data to es end")
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
	err = this_.onDataSourceData(this_.datasourceToTxt)
	util.Logger.Info("data to excel end")
	return
}

func (this_ *Executor) onDataSourceData(on func(datasource DataSource) (err error)) (err error) {
	datasource := NewDataSourceData()
	datasource.DataList = this_.DataList
	datasource.ColumnList = this_.ColumnList
	err = on(datasource)
	return
}
