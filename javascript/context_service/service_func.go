package context_service

import (
	"github.com/team-ide/go-tool/javascript/context_map"
)

func init() {

	dbModule.Service = &context_map.ServiceInfo{
		Name:    "dbService",
		Comment: "",
		FuncList: []*context_map.FuncInfo{

			{
				Name:    "close",
				Comment: "关闭 db 客户端",
			},
			{
				Name:    "getConfig",
				Comment: "获取 数据库 配置",
			},
			{
				Name:    "getDialect",
				Comment: "获取 方言",
			},
			{
				Name:    "getDb",
				Comment: "获取 sql.DB",
			},
			{
				Name:    "info",
				Comment: "查询数据库信息",
			},
			{
				Name:    "exec",
				Comment: "执行 SQL",
			},
			{
				Name:    "execs",
				Comment: "批量执行 SQL",
			},
			{
				Name:    "count",
				Comment: "统计查询 SQL",
			},
			{
				Name:    "queryOne",
				Comment: "查询 单个 结构体 SQL  需要 传入 结构体 接收",
			},
			{
				Name:    "query",
				Comment: "查询 列表 结构体 SQL  需要 传入 结构体 接收",
			},
			{
				Name:    "queryMap",
				Comment: "查询 列表 map SQL  返回 list < map >",
			},
			{
				Name:    "queryPage",
				Comment: "分页查询 列表 结构体 SQL  需要 传入 结构体 接收",
			},
			{
				Name:    "queryMapPage",
				Comment: "分页查询 列表 map 列表 SQL  返回 list < map >",
			},
			{
				Name:    "ownerCreateSql",
				Comment: "创建 库|模式|用户等 SQL 取决于哪种数据库",
			},
			{
				Name:    "ownersSelect",
				Comment: "查询 库|模式|用户等 取决于哪种数据库",
			},
			{
				Name:    "tablesSelect",
				Comment: "查询 库|模式|用户 下所有表",
			},
			{
				Name:    "tableDetail",
				Comment: "查询 某个表明细",
			},
			{
				Name:    "ownerCreate",
				Comment: "创建 库|模式|用户等 取决于哪种数据库",
			},
			{
				Name:    "ownerDelete",
				Comment: "删除 库|模式|用户等 取决于哪种数据库",
			},
			{
				Name:    "ownerDataTrim",
				Comment: "清空 库|模式|用户等 数据 取决于哪种数据库",
			},
			{
				Name:    "DDL",
				Comment: "查看 库 或 表 SQL语句",
			},
			{
				Name:    "model",
				Comment: "将 表 转换某些语言的模型",
			},
			{
				Name:    "tableCreate",
				Comment: "创建表",
			},
			{
				Name:    "tableDelete",
				Comment: "删除表",
			},
			{
				Name:    "tableDataTrim",
				Comment: "表数据清理",
			},
			{
				Name:    "tableCreateSql",
				Comment: "获取建表 SQL",
			},
			{
				Name:    "tableUpdateSql",
				Comment: "获取更新表的 SQL",
			},
			{
				Name:    "tableUpdate",
				Comment: "更新表",
			},
			{
				Name:    "dataListSql",
				Comment: "将 需要 新增、修改、删除的 数据 转为 SQL",
			},
			{
				Name:    "dataListExec",
				Comment: "执行 需要 新增、修改、删除的 数据",
			},
			{
				Name:    "tableData",
				Comment: "根据 一些 条件 查询 表数据",
			},
			{
				Name:    "getTargetDialect",
				Comment: "获取 目标数据库方言",
			},
			{
				Name:    "executeSQL",
				Comment: "执行某段 SQL",
			},
			{
				Name:    "startExport",
				Comment: "开始 导出",
			},
			{
				Name:    "startImport",
				Comment: "开始 导入",
			},
			{
				Name:    "startSync",
				Comment: "开始 同步",
			},
		},
	}

	elasticsearchModule.Service = &context_map.ServiceInfo{
		Name:    "elasticsearchService",
		Comment: "",
		FuncList: []*context_map.FuncInfo{

			{
				Name:    "close",
				Comment: " 关闭 elasticsearch 客户端",
			},
			{
				Name:    "info",
				Comment: " 获取 elasticsearch 信息",
			},
			{
				Name:    "deleteIndex",
				Comment: " 删除 索引",
			},
			{
				Name:    "createIndex",
				Comment: " 创建 索引",
			},
			{
				Name:    "indexes",
				Comment: " 查询 索引",
			},
			{
				Name:    "getMapping",
				Comment: " 查询 索引 配置",
			},
			{
				Name:    "putMapping",
				Comment: " 设置 索引 配置",
			},
			{
				Name:    "setFieldType",
				Comment: " 设置 索引 字段类型",
			},
			{
				Name:    "search",
				Comment: " 搜索",
			},
			{
				Name:    "insert",
				Comment: " 插入数据 并且 等待刷新",
			},
			{
				Name:    "insertNotWait",
				Comment: " 插入数据 不 等待刷新",
			},
			{
				Name:    "batchInsertNotWait",
				Comment: " 批量插入数据 不 等待刷新",
			},
			{
				Name:    "update",
				Comment: " 更新数据 并且 等待刷新",
			},
			{
				Name:    "updateNotWait",
				Comment: " 更新数据 不 等待刷新",
			},
			{
				Name:    "delete",
				Comment: " 删除 并且 等待刷新",
			},
			{
				Name:    "deleteNotWait",
				Comment: " 删除 不 等待刷新",
			},
			{
				Name:    "reindex",
				Comment: "修改索引名称",
			},
			{
				Name:    "indexStat",
				Comment: "索引状态",
			},
			{
				Name:    "scroll",
				Comment: "滚动查询",
			},
			{
				Name:    "indexAlias",
				Comment: "索引别名",
			},
		},
	}

	kafkaModule.Service = &context_map.ServiceInfo{
		Name:    "kafkaService",
		Comment: "",
		FuncList: []*context_map.FuncInfo{

			{
				Name:    "close",
				Comment: "关闭 kafka 客户端",
			},
			{
				Name:    "info",
				Comment: "查看 kafka 信息",
			},
			{
				Name:    "getTopics",
				Comment: "获取主题",
			},
			{
				Name:    "getTopic",
				Comment: "获取主题",
			},
			{
				Name:    "pull",
				Comment: "拉取消息",
			},
			{
				Name:    "markOffset",
				Comment: "提交 位置",
			},
			{
				Name:    "resetOffset",
				Comment: "重置 位置",
			},
			{
				Name:    "createPartitions",
				Comment: "创建 主题 分区",
			},
			{
				Name:    "createTopic",
				Comment: "创建主题",
			},
			{
				Name:    "deleteTopic",
				Comment: "删除 主题",
			},
			{
				Name:    "deleteConsumerGroup",
				Comment: "删除 某个 消费组",
			},
			{
				Name:    "deleteRecords",
				Comment: "删除 主题 数据",
			},
			{
				Name:    "newSyncProducer",
				Comment: "创建 提供者",
			},
			{
				Name:    "push",
				Comment: "推送",
			},
			{
				Name:    "getOffset",
				Comment: "获取 主题 某个 分区 最新 位置",
			},
			{
				Name:    "partitions",
				Comment: "获取 主题 分区",
			},
			{
				Name:    "listConsumerGroups",
				Comment: "查询 所有 消费组",
			},
			{
				Name:    "describeConsumerGroups",
				Comment: "查询 消费组 明细",
			},
			{
				Name:    "deleteConsumerGroupOffset",
				Comment: "删除 消费组 某个主题 分区",
			},
			{
				Name:    "listConsumerGroupOffsets",
				Comment: "查询 消费组 主题分区 信息",
			},
			{
				Name:    "removeMemberFromConsumerGroup",
				Comment: "删除 消费组 成员",
			},
			{
				Name:    "describeTopics",
				Comment: "主题 元数据",
			},
			{
				Name:    "getClient",
				Comment: "获取 kafka 客户端",
			},
		},
	}

	redisModule.Service = &context_map.ServiceInfo{
		Name:    "redisService",
		Comment: "",
		FuncList: []*context_map.FuncInfo{

			{
				Name:    "close",
				Comment: "关闭 redis 客户端",
			},
			{
				Name:    "getClient",
				Comment: "获取 redis 指令客户端",
			},
			{
				Name:    "info",
				Comment: "获取 redis 信息",
			},
			{
				Name:    "keys",
				Comment: "模糊 搜索 key，如 `xx*` 搜索",
			},
			{
				Name:    "expire",
				Comment: "设置 key 过期时间 让给定键在指定的秒数之后过期",
			},
			{
				Name:    "ttl",
				Comment: "查看给定键距离过期还有多少秒",
			},
			{
				Name:    "persist",
				Comment: "移除键的过期时间",
			},
			{
				Name:    "exists",
				Comment: "判断 key 是否存在",
			},
			{
				Name:    "getValueInfo",
				Comment: "获取 key 的值信息  string、set、list、hash等值 多个值的情况下 使用 StartArg 和 SizeArg 查询一定数量的值",
			},
			{
				Name:    "get",
				Comment: "获取 string 类型的值",
			},
			{
				Name:    "set",
				Comment: "设置 string 类型的值",
			},
			{
				Name:    "setAdd",
				Comment: "在 set 中 添加 string 类型的值",
			},
			{
				Name:    "setRem",
				Comment: "在 set 中 移除 string 类型的值",
			},
			{
				Name:    "setCard",
				Comment: "在 set 中 移除 string 类型的值",
			},
			{
				Name:    "listPush",
				Comment: "在 list 中 往 头部 追加 string 类型的值",
			},
			{
				Name:    "listRPush",
				Comment: "在 list 中 往 尾部 追加 string 类型的值",
			},
			{
				Name:    "listSet",
				Comment: "在 list 中 往 某个 索引位 设置 string 类型的值",
			},
			{
				Name:    "listRem",
				Comment: "在 list 中 根据参数 COUNT 的值，移除列表中与参数 VALUE 相等的元素 count > 0 : 从表头开始向表尾搜索，移除与 VALUE 相等的元素，数量为 COUNT count < 0 : 从表尾开始向表头搜索，移除与 VALUE 相等的元素，数量为 COUNT 的绝对值 count = 0 : 移除表中所有与 VALUE 相等的值",
			},
			{
				Name:    "hashSet",
				Comment: "在 hash 中 设置 字段 的值",
			},
			{
				Name:    "hashGet",
				Comment: "在 hash 中 获取 字段 的值",
			},
			{
				Name:    "hashGetAll",
				Comment: "在 hash 中 获取 所有 字段 的值",
			},
			{
				Name:    "hashDel",
				Comment: "在 hash 中 删除 字段 的值",
			},
			{
				Name:    "del",
				Comment: "删除 某个 key",
			},
			{
				Name:    "delPattern",
				Comment: "模糊删除 匹配 key",
			},
			{
				Name:    "bitSet",
				Comment: "在 bitmap 中 设置 某个 位置的值",
			},
			{
				Name:    "bitCount",
				Comment: "在 bitmap 中 统计 所有 值",
			},
		},
	}

	zookeeperModule.Service = &context_map.ServiceInfo{
		Name:    "zookeeperService",
		Comment: "",
		FuncList: []*context_map.FuncInfo{

			{
				Name:    "close",
				Comment: "关闭 客户端",
			},
			{
				Name:    "getConn",
				Comment: "获取 zk Conn",
			},
			{
				Name:    "info",
				Comment: "查看 zk 相关信息",
			},
			{
				Name:    "create",
				Comment: "创建 永久 节点",
			},
			{
				Name:    "createIfNotExists",
				Comment: "如果不存在 则创建 永久 节点 创建时候如果已存在不报错  如果 父节点不存在 则先创建父节点",
			},
			{
				Name:    "createEphemeral",
				Comment: "创建 临时 节点",
			},
			{
				Name:    "createEphemeralIfNotExists",
				Comment: "如果不存在 则创建 临时 节点 创建时候如果已存在不报错 如果 父节点不存在 则先创建父节点",
			},
			{
				Name:    "exists",
				Comment: "查看节点是否存在",
			},
			{
				Name:    "set",
				Comment: "设置 节点 值",
			},
			{
				Name:    "get",
				Comment: "查看 节点 数据",
			},
			{
				Name:    "getInfo",
				Comment: "查看 节点 信息",
			},
			{
				Name:    "stat",
				Comment: "节点 状态",
			},
			{
				Name:    "getChildren",
				Comment: "查询 子节点",
			},
			{
				Name:    "delete",
				Comment: "删除节点 如果 包含子节点 则先删除所有子节点",
			},
			{
				Name:    "watchChildren",
				Comment: "监听 子节点 只监听当前节点下的子节点 NodeEventError 监听异常 NodeEventStopped zk客户端关闭 NodeEventAdded 节点新增 NodeEventDeleted 节点删除 NodeEventNodeNotFound 节点不存在",
			},
		},
	}

}