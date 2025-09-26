package redis

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"net"
	"sort"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/team-ide/go-tool/util"
	"golang.org/x/crypto/ssh"
)

// NewClusterService 创建集群客户端
func NewClusterService(config *Config) (IService, error) {
	service := &ClusterService{
		Config: config,
		CmdService: &CmdService{
			ThrowNotFoundErr: config.ThrowNotFoundErr,
		},
	}
	service.CmdService.GetClient = service.getClient
	err := service.init(config.SSHClient)
	if err != nil {
		return nil, err
	}
	return service, nil
}

type ClusterService struct {
	*Config
	redisCluster *redis.ClusterClient
	*CmdService
	isClosed bool
}

func (this_ *ClusterService) init(sshClient *ssh.Client) (err error) {
	options := &redis.ClusterOptions{
		Addrs:        this_.Servers,
		DialTimeout:  100 * time.Second,
		ReadTimeout:  100 * time.Second,
		WriteTimeout: 100 * time.Second,
		Username:     this_.Username,
		Password:     this_.Auth,
	}
	TLSClientConfig := &tls.Config{}
	if this_.InsecureSkipVerify {
		TLSClientConfig.InsecureSkipVerify = true
	}
	if this_.CertPath != "" {
		certPool := x509.NewCertPool()
		var pemCerts []byte
		pemCerts, err = util.ReadFile(this_.CertPath)
		if err != nil {
			return
		}

		if !certPool.AppendCertsFromPEM(pemCerts) {
			err = errors.New("证书[" + this_.CertPath + "]解析失败")
			return
		}
		TLSClientConfig.RootCAs = certPool
	}
	if this_.InsecureSkipVerify || this_.CertPath != "" {
		options.TLSConfig = TLSClientConfig
	}

	if sshClient != nil {
		options.Dialer = func(ctx context.Context, network, addr string) (net.Conn, error) {
			conn, e := sshClient.Dial("tcp", addr)
			return &util.SSHChanConn{Conn: conn}, e
		}
	}

	redisCluster := redis.NewClusterClient(options)
	this_.redisCluster = redisCluster
	this_.isClosed = false
	return
}

func (this_ *ClusterService) Close() {
	if this_ == nil {
		return
	}
	if this_.isClosed {
		return
	}
	this_.isClosed = true
	if this_.redisCluster != nil {
		_ = this_.redisCluster.Close()
	}
}

func (this_ *ClusterService) getClient(param *Param) (client redis.Cmdable, err error) {
	param = formatParam(param)
	client = this_.redisCluster
	if param.Ctx != nil && param.Database >= 0 {
		return
	}
	return
}

func (this_ *ClusterService) GetClient(args ...Arg) (client redis.Cmdable, err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)
	return this_.getClient(param)
}

func (this_ *ClusterService) Keys(pattern string, args ...Arg) (keysResult *KeysResult, err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	var size int64 = -1
	if argCache.SizeArg != nil {
		size = argCache.SizeArg.Size
	}

	return ClusterKeys(param.Ctx, client.(*redis.ClusterClient), param.Database, pattern, size)
}

func (this_ *ClusterService) Scan(pattern string, args ...Arg) (keysResult *KeysResult, err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	var size int64 = -1
	if argCache.SizeArg != nil {
		size = argCache.SizeArg.Size
	}
	var count int64 = 10000
	if argCache.CountArg != nil {
		count = argCache.CountArg.Count
	}

	return ClusterScan(param.Ctx, client.(*redis.ClusterClient), param.Database, pattern, size, count)
}

func (this_ *ClusterService) ValueType(key string, args ...Arg) (valueType string, err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}
	return ValueType(param.Ctx, client, key)
}

func (this_ *ClusterService) GetValueInfo(key string, args ...Arg) (valueInfo *ValueInfo, err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}
	var valueStart int64 = -1
	var valueSize int64 = -1
	if argCache.StartArg != nil {
		valueStart = int64(argCache.StartArg.Start)
	}
	if argCache.SizeArg != nil {
		valueSize = int64(argCache.SizeArg.Size)
	}

	return GetValueInfo(param.Ctx, client, param.Database, key, valueStart, valueSize)
}

func (this_ *ClusterService) DelPattern(pattern string, args ...Arg) (count int, err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	keysResult, err := ClusterScan(param.Ctx, client.(*redis.ClusterClient), param.Database, pattern, -1, 1000)
	if err != nil {
		return
	}

	count = 0
	for _, key := range keysResult.KeyList {
		cmd := client.Del(param.Ctx, key)
		_, err = cmd.Result()
		if err == nil {
			count++
		}
	}
	return
}

func ClusterKeys(ctx context.Context, client *redis.ClusterClient, database int, pattern string, size int64) (keysResult *KeysResult, err error) {
	keysResult = &KeysResult{
		Database: database,
	}
	var list []string
	err = client.ForEachMaster(ctx, func(ctx context.Context, client *redis.Client) (e error) {
		_, e = client.Do(ctx, "select", database).Result()
		if e != nil {
			return
		}
		var ls []string
		cmd := client.Keys(ctx, pattern)
		ls, e = cmd.Result()
		if e != nil {
			return
		}
		keysResult.Count += int64(len(ls))
		list = append(list, ls...)
		return
	})
	if err != nil {
		return
	}

	sort.Slice(list, func(i, j int) bool {
		return strings.ToLower(list[i]) < strings.ToLower(list[j]) //升序  即前面的值比后面的小 忽略大小写排序
	})
	if keysResult.Count <= size || size <= 0 {
		keysResult.KeyList = list
	} else {
		keysResult.KeyList = list[0:size]
	}
	return
}

func ClusterScan(ctx context.Context, client *redis.ClusterClient, database int, match string, size int64, count int64) (keysResult *KeysResult, err error) {
	keysResult = &KeysResult{
		Database: database,
	}
	err = client.ForEachMaster(ctx, func(ctx context.Context, client *redis.Client) (e error) {
		if size > 0 && keysResult.Count >= size {
			return
		}

		_, e = client.Do(ctx, "select", database).Result()
		if e != nil {
			return
		}

		cmd := client.Scan(ctx, 0, match, count)
		scanIterator := cmd.Iterator()
		for scanIterator.Next(ctx) {
			key := scanIterator.Val()
			keysResult.KeyList = append(keysResult.KeyList, key)
			keysResult.Count++
			if size > 0 && keysResult.Count >= size {
				break
			}
		}
		e = cmd.Err()
		return
	})
	if err != nil {
		return
	}

	return
}
