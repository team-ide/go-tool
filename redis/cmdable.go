package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"time"
)

func ValueType(ctx context.Context, client redis.Cmdable, key string) (ValueType string, err error) {

	cmdType := client.Type(ctx, key)
	ValueType, err = cmdType.Result()
	if errors.Is(err, redis.Nil) {
		err = nil
		return
	}
	return
}

func MemoryUsage(ctx context.Context, client redis.Cmdable, key string) (size int64, err error) {
	cmd := client.MemoryUsage(ctx, key)
	size, err = cmd.Result()
	return
}

func GetValueInfo(ctx context.Context, client redis.Cmdable, database int, key string, valueStart, valueSize int64) (valueInfo *ValueInfo, err error) {
	var valueType string
	valueType, err = ValueType(ctx, client, key)
	if err != nil {
		return
	}
	valueInfo = &ValueInfo{
		Key:      key,
		Database: database,
	}
	MemoryUsageCmd := client.MemoryUsage(ctx, key)

	valueInfo.MemoryUsage, _ = MemoryUsageCmd.Result()

	TTLCmd := client.TTL(ctx, key)
	ttlV, _ := TTLCmd.Result()
	if ttlV > 0 {
		valueInfo.TTL = int64(ttlV / time.Second)
	}
	var value interface{}

	if valueType == "none" {

	} else if valueType == "string" {
		cmd := client.Get(ctx, key)
		value, err = cmd.Result()
		if err != nil {
			util.Logger.Error("Get Error", zap.Any("key", key), zap.Error(err))
			return
		}
	} else if valueType == "list" {

		cmd := client.LLen(ctx, key)

		valueInfo.ValueCount, err = cmd.Result()
		if err != nil {
			util.Logger.Error("LLen Error", zap.Any("key", key), zap.Error(err))
			return
		}
		valueInfo.ValueStart = valueStart
		valueInfo.ValueEnd = valueInfo.ValueStart + valueSize
		if valueSize <= 0 {
			valueInfo.ValueEnd = valueInfo.ValueCount
		}

		var list []string
		cmdRange := client.LRange(ctx, key, valueInfo.ValueStart, valueInfo.ValueEnd)
		list, err = cmdRange.Result()
		if err != nil {
			util.Logger.Error("LRange Error", zap.Any("key", key), zap.Error(err))
			return
		}

		if int64(len(list)) <= valueSize || valueSize <= 0 {
			value = list
		} else {
			value = list[0:valueSize]
		}

	} else if valueType == "set" {

		cmdSCard := client.SCard(ctx, key)
		valueInfo.ValueCount, err = cmdSCard.Result()
		if err != nil {
			util.Logger.Error("SCard Error", zap.Any("key", key), zap.Error(err))
			return
		}
		valueInfo.ValueStart = valueStart
		valueInfo.ValueEnd = valueInfo.ValueStart + valueSize
		if valueSize <= 0 {
			valueInfo.ValueEnd = valueInfo.ValueCount
		}

		var list []string
		cmd := client.SScan(ctx, key, uint64(valueInfo.ValueStart), "*", valueInfo.ValueEnd)
		list, valueInfo.Cursor, err = cmd.Result()
		if err != nil {
			util.Logger.Error("SScan Error", zap.Any("key", key), zap.Error(err))
			return
		}

		if int64(len(list)) <= valueSize || valueSize <= 0 {
			value = list
		} else {
			value = list[0:valueSize]
		}
	} else if valueType == "hash" {

		cmdHLen := client.HLen(ctx, key)
		valueInfo.ValueCount, err = cmdHLen.Result()
		if err != nil {
			util.Logger.Error("HLen Error", zap.Any("key", key), zap.Error(err))
			return
		}
		if valueSize <= 0 {
			valueSize = valueInfo.ValueCount
		}
		valueInfo.ValueStart = valueStart * 2
		valueInfo.ValueEnd = valueInfo.ValueStart + valueSize*2

		cmdHScan := client.HScan(ctx, key, uint64(valueInfo.ValueStart), "*", valueInfo.ValueEnd)

		var keyValueList []string
		keyValueList, valueInfo.Cursor, err = cmdHScan.Result()
		if err != nil {
			util.Logger.Error("HScan Error", zap.Any("key", key), zap.Error(err))
			return
		}
		var keyValueListSize = int64(len(keyValueList))
		var keyValue = map[string]string{}
		for i := int64(0); i < valueSize*2; i++ {
			if i >= keyValueListSize {
				break
			}
			filed := keyValueList[i]
			filedValue := ""
			if i+1 < keyValueListSize {
				filedValue = keyValueList[i+1]
			}
			keyValue[filed] = filedValue
			i++
		}

		value = keyValue
	} else {
		util.Logger.Warn("valueType not support", zap.Any("valueType", valueType), zap.Any("key", key))
	}
	valueInfo.ValueType = valueType
	valueInfo.Value = value

	return
}

type CmdService struct {
	GetClient func(param *Param) (client redis.Cmdable, err error) `json:"-"`
}

func (this_ *CmdService) Info(args ...Arg) (res string, err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	cmd := client.Info(param.Ctx)
	res, err = cmd.Result()
	return
}

// Exists key是否存在
func (this_ *CmdService) Exists(key string, args ...Arg) (res int64, err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	cmdExists := client.Exists(param.Ctx, key)
	res, err = cmdExists.Result()
	return
}

// Expire 让给定键在指定的秒数之后过期
func (this_ *CmdService) Expire(key string, expire int64, args ...Arg) (res bool, err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	cmd := client.Expire(param.Ctx, key, time.Duration(expire)*time.Second)
	res, err = cmd.Result()
	if err == redis.Nil {
		err = nil
		return
	}
	return
}

// Persist 移除键的过期时间
func (this_ *CmdService) Persist(key string, args ...Arg) (res bool, err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	cmd := client.Persist(param.Ctx, key)
	res, err = cmd.Result()
	if err == redis.Nil {
		err = nil
		return
	}
	return
}

// Ttl 查看给定键距离过期还有多少秒
func (this_ *CmdService) Ttl(key string, args ...Arg) (res int64, err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	cmd := client.TTL(param.Ctx, key)
	r, err := cmd.Result()
	if err == redis.Nil {
		err = nil
		return
	}
	if err != nil {
		return
	}
	if r > 0 {
		res = int64(r / time.Second)
	}
	return
}

func (this_ *CmdService) Get(key string, args ...Arg) (value string, err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	cmd := client.Get(param.Ctx, key)
	value, err = cmd.Result()
	if err == redis.Nil {
		err = nil
		//notFound = true
		return
	}
	return
}

func (this_ *CmdService) Set(key string, value string, args ...Arg) (err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	cmd := client.Set(param.Ctx, key, value, time.Duration(0))
	_, err = cmd.Result()
	return
}

func (this_ *CmdService) SetAdd(key string, value string, args ...Arg) (err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	cmd := client.SAdd(param.Ctx, key, value)
	_, err = cmd.Result()
	return
}

func (this_ *CmdService) SetRem(key string, value string, args ...Arg) (err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	cmd := client.SRem(param.Ctx, key, value)
	_, err = cmd.Result()
	return
}

func (this_ *CmdService) SetCard(key string, args ...Arg) (res int64, err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	cmd := client.SCard(param.Ctx, key)
	res, err = cmd.Result()
	return
}

func (this_ *CmdService) ListPush(key string, value string, args ...Arg) (err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	cmd := client.LPush(param.Ctx, key, value)
	_, err = cmd.Result()
	return
}

func (this_ *CmdService) ListSet(key string, index int64, value string, args ...Arg) (err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	cmd := client.LSet(param.Ctx, key, index, value)
	_, err = cmd.Result()
	return
}

func (this_ *CmdService) ListRem(key string, count int64, value string, args ...Arg) (err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	cmd := client.LRem(param.Ctx, key, count, value)
	_, err = cmd.Result()
	return
}

func (this_ *CmdService) ListRPush(key string, value string, args ...Arg) (err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	cmd := client.RPush(param.Ctx, key, value)
	_, err = cmd.Result()
	return
}

func (this_ *CmdService) HashSet(key string, field string, value string, args ...Arg) (err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	cmd := client.HSet(param.Ctx, key, field, value)
	_, err = cmd.Result()
	return
}

func (this_ *CmdService) HashGet(key string, field string, args ...Arg) (value string, err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	cmd := client.HGet(param.Ctx, key, field)
	value, err = cmd.Result()
	if err == redis.Nil {
		err = nil
		//notFound = true
		return
	}
	return
}

func (this_ *CmdService) ScriptRun(src string, KEYS []string, ARGV interface{}, args ...Arg) (value interface{}, err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	script := redis.NewScript(src)
	ret, err := script.Run(param.Ctx, client, KEYS, ARGV).Result()
	if err != nil {
		return
	}
	value = ret
	return
}

func (this_ *CmdService) HashGetAll(key string, args ...Arg) (value map[string]string, err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	cmd := client.HGetAll(param.Ctx, key)
	value, err = cmd.Result()
	if err == redis.Nil {
		err = nil
		//notFound = true
		return
	}
	return
}

func (this_ *CmdService) HashDel(key string, field string, args ...Arg) (err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	cmd := client.HDel(param.Ctx, key, field)
	_, err = cmd.Result()
	return
}

func (this_ *CmdService) BitSet(key string, offset int64, value int, args ...Arg) (err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	cmd := client.SetBit(param.Ctx, key, offset, value)
	err = cmd.Err()
	return
}

func (this_ *CmdService) BitCount(key string, args ...Arg) (count int64, err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	cmd := client.BitCount(param.Ctx, key, nil)
	count, err = cmd.Result()
	return
}

func (this_ *CmdService) Del(key string, args ...Arg) (count int, err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	cmd := client.Del(param.Ctx, key)
	_, err = cmd.Result()
	if err == nil {
		count++
	}
	return
}
