package datamove

import (
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

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

func (this_ *Executor) excelToKafka() (err error) {
	util.Logger.Info("excel to kafka start")
	err = this_.onExcelSourceData(this_.datasourceToKafka)
	util.Logger.Info("excel to kafka end")
	return
}

func (this_ *Executor) excelToRedis() (err error) {
	util.Logger.Info("excel to redis start")
	err = this_.onExcelSourceData(this_.datasourceToRedis)
	util.Logger.Info("excel to redis end")
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
	datasource.FilePath = this_.From.FilePath
	datasource.ColumnList = this_.From.ColumnList
	datasource.FillColumn = this_.From.FillColumn
	datasource.DataSourceExcelParam = &this_.From.DataSourceExcelParam
	err = on(datasource)
	return
}

func (this_ *Executor) datasourceToExcel(from DataSource) (err error) {
	util.Logger.Info("datasource to excel start")
	to := NewDataSourceExcel()
	to.ColumnList = this_.To.ColumnList
	to.FillColumn = this_.To.FillColumn
	to.DataSourceExcelParam = &this_.To.DataSourceExcelParam
	to.FilePath = this_.getFilePath("", this_.To.GetFileName(), "xlsx")
	err = DateMove(this_.Progress, from, to)
	if err != nil {
		util.Logger.Error("datasource to excel error", zap.Error(err))
		return
	}
	util.Logger.Info("datasource to excel end")
	return
}
