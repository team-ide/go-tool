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
				Name: "newService",
				Comment: `新建 Redis 服务
redisService = redis.newService({address:"127.0.0.1:6379",auth:""})`,
				Func: redis.NewServiceScript,
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
				Name: "newService",
				Comment: `新建 Zookeeper 服务
zookeeperService = zookeeper.newService({address:"127.0.0.1:2181"})`,
				Func: zookeeper.NewServiceScript,
			},
		},
	}

	kafkaModule = &context_map.ModuleInfo{
		Name:    "kafka",
		Comment: "Kafka 模块",
		FuncList: []*context_map.FuncInfo{
			{
				Name: "newService",
				Comment: `新建 Kafka 服务
kafkaService = kafka.newService({address:"127.0.0.1:9092"})`,
				Func: kafka.NewServiceScript,
			},
		},
	}

	elasticsearchModule = &context_map.ModuleInfo{
		Name:    "elasticsearch",
		Comment: "Elasticsearch 模块",
		FuncList: []*context_map.FuncInfo{
			{
				Name: "newService",
				Comment: `新建 Elasticsearch 服务
elasticsearchService = elasticsearch.newService({url:"http://127.0.0.1:9200"})`,
				Func: elasticsearch.NewServiceScript,
			},
		},
	}

	dbModule = &context_map.ModuleInfo{
		Name:    "db",
		Comment: "Db 模块",
		FuncList: []*context_map.FuncInfo{
			{
				Name: "newService",
				Comment: `新建 Db 服务
dbService = db.newService({host:"127.0.0.1", port: 3306, username: "root", password: "123456", database: "test_db"})`,
				Func: db.NewServiceScript,
			},
			{
				Name:    "toConfig",
				Comment: `任意对象转为 Config 用于数据库连接等`,
				Func:    db.ToConfig,
			},
			{
				Name:    "toOwnerModel",
				Comment: `任意对象转为 OwnerModel 用于创建 数据库等`,
				Func:    db.ToOwnerModel,
			},
			{
				Name:    "toTableModel",
				Comment: `任意对象转为 TableModel 用于创建 表等`,
				Func:    db.ToTableModel,
			},
			{
				Name:    "toColumnModel",
				Comment: `任意对象转为 ColumnModel 用于创建 字段等`,
				Func:    db.ToColumnModel,
			},
			{
				Name:    "toIndexModel",
				Comment: "任意对象转为 IndexModel 用于创建 索引等",
				Func:    db.ToIndexModel,
			},
			{
				Name:    "toPrimaryKeyModel",
				Comment: "任意对象转为 PrimaryKeyModel 用于创建 主键等",
				Func:    db.ToPrimaryKeyModel,
			},
			{
				Name:    "toParam",
				Comment: "任意对象转为 Param 用于 建库、建表 配置参数等",
				Func:    db.ToParam,
			},
			{
				Name:    "toPage",
				Comment: "任意对象转为 Page 用于 分页查询等",
				Func:    db.ToPage,
			},
			{
				Name:    "toTaskExportParam",
				Comment: "任意对象转为 TaskExportParam 用于 导出任务",
				Func:    db.ToTaskExportParam,
			},
			{
				Name:    "toTaskImportParam",
				Comment: "任意对象转为 TaskImportParam 用于 导入任务",
				Func:    db.ToTaskImportParam,
			},
			{
				Name:    "toTaskSyncParam",
				Comment: "任意对象转为 TaskSyncParam 用于 同步任务",
				Func:    db.ToTaskSyncParam,
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
