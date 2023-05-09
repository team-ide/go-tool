package context_service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/team-ide/go-tool/db"
	"github.com/team-ide/go-tool/elasticsearch"
	"github.com/team-ide/go-tool/kafka"
	"github.com/team-ide/go-tool/redis"
	"github.com/team-ide/go-tool/zookeeper"
)

func NewRedisService(config interface{}) (res redis.IService, err error) {
	if config == nil {
		err = errors.New("config is null")
		return
	}
	c1, ok := config.(*redis.Config)
	if ok {
		return redis.New(c1)
	}
	c2, ok := config.(redis.Config)
	if ok {
		return redis.New(&c2)
	}
	bs, err := json.Marshal(config)
	if err != nil {
		err = errors.New("config to json error:" + err.Error())
		return
	}
	var c3 = &redis.Config{}
	err = json.Unmarshal(bs, c3)
	if err != nil {
		err = errors.New("config to json error:" + err.Error())
		return
	}
	return redis.New(c3)
}

func NewRedisParam(args ...interface{}) *redis.Param {
	res := &redis.Param{}
	for _, arg := range args {
		if v, ok := arg.(int); ok {
			res.Database = v
		} else if v, ok := arg.(int64); ok {
			res.Database = int(v)
		} else if v, ok := arg.(float64); ok {
			res.Database = int(v)
		} else if v, ok := arg.(context.Context); ok {
			res.Ctx = v
		}
	}
	return res
}

func NewZookeeperService(config interface{}) (res zookeeper.IService, err error) {
	if config == nil {
		err = errors.New("zookeeper config is null")
		return
	}
	c1, ok := config.(*zookeeper.Config)
	if ok {
		return zookeeper.New(c1)
	}
	c2, ok := config.(zookeeper.Config)
	if ok {
		return zookeeper.New(&c2)
	}
	bs, err := json.Marshal(config)
	if err != nil {
		err = errors.New("zookeeper config to json error:" + err.Error())
		return
	}
	var c3 = &zookeeper.Config{}
	err = json.Unmarshal(bs, c3)
	if err != nil {
		err = errors.New("zookeeper config to json error:" + err.Error())
		return
	}
	return zookeeper.New(c3)
}

func NewKafkaService(config interface{}) (res kafka.IService, err error) {
	if config == nil {
		err = errors.New("kafka config is null")
		return
	}
	c1, ok := config.(*kafka.Config)
	if ok {
		return kafka.New(c1)
	}
	c2, ok := config.(kafka.Config)
	if ok {
		return kafka.New(&c2)
	}
	bs, err := json.Marshal(config)
	if err != nil {
		err = errors.New("kafka config to json error:" + err.Error())
		return
	}
	var c3 = &kafka.Config{}
	err = json.Unmarshal(bs, c3)
	if err != nil {
		err = errors.New("kafka config to json error:" + err.Error())
		return
	}
	return kafka.New(c3)
}

func NewElasticsearchService(config interface{}) (res elasticsearch.IService, err error) {
	if config == nil {
		err = errors.New("elasticsearch config is null")
		return
	}
	c1, ok := config.(*elasticsearch.Config)
	if ok {
		return elasticsearch.New(c1)
	}
	c2, ok := config.(elasticsearch.Config)
	if ok {
		return elasticsearch.New(&c2)
	}
	bs, err := json.Marshal(config)
	if err != nil {
		err = errors.New("elasticsearch config to json error:" + err.Error())
		return
	}
	var c3 = &elasticsearch.Config{}
	err = json.Unmarshal(bs, c3)
	if err != nil {
		err = errors.New("elasticsearch config to json error:" + err.Error())
		return
	}
	return elasticsearch.New(c3)
}

func NewDbService(config interface{}) (res db.IService, err error) {
	if config == nil {
		err = errors.New("db config is null")
		return
	}
	c1, ok := config.(*db.Config)
	if ok {
		return db.New(c1)
	}
	c2, ok := config.(db.Config)
	if ok {
		return db.New(&c2)
	}
	bs, err := json.Marshal(config)
	if err != nil {
		err = errors.New("db config to json error:" + err.Error())
		return
	}
	var c3 = &db.Config{}
	err = json.Unmarshal(bs, c3)
	if err != nil {
		err = errors.New("db config to json error:" + err.Error())
		return
	}
	return db.New(c3)
}
