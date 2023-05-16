package context_service

import (
	"github.com/team-ide/go-tool/db"
	"github.com/team-ide/go-tool/elasticsearch"
	"github.com/team-ide/go-tool/javascript/context_map"
	"github.com/team-ide/go-tool/kafka"
	"github.com/team-ide/go-tool/redis"
	"github.com/team-ide/go-tool/zookeeper"
)

func init() {

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "newRedisService",
		Comment: "新建 Redis 服务",
		Func:    redis.NewServiceScript,
	})
	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "newRedisParam",
		Comment: "新建 Redis 参数",
		Func:    redis.NewParam,
	})
	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "newRedisSizeArg",
		Comment: "新建 Redis 参数",
		Func:    redis.NewSizeArg,
	})
	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "newRedisStartArg",
		Comment: "新建 Redis 参数",
		Func:    redis.NewStartArg,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "newZookeeperService",
		Comment: "新建 Zookeeper 服务",
		Func:    zookeeper.NewServiceScript,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "newKafkaService",
		Comment: "新建 Kafka 服务",
		Func:    kafka.NewServiceScript,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "newElasticsearchService",
		Comment: "新建 Elasticsearch 服务",
		Func:    elasticsearch.NewServiceScript,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "newDbService",
		Comment: "新建 Db 服务",
		Func:    db.NewServiceScript,
	})
}
