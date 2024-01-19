package datamove

import (
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

func (this_ *Executor) txtToDb() (err error) {
	util.Logger.Info("txt to db start")
	err = this_.onTxtSourceData(this_.datasourceToDb)
	util.Logger.Info("txt to db end")
	return
}

func (this_ *Executor) txtToEs() (err error) {
	util.Logger.Info("txt to es start")
	err = this_.onTxtSourceData(this_.datasourceToEs)
	util.Logger.Info("txt to es end")
	return
}

func (this_ *Executor) txtToKafka() (err error) {
	util.Logger.Info("txt to kafka start")
	err = this_.onTxtSourceData(this_.datasourceToKafka)
	util.Logger.Info("txt to kafka end")
	return
}

func (this_ *Executor) txtToRedis() (err error) {
	util.Logger.Info("txt to redis start")
	err = this_.onTxtSourceData(this_.datasourceToRedis)
	util.Logger.Info("txt to redis end")
	return
}

func (this_ *Executor) txtToSql() (err error) {
	util.Logger.Info("txt to sql start")
	err = this_.onTxtSourceData(this_.datasourceToSql)
	util.Logger.Info("txt to sql end")
	return
}

func (this_ *Executor) txtToTxt() (err error) {
	util.Logger.Info("txt to txt start")
	err = this_.onTxtSourceData(this_.datasourceToTxt)
	util.Logger.Info("txt to txt end")
	return
}

func (this_ *Executor) txtToExcel() (err error) {
	util.Logger.Info("txt to excel start")
	err = this_.onTxtSourceData(this_.datasourceToExcel)
	util.Logger.Info("txt to excel end")
	return
}

func (this_ *Executor) onTxtSourceData(on func(datasource DataSource) (err error)) (err error) {
	datasource := NewDataSourceTxt()
	datasource.FilePath = this_.From.FilePath
	datasource.ColumnList = this_.From.ColumnList
	datasource.FillColumn = this_.From.FillColumn
	if this_.From.DataSourceTxtParam != nil {
		datasource.DataSourceTxtParam = this_.From.DataSourceTxtParam
	}
	err = on(datasource)
	return
}

func (this_ *Executor) datasourceToTxt(from DataSource) (err error) {
	util.Logger.Info("datasource to text start")
	to := NewDataSourceTxt()
	to.ColumnList = this_.To.ColumnList
	to.FillColumn = this_.To.FillColumn
	if this_.To.DataSourceTxtParam != nil {
		to.DataSourceTxtParam = this_.To.DataSourceTxtParam
	}

	to.FilePath = this_.getFilePath("", this_.To.GetFileName(), this_.To.GetTxtFileType())
	err = DateMove(this_.Progress, from, to)
	if err != nil {
		util.Logger.Error("datasource to text error", zap.Error(err))
		return
	}
	util.Logger.Info("datasource to text end")
	return
}
