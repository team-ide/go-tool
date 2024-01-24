package datamove

import (
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

func (this_ *Executor) sqlToDb() (err error) {
	util.Logger.Info("sql to db start")
	err = this_.onSqlSourceData(this_.datasourceToDb)
	util.Logger.Info("sql to db end")
	return
}

func (this_ *Executor) onSqlSourceData(on func(datasource DataSource) (err error)) (err error) {
	datasource := NewDataSourceSql()
	datasource.FilePath = this_.From.FilePath
	datasource.ColumnList = this_.From.ColumnList
	datasource.FillColumn = this_.From.FillColumn
	datasource.DataSourceSqlParam = &this_.From.DataSourceSqlParam

	err = on(datasource)
	return
}

func (this_ *Executor) datasourceToSql(from DataSource) (err error) {
	util.Logger.Info("datasource to sql start")
	to := NewDataSourceSql()
	to.ParamModel = this_.To.GetDialectParam()
	to.ColumnList = this_.To.ColumnList
	to.FillColumn = this_.To.FillColumn
	to.DataSourceSqlParam = &this_.To.DataSourceSqlParam
	to.FilePath = this_.getFilePath("", this_.To.GetFileName(), "sql")
	err = DateMove(this_.Progress, from, to)
	if err != nil {
		util.Logger.Error("datasource to sql error", zap.Error(err))
		return
	}
	util.Logger.Info("datasource to sql end")
	return
}
