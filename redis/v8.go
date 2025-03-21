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

// NewRedisService 创建客户端
func NewRedisService(config *Config) (IService, error) {
	service := &V8Service{
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

type V8Service struct {
	*Config
	client *redis.Client
	*CmdService
	isClosed bool
}

func (this_ *V8Service) init(sshClient *ssh.Client) (err error) {
	options := &redis.Options{
		Addr:         this_.Address,
		DialTimeout:  100 * time.Second,
		ReadTimeout:  100 * time.Second,
		WriteTimeout: 100 * time.Second,
		Password:     this_.Auth,
		Username:     this_.Username,
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

	client := redis.NewClient(options)
	this_.client = client
	this_.isClosed = false
	return
}

func (this_ *V8Service) Close() {
	if this_ == nil {
		return
	}
	if this_.isClosed {
		return
	}
	this_.isClosed = true
	if this_.client != nil {
		_ = this_.client.Close()
	}
}

func formatParam(param *Param) *Param {
	if param == nil {
		param = &Param{}
	}
	if param.Ctx == nil {
		param.Ctx = context.Background()
	}
	return param
}
func (this_ *V8Service) getClient(param *Param) (client redis.Cmdable, err error) {
	param = formatParam(param)
	cmd := this_.client.Do(param.Ctx, "select", param.Database)
	_, err = cmd.Result()
	if err != nil {
		return
	}
	client = this_.client
	return
}

func (this_ *V8Service) GetClient(args ...Arg) (client redis.Cmdable, err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)
	return this_.getClient(param)
}

func (this_ *V8Service) Keys(pattern string, args ...Arg) (keysResult *KeysResult, err error) {
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
	return Keys(param.Ctx, client, param.Database, pattern, size)
}

func (this_ *V8Service) Scan(pattern string, args ...Arg) (keysResult *KeysResult, err error) {
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
	return Scan(param.Ctx, client, param.Database, pattern, size, count)
}

func (this_ *V8Service) ValueType(key string, args ...Arg) (valueType string, err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}
	return ValueType(param.Ctx, client, key)
}

func (this_ *V8Service) GetValueInfo(key string, args ...Arg) (valueInfo *ValueInfo, err error) {
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

func (this_ *V8Service) DelPattern(pattern string, args ...Arg) (count int, err error) {
	argCache := getArgCache(args...)
	param := formatParam(argCache.Param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	count = 0
	keysResult, err := Scan(param.Ctx, client, param.Database, pattern, -1, 1000)
	if err != nil {
		return
	}
	for _, key := range keysResult.KeyList {
		cmd := client.Del(param.Ctx, key)
		_, err = cmd.Result()
		if err == nil {
			count++
		} else {
			return
		}
	}
	return
}

func Keys(ctx context.Context, client redis.Cmdable, database int, pattern string, size int64) (keysResult *KeysResult, err error) {
	keysResult = &KeysResult{
		Database: database,
	}
	var list []string
	cmdKeys := client.Keys(ctx, pattern)
	list, err = cmdKeys.Result()
	if err != nil {
		return
	}
	keysResult.Count = int64(len(list))

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

func Scan(ctx context.Context, client redis.Cmdable, database int, match string, size int64, count int64) (keysResult *KeysResult, err error) {

	keysResult = &KeysResult{
		Database: database,
	}
	cmd := client.Scan(ctx, 0, match, count)
	scanIterator := cmd.Iterator()
	for scanIterator.Next(ctx) {
		keysResult.KeyList = append(keysResult.KeyList, scanIterator.Val())
		keysResult.Count++
		if size > 0 && keysResult.Count >= size {
			break
		}
	}
	err = cmd.Err()
	if err != nil {
		return
	}
	return
}
