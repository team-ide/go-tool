package datamove

import (
	"github.com/team-ide/go-tool/redis"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

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

func (this_ *Executor) redisToKafka() (err error) {
	util.Logger.Info("redis to kafka start")
	err = this_.onRedisSourceData(this_.datasourceToKafka)
	util.Logger.Info("redis to kafka end")
	return
}

func (this_ *Executor) redisToRedis() (err error) {
	util.Logger.Info("redis to redis start")
	err = this_.onRedisSourceData(this_.datasourceToRedis)
	util.Logger.Info("redis to redis end")
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
	datasource.FillColumn = this_.From.FillColumn
	datasource.Service, err = redis.New(this_.From.RedisConfig)
	if err != nil {
		util.Logger.Error("redis client new error", zap.Error(err))
		return
	}
	err = on(datasource)
	return
}
