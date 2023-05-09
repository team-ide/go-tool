package context_service

import (
	"github.com/team-ide/go-tool/javascript/context_map"
)

func init() {

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "newRedisService",
		Comment: "新建 Redis 服务",
		Func:    NewRedisService,
	})
	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "newRedisParam",
		Comment: "新建 Redis 参数",
		Func:    NewRedisParam,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "newZookeeperService",
		Comment: "新建 Zookeeper 服务",
		Func:    NewZookeeperService,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "newKafkaService",
		Comment: "新建 Kafka 服务",
		Func:    NewKafkaService,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "newElasticsearchService",
		Comment: "新建 Elasticsearch 服务",
		Func:    NewElasticsearchService,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "newDbService",
		Comment: "新建 Db 服务",
		Func:    NewDbService,
	})
}
