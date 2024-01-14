package datamove

import "github.com/team-ide/go-tool/util"

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
	err = on(datasource)
	return
}
