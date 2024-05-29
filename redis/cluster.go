package redis

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/team-ide/go-tool/util"
	"golang.org/x/crypto/ssh"
	"net"
	"sort"
	"strings"
	"time"
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
		TLSClientConfig := &tls.Config{
			InsecureSkipVerify: true,
		}
		TLSClientConfig.RootCAs = certPool
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
	return
}

func (this_ *ClusterService) Close() {
	if this_ != nil && this_.redisCluster != nil {
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

	var size = -1
	if argCache.SizeArg != nil {
		size = argCache.SizeArg.Size
	}

	return ClusterKeys(param.Ctx, client.(*redis.ClusterClient), param.Database, pattern, size)
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

	keysResult, err := ClusterKeys(param.Ctx, client.(*redis.ClusterClient), param.Database, pattern, -1)
	if err != nil {
		return
	}

	count = 0
	for _, keyInfo := range keysResult.KeyList {
		cmd := client.Del(param.Ctx, keyInfo.Key)
		_, err = cmd.Result()
		if err == nil {
			count++
		}
	}
	return
}

func ClusterKeys(ctx context.Context, client *redis.ClusterClient, database int, pattern string, size int) (keysResult *KeysResult, err error) {
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

	sort.Slice(list, func(i, j int) bool {
		return strings.ToLower(list[i]) < strings.ToLower(list[j]) //升序  即前面的值比后面的小 忽略大小写排序
	})
	var keys []string
	if keysResult.Count <= size || size < 0 {
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
