package context_service

import (
	"github.com/team-ide/go-tool/db"
	"github.com/team-ide/go-tool/elasticsearch"
	"github.com/team-ide/go-tool/javascript/context_map"
	"github.com/team-ide/go-tool/kafka"
	"github.com/team-ide/go-tool/redis"
	"github.com/team-ide/go-tool/zookeeper"
)

var (
	redisModule = &context_map.ModuleInfo{
		Name:    "redis",
		Comment: "Redis 模块",
		FuncList: []*context_map.FuncInfo{
			{
				Name:    "newService",
				Comment: "新建 Redis 服务",
				Func:    redis.NewServiceScript,
			}, {
				Name:    "newParam",
				Comment: "新建 Redis 参数",
				Func:    redis.NewParam,
			}, {
				Name:    "newSizeArg",
				Comment: "新建 Redis 参数",
				Func:    redis.NewSizeArg,
			}, {
				Name:    "newStartArg",
				Comment: "新建 Redis 参数",
				Func:    redis.NewStartArg,
			},
		},
	}

	zookeeperModule = &context_map.ModuleInfo{
		Name:    "zookeeper",
		Comment: "Zookeeper 模块",
		FuncList: []*context_map.FuncInfo{
			{
				Name:    "newService",
				Comment: "新建 Zookeeper 服务",
				Func:    zookeeper.NewServiceScript,
			},
		},
	}

	kafkaModule = &context_map.ModuleInfo{
		Name:    "kafka",
		Comment: "Kafka 模块",
		FuncList: []*context_map.FuncInfo{
			{
				Name:    "newService",
				Comment: "新建 Kafka 服务",
				Func:    kafka.NewServiceScript,
			},
		},
	}

	elasticsearchModule = &context_map.ModuleInfo{
		Name:    "elasticsearch",
		Comment: "Elasticsearch 模块",
		FuncList: []*context_map.FuncInfo{
			{
				Name:    "newService",
				Comment: "新建 Elasticsearch 服务",
				Func:    elasticsearch.NewServiceScript,
			},
		},
	}

	dbModule = &context_map.ModuleInfo{
		Name:    "db",
		Comment: "Db 模块",
		FuncList: []*context_map.FuncInfo{
			{
				Name:    "newService",
				Comment: "新建 Db 服务",
				Func:    db.NewServiceScript,
			},
		},
	}
)

func init() {

	context_map.AddModule(dbModule)
	context_map.AddModule(redisModule)
	context_map.AddModule(zookeeperModule)
	context_map.AddModule(elasticsearchModule)
	context_map.AddModule(kafkaModule)
}
