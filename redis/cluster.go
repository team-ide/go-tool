package redis

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/team-ide/go-tool/util"
	"sort"
	"time"
)

// NewClusterService 创建集群客户端
func NewClusterService(servers []string, username string, auth string, certPath string) (IService, error) {
	service := &ClusterService{
		servers:  servers,
		username: username,
		auth:     auth,
		certPath: certPath,
	}
	err := service.init()
	if err != nil {
		return nil, err
	}
	return service, nil
}

type ClusterService struct {
	servers      []string
	auth         string
	username     string
	certPath     string
	redisCluster *redis.ClusterClient
}

func (this_ *ClusterService) init() (err error) {
	options := &redis.ClusterOptions{
		Addrs:        this_.servers,
		DialTimeout:  100 * time.Second,
		ReadTimeout:  100 * time.Second,
		WriteTimeout: 100 * time.Second,
		Username:     this_.username,
		Password:     this_.auth,
	}
	if this_.certPath != "" {
		certPool := x509.NewCertPool()
		var pemCerts []byte
		pemCerts, err = util.ReadFile(this_.certPath)
		if err != nil {
			return
		}

		if !certPool.AppendCertsFromPEM(pemCerts) {
			err = errors.New("证书[" + this_.certPath + "]解析失败")
			return
		}
		TLSClientConfig := &tls.Config{
			InsecureSkipVerify: true,
		}
		TLSClientConfig.RootCAs = certPool
		options.TLSConfig = TLSClientConfig
	}

	redisCluster := redis.NewClusterClient(options)
	this_.redisCluster = redisCluster
	return
}

func (this_ *ClusterService) Stop() {
	_ = this_.redisCluster.Close()
}

func (this_ *ClusterService) GetClient(param *Param) (client redis.Cmdable, err error) {
	param = formatParam(param)
	client = this_.redisCluster
	if param.Ctx != nil && param.Database >= 0 {
		return
	}
	return
}

func (this_ *ClusterService) Info(param *Param) (res string, err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return Info(param.Ctx, client)
}

func (this_ *ClusterService) Keys(param *Param, pattern string, size int64) (keysResult *KeysResult, err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return ClusterKeys(param.Ctx, client.(*redis.ClusterClient), param.Database, pattern, size)
}

func (this_ *ClusterService) Expire(param *Param, key string, expire int64) (res bool, err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return Expire(param.Ctx, client, key, expire)
}

func (this_ *ClusterService) TTL(param *Param, key string) (res int64, err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return TTL(param.Ctx, client, key)
}

func (this_ *ClusterService) Persist(param *Param, key string) (res bool, err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return Persist(param.Ctx, client, key)
}

func (this_ *ClusterService) Exists(param *Param, key string) (res int64, err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return Exists(param.Ctx, client, key)
}

func (this_ *ClusterService) GetValueInfo(param *Param, key string, valueStart, valueSize int64) (valueInfo *ValueInfo, err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return GetValueInfo(param.Ctx, client, param.Database, key, valueStart, valueSize)
}

func (this_ *ClusterService) Set(param *Param, key string, value string) (err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return Set(param.Ctx, client, key, value)
}

func (this_ *ClusterService) SAdd(param *Param, key string, value string) (err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return SAdd(param.Ctx, client, key, value)
}

func (this_ *ClusterService) SRem(param *Param, key string, value string) (err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return SRem(param.Ctx, client, key, value)
}

func (this_ *ClusterService) LPush(param *Param, key string, value string) (err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return LPush(param.Ctx, client, key, value)
}

func (this_ *ClusterService) RPush(param *Param, key string, value string) (err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return RPush(param.Ctx, client, key, value)
}

func (this_ *ClusterService) LSet(param *Param, key string, index int64, value string) (err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return LSet(param.Ctx, client, key, index, value)
}

func (this_ *ClusterService) LRem(param *Param, key string, count int64, value string) (err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return LRem(param.Ctx, client, key, count, value)
}

func (this_ *ClusterService) HSet(param *Param, key string, field string, value string) (err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return HSet(param.Ctx, client, key, field, value)
}

func (this_ *ClusterService) HGet(param *Param, key string, field string) (value string, notFound bool, err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return HGet(param.Ctx, client, key, field)
}

func (this_ *ClusterService) Get(param *Param, key string) (value string, notFound bool, err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return Get(param.Ctx, client, key)
}

func (this_ *ClusterService) SetBit(param *Param, key string, offset int64, value int) (err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return SetBit(param.Ctx, client, key, offset, value)
}

func (this_ *ClusterService) BitCount(param *Param, key string) (count int64, err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return BitCount(param.Ctx, client, key)
}

func (this_ *ClusterService) HDel(param *Param, key string, field string) (err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return HDel(param.Ctx, client, key, field)
}

func (this_ *ClusterService) Del(param *Param, key string) (count int, err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return Del(param.Ctx, client, key)
}

func (this_ *ClusterService) DelPattern(param *Param, pattern string) (count int, err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	keysResult, err := ClusterKeys(param.Ctx, client.(*redis.ClusterClient), param.Database, pattern, -1)
	if err != nil {
		return
	}

	count = 0
	for _, keyInfo := range keysResult.KeyList {
		_, err = Del(param.Ctx, client, keyInfo.Key)
		if err == nil {
			count++
		}
	}
	return
}

func ClusterKeys(ctx context.Context, client *redis.ClusterClient, database int, pattern string, size int64) (keysResult *KeysResult, err error) {
	keysResult = &KeysResult{}
	var list []string
	err = client.ForEachMaster(ctx, func(ctx context.Context, client *redis.Client) (err error) {

		var ls []string
		cmd := client.Keys(ctx, pattern)
		ls, err = cmd.Result()
		if err != nil {
			return
		}
		keysResult.Count += len(ls)
		list = append(list, ls...)
		return
	})
	sor := sort.StringSlice(list)
	sor.Sort()
	var keys []string
	if int64(keysResult.Count) <= size || size < 0 {
		keys = list
	} else {
		keys = list[0:size]
	}
	for _, key := range keys {
		info := &KeyInfo{
			Key:      key,
			Database: database,
		}
		keysResult.KeyList = append(keysResult.KeyList, info)
	}
	return
}
