package datamove

import "github.com/team-ide/go-tool/util"

func (this_ *Executor) excelToDb() (err error) {
	util.Logger.Info("excel to db start")
	err = this_.onExcelSourceData(this_.datasourceToDb)
	util.Logger.Info("excel to db end")
	return
}

func (this_ *Executor) excelToEs() (err error) {
	util.Logger.Info("excel to es start")
	err = this_.onExcelSourceData(this_.datasourceToEs)
	util.Logger.Info("excel to es end")
	return
}

func (this_ *Executor) excelToSql() (err error) {
	util.Logger.Info("excel to sql start")
	err = this_.onExcelSourceData(this_.datasourceToSql)
	util.Logger.Info("excel to sql end")
	return
}

func (this_ *Executor) excelToTxt() (err error) {
	util.Logger.Info("excel to txt start")
	err = this_.onExcelSourceData(this_.datasourceToTxt)
	util.Logger.Info("excel to txt end")
	return
}

func (this_ *Executor) excelToExcel() (err error) {
	util.Logger.Info("excel to excel start")
	err = this_.onExcelSourceData(this_.datasourceToExcel)
	util.Logger.Info("excel to excel end")
	return
}

func (this_ *Executor) onExcelSourceData(on func(datasource DataSource) (err error)) (err error) {
	datasource := NewDataSourceExcel()
	datasource.FilePath = this_.FilePath
	datasource.ColumnList = this_.ColumnList
	err = on(datasource)
	return
}
