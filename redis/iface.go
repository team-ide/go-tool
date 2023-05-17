package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type IService interface {
	// Close 关闭 redis 客户端
	Close()
	// GetClient 获取 redis 指令客户端
	GetClient(args ...Arg) (client redis.Cmdable, err error)
	// Info 获取 redis 信息
	Info(args ...Arg) (res string, err error)
	// Keys 模糊 搜索 key，如 `xx*` 搜索
	Keys(pattern string, args ...Arg) (keysResult *KeysResult, err error)
	// Expire 设置 key 过期时间 让给定键在指定的秒数之后过期
	Expire(key string, expire int64, args ...Arg) (res bool, err error)
	// Ttl 查看给定键距离过期还有多少秒
	Ttl(key string, args ...Arg) (res int64, err error)
	// Persist 移除键的过期时间
	Persist(key string, args ...Arg) (res bool, err error)
	// Exists 判断 key 是否存在
	Exists(key string, args ...Arg) (res int64, err error)
	// GetValueInfo 获取 key 的值信息  string、set、list、hash等值 多个值的情况下 使用 StartArg 和 SizeArg 查询一定数量的值
	GetValueInfo(key string, args ...Arg) (valueInfo *ValueInfo, err error)

	// Get 获取 string 类型的值
	Get(key string, args ...Arg) (value string, err error)
	// Set 设置 string 类型的值
	Set(key string, value string, args ...Arg) (err error)

	// SetAdd 在 set 中 添加 string 类型的值
	SetAdd(key string, value string, args ...Arg) (err error)
	// SetRem 在 set 中 移除 string 类型的值
	SetRem(key string, value string, args ...Arg) (err error)
	// SetCard 在 set 中 移除 string 类型的值
	SetCard(key string, args ...Arg) (res int64, err error)

	// ListPush 在 list 中 往 头部 追加 string 类型的值
	ListPush(key string, value string, args ...Arg) (err error)
	// ListRPush 在 list 中 往 尾部 追加 string 类型的值
	ListRPush(key string, value string, args ...Arg) (err error)
	// ListSet 在 list 中 往 某个 索引位 设置 string 类型的值
	ListSet(key string, index int64, value string, args ...Arg) (err error)
	// ListRem 在 list 中 根据参数 COUNT 的值，移除列表中与参数 VALUE 相等的元素 count > 0 : 从表头开始向表尾搜索，移除与 VALUE 相等的元素，数量为 COUNT count < 0 : 从表尾开始向表头搜索，移除与 VALUE 相等的元素，数量为 COUNT 的绝对值 count = 0 : 移除表中所有与 VALUE 相等的值
	ListRem(key string, count int64, value string, args ...Arg) (err error)

	// HashSet 在 hash 中 设置 字段 的值
	HashSet(key string, field string, value string, args ...Arg) (err error)
	// HashGet 在 hash 中 获取 字段 的值
	HashGet(key string, field string, args ...Arg) (value string, err error)
	// HashGetAll 在 hash 中 获取 所有 字段 的值
	HashGetAll(key string, args ...Arg) (value map[string]string, err error)
	// HashDel 在 hash 中 删除 字段 的值
	HashDel(key string, field string, args ...Arg) (err error)

	// Del 删除 某个 key
	Del(key string, args ...Arg) (count int, err error)
	// DelPattern 模糊删除 匹配 key
	DelPattern(pattern string, args ...Arg) (count int, err error)
	// BitSet 在 bitmap 中 设置 某个 位置的值
	BitSet(key string, offset int64, value int, args ...Arg) (err error)
	// BitCount 在 bitmap 中 统计 所有 值
	BitCount(key string, args ...Arg) (count int64, err error)
}

type Arg interface {
	IsArg()
}

type ValueInfo struct {
	Database    int         `json:"database"`
	Key         string      `json:"key"`
	ValueType   string      `json:"valueType"`
	Value       interface{} `json:"value"`
	ValueCount  int64       `json:"valueCount"`
	ValueStart  int64       `json:"valueStart"`
	ValueEnd    int64       `json:"valueEnd"`
	Cursor      uint64      `json:"cursor"`
	MemoryUsage int64       `json:"memoryUsage"`
	TTL         int64       `json:"ttl"`
}
type KeysResult struct {
	Count   int        `json:"count"`
	KeyList []*KeyInfo `json:"keyList"`
}
type KeyInfo struct {
	Database int    `json:"database"`
	Key      string `json:"key"`
}

type Param struct {
	Ctx      context.Context
	Database int
}

func (this_ *Param) IsArg() {}

type SizeArg struct {
	Size int
}

func (this_ *SizeArg) IsArg() {}

func NewSizeArg(size int) *SizeArg {
	return &SizeArg{Size: size}
}

type StartArg struct {
	Start int
}

func (this_ *StartArg) IsArg() {}

func NewStartArg(start int) *StartArg {
	return &StartArg{Start: start}
}

type ArgCache struct {
	Param    *Param
	SizeArg  *SizeArg
	StartArg *StartArg
}

func getArgCache(args ...Arg) (res *ArgCache) {
	res = &ArgCache{}
	for _, arg := range args {
		if arg == nil {
			continue
		}
		switch tV := arg.(type) {
		case *Param:
			res.Param = tV
		case *SizeArg:
			res.SizeArg = tV
		case *StartArg:
			res.StartArg = tV
		}
	}
	return
}
