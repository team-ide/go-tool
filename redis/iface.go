package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type IService interface {
	Stop()
	GetClient(args ...Arg) (client redis.Cmdable, err error)
	Info(args ...Arg) (res string, err error)
	Keys(pattern string, args ...Arg) (keysResult *KeysResult, err error)
	Expire(key string, expire int64, args ...Arg) (res bool, err error)
	TTL(key string, args ...Arg) (res int64, err error)
	Persist(key string, args ...Arg) (res bool, err error)
	Exists(key string, args ...Arg) (res int64, err error)
	GetValueInfo(key string, args ...Arg) (valueInfo *ValueInfo, err error)

	Get(key string, args ...Arg) (value string, err error)
	Set(key string, value string, args ...Arg) (err error)

	SAdd(key string, value string, args ...Arg) (err error)
	SRem(key string, value string, args ...Arg) (err error)
	SCard(key string, args ...Arg) (res int64, err error)

	LPush(key string, value string, args ...Arg) (err error)
	RPush(key string, value string, args ...Arg) (err error)
	LSet(key string, index int64, value string, args ...Arg) (err error)
	LRem(key string, count int64, value string, args ...Arg) (err error)
	HSet(key string, field string, value string, args ...Arg) (err error)
	HGet(key string, field string, args ...Arg) (value string, err error)
	HGetAll(key string, args ...Arg) (value map[string]string, err error)
	HDel(key string, field string, args ...Arg) (err error)

	Del(key string, args ...Arg) (count int, err error)
	DelPattern(pattern string, args ...Arg) (count int, err error)
	SetBit(key string, offset int64, value int, args ...Arg) (err error)
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
