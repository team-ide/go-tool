package datamove

import "github.com/team-ide/go-tool/util"

func (this_ *Executor) redisToDb() (err error) {
	util.Logger.Info("redis to db start")
	err = this_.onRedisSourceData(this_.datasourceToDb)
	util.Logger.Info("redis to db end")
	return
}

func (this_ *Executor) redisToEs() (err error) {
	util.Logger.Info("redis to es start")
	err = this_.onRedisSourceData(this_.datasourceToEs)
	util.Logger.Info("redis to es end")
	return
}

func (this_ *Executor) redisToSql() (err error) {
	util.Logger.Info("redis to sql start")
	err = this_.onRedisSourceData(this_.datasourceToSql)
	util.Logger.Info("redis to sql end")
	return
}

func (this_ *Executor) redisToTxt() (err error) {
	util.Logger.Info("redis to txt start")
	err = this_.onRedisSourceData(this_.datasourceToTxt)
	util.Logger.Info("redis to txt end")
	return
}

func (this_ *Executor) redisToExcel() (err error) {
	util.Logger.Info("redis to excel start")
	err = this_.onRedisSourceData(this_.datasourceToExcel)
	util.Logger.Info("redis to excel end")
	return
}

func (this_ *Executor) onRedisSourceData(on func(datasource DataSource) (err error)) (err error) {
	datasource := NewDataSourceRedis()
	err = on(datasource)
	return
}
