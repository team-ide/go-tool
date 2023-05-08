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
	"time"
)

// NewRedisService 创建客户端
func NewRedisService(address string, username string, auth string, certPath string) (IService, error) {
	service := &V8Service{
		address:  address,
		username: username,
		auth:     auth,
		certPath: certPath,
	}
	err := service.init(nil)
	if err != nil {
		return nil, err
	}
	return service, nil
}

// NewRedisServiceForSSH 创建 SSH 客户端
func NewRedisServiceForSSH(address string, username string, auth string, certPath string, sshClient *ssh.Client) (IService, error) {
	service := &V8Service{
		address:  address,
		username: username,
		auth:     auth,
		certPath: certPath,
	}
	err := service.init(sshClient)
	if err != nil {
		return nil, err
	}
	return service, nil
}

type V8Service struct {
	address  string
	auth     string
	username string
	client   *redis.Client
	certPath string
}

func (this_ *V8Service) init(sshClient *ssh.Client) (err error) {
	options := &redis.Options{
		Addr:         this_.address,
		DialTimeout:  100 * time.Second,
		ReadTimeout:  100 * time.Second,
		WriteTimeout: 100 * time.Second,
		Password:     this_.auth,
		Username:     this_.username,
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
	if sshClient != nil {
		options.Dialer = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return sshClient.Dial(network, addr)
		}
	}

	client := redis.NewClient(options)
	this_.client = client
	return
}

func (this_ *V8Service) Stop() {
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
func (this_ *V8Service) GetClient(param *Param) (client redis.Cmdable, err error) {
	param = formatParam(param)
	cmd := this_.client.Do(param.Ctx, "select", param.Database)
	_, err = cmd.Result()
	if err != nil {
		return
	}
	client = this_.client
	return
}

func (this_ *V8Service) Info(param *Param) (res string, err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return Info(param.Ctx, client)
}

func (this_ *V8Service) Keys(param *Param, pattern string, size int64) (keysResult *KeysResult, err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return Keys(param.Ctx, client, param.Database, pattern, size)
}

func (this_ *V8Service) Expire(param *Param, key string, expire int64) (res bool, err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return Expire(param.Ctx, client, key, expire)
}

func (this_ *V8Service) TTL(param *Param, key string) (res int64, err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return TTL(param.Ctx, client, key)
}

func (this_ *V8Service) Persist(param *Param, key string) (res bool, err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return Persist(param.Ctx, client, key)
}

func (this_ *V8Service) Exists(param *Param, key string) (res int64, err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return Exists(param.Ctx, client, key)
}

func (this_ *V8Service) GetValueInfo(param *Param, key string, valueStart, valueSize int64) (valueInfo *ValueInfo, err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return GetValueInfo(param.Ctx, client, param.Database, key, valueStart, valueSize)
}

func (this_ *V8Service) Set(param *Param, key string, value string) (err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return Set(param.Ctx, client, key, value)
}

func (this_ *V8Service) SAdd(param *Param, key string, value string) (err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return SAdd(param.Ctx, client, key, value)
}

func (this_ *V8Service) SRem(param *Param, key string, value string) (err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return SRem(param.Ctx, client, key, value)
}

func (this_ *V8Service) LPush(param *Param, key string, value string) (err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return LPush(param.Ctx, client, key, value)
}

func (this_ *V8Service) RPush(param *Param, key string, value string) (err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return RPush(param.Ctx, client, key, value)
}

func (this_ *V8Service) LSet(param *Param, key string, index int64, value string) (err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return LSet(param.Ctx, client, key, index, value)
}

func (this_ *V8Service) LRem(param *Param, key string, count int64, value string) (err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return LRem(param.Ctx, client, key, count, value)
}

func (this_ *V8Service) HSet(param *Param, key string, field string, value string) (err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return HSet(param.Ctx, client, key, field, value)
}

func (this_ *V8Service) HGet(param *Param, key string, field string) (value string, notFound bool, err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return HGet(param.Ctx, client, key, field)
}

func (this_ *V8Service) HGetAll(param *Param, key string) (value map[string]string, notFound bool, err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return HGetAll(param.Ctx, client, key)
}

func (this_ *V8Service) Get(param *Param, key string) (value string, notFound bool, err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return Get(param.Ctx, client, key)
}

func (this_ *V8Service) SetBit(param *Param, key string, offset int64, value int) (err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return SetBit(param.Ctx, client, key, offset, value)
}

func (this_ *V8Service) BitCount(param *Param, key string) (count int64, err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return BitCount(param.Ctx, client, key)
}

func (this_ *V8Service) HDel(param *Param, key string, field string) (err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return HDel(param.Ctx, client, key, field)
}

func (this_ *V8Service) Del(param *Param, key string) (count int, err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return Del(param.Ctx, client, key)
}

func (this_ *V8Service) DelPattern(param *Param, pattern string) (count int, err error) {
	param = formatParam(param)

	client, err := this_.GetClient(param)
	if err != nil {
		return
	}

	return DelPattern(param.Ctx, client, param.Database, pattern)
}
