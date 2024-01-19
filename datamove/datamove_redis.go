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
	defer func() {
		if datasource.Service != nil {
			datasource.Service.Close()
		}
		if this_.From.RedisConfig.SSHClient != nil {
			_ = this_.From.RedisConfig.SSHClient.Close()
		}
	}()
	datasource.FillColumn = this_.From.FillColumn
	datasource.ColumnList = this_.From.ColumnList
	if this_.From.DataSourceRedisParam != nil {
		datasource.DataSourceRedisParam = this_.From.DataSourceRedisParam
	}
	datasource.Service, err = redis.New(this_.From.RedisConfig)
	if err != nil {
		util.Logger.Error("redis client new error", zap.Error(err))
		return
	}
	err = on(datasource)
	return
}

func (this_ *Executor) datasourceToRedis(from DataSource) (err error) {
	util.Logger.Info("datasource to redis start")
	to := NewDataSourceRedis()
	defer func() {
		if to.Service != nil {
			to.Service.Close()
		}
		if this_.To.RedisConfig.SSHClient != nil {
			_ = this_.To.RedisConfig.SSHClient.Close()
		}
	}()
	to.ColumnList = this_.To.ColumnList
	to.FillColumn = this_.To.FillColumn
	if this_.To.DataSourceRedisParam != nil {
		to.DataSourceRedisParam = this_.To.DataSourceRedisParam
	}
	to.Service, err = redis.New(this_.To.RedisConfig)
	if err != nil {
		util.Logger.Error("redis client new error", zap.Error(err))
		return
	}
	err = DateMove(this_.Progress, from, to)
	if err != nil {
		util.Logger.Error("datasource to redis error", zap.Error(err))
		return
	}
	util.Logger.Info("datasource to redis end")
	return
}
