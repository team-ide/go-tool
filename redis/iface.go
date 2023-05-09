package redis

import "context"

type IService interface {
	Stop()
	Info(param *Param) (res string, err error)
	Keys(param *Param, pattern string, size int64) (keysResult *KeysResult, err error)
	Expire(param *Param, key string, expire int64) (res bool, err error)
	TTL(param *Param, key string) (res int64, err error)
	Persist(param *Param, key string) (res bool, err error)
	Exists(param *Param, key string) (res int64, err error)
	GetValueInfo(param *Param, key string, valueStart, valueSize int64) (valueInfo *ValueInfo, err error)
	Set(param *Param, key string, value string) (err error)
	SAdd(param *Param, key string, value string) (err error)
	SRem(param *Param, key string, value string) (err error)
	LPush(param *Param, key string, value string) (err error)
	RPush(param *Param, key string, value string) (err error)
	LSet(param *Param, key string, index int64, value string) (err error)
	LRem(param *Param, key string, count int64, value string) (err error)
	HSet(param *Param, key string, field string, value string) (err error)
	HGet(param *Param, key string, field string) (value string, err error)
	HGetAll(param *Param, key string) (value map[string]string, err error)
	HDel(param *Param, key string, field string) (err error)
	Del(param *Param, key string) (count int, err error)
	DelPattern(param *Param, pattern string) (count int, err error)
	SetBit(param *Param, key string, offset int64, value int) (err error)
	BitCount(param *Param, key string) (count int64, err error)
	Get(param *Param, key string) (value string, err error)
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
