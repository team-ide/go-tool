package datamove

import "github.com/team-ide/go-tool/util"

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
	err = on(datasource)
	return
}
