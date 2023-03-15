package redis

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/team-ide/go-tool/util"
	"io/ioutil"
	"time"
)

func CreateRedisService(address string, username string, auth string, certPath string) (service *V8Service, err error) {
	service = &V8Service{
		address:  address,
		username: username,
		auth:     auth,
		certPath: certPath,
	}
	err = service.init()
	return
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

type V8Service struct {
	address     string
	auth        string
	username    string
	client      *redis.Client
	lastUseTime int64
	certPath    string
}

func (this_ *V8Service) init() (err error) {
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
		pemCerts, err = ioutil.ReadFile(this_.certPath)
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

	client := redis.NewClient(options)
	this_.client = client
	return
}

func (this_ *V8Service) GetWaitTime() int64 {
	return 10 * 60 * 1000
}

func (this_ *V8Service) GetLastUseTime() int64 {
	return this_.lastUseTime
}

func (this_ *V8Service) SetLastUseTime() {
	this_.lastUseTime = util.GetNowTime()
}
func (this_ *V8Service) Stop() {
	_ = this_.client.Close()
}

func (this_ *V8Service) GetClient(ctx context.Context, database int) (client redis.Cmdable, err error) {
	defer func() {
		this_.lastUseTime = util.GetNowTime()
	}()
	cmd := this_.client.Do(ctx, "select", database)
	_, err = cmd.Result()
	if err != nil {
		return
	}
	client = this_.client
	return
}

func (this_ *V8Service) Info(ctx context.Context) (res string, err error) {

	client, err := this_.GetClient(ctx, 0)
	if err != nil {
		return
	}

	return Info(ctx, client)
}

func (this_ *V8Service) Keys(ctx context.Context, database int, pattern string, size int64) (keysResult *KeysResult, err error) {

	client, err := this_.GetClient(ctx, database)
	if err != nil {
		return
	}

	return Keys(ctx, client, database, pattern, size)
}

func (this_ *V8Service) Expire(ctx context.Context, database int, key string, expire int64) (res bool, err error) {

	client, err := this_.GetClient(ctx, database)
	if err != nil {
		return
	}

	return Expire(ctx, client, key, expire)
}

func (this_ *V8Service) TTL(ctx context.Context, database int, key string) (res int64, err error) {

	client, err := this_.GetClient(ctx, database)
	if err != nil {
		return
	}

	return TTL(ctx, client, key)
}

func (this_ *V8Service) Persist(ctx context.Context, database int, key string) (res bool, err error) {

	client, err := this_.GetClient(ctx, database)
	if err != nil {
		return
	}

	return Persist(ctx, client, key)
}

func (this_ *V8Service) Exists(ctx context.Context, database int, key string) (res int64, err error) {

	client, err := this_.GetClient(ctx, database)
	if err != nil {
		return
	}

	return Exists(ctx, client, key)
}

func (this_ *V8Service) Get(ctx context.Context, database int, key string, valueStart, valueSize int64) (valueInfo *ValueInfo, err error) {

	client, err := this_.GetClient(ctx, database)
	if err != nil {
		return
	}

	return Get(ctx, client, database, key, valueStart, valueSize)
}

func (this_ *V8Service) Set(ctx context.Context, database int, key string, value string) (err error) {

	client, err := this_.GetClient(ctx, database)
	if err != nil {
		return
	}

	return Set(ctx, client, key, value)
}

func (this_ *V8Service) SAdd(ctx context.Context, database int, key string, value string) (err error) {

	client, err := this_.GetClient(ctx, database)
	if err != nil {
		return
	}

	return SAdd(ctx, client, key, value)
}

func (this_ *V8Service) SRem(ctx context.Context, database int, key string, value string) (err error) {

	client, err := this_.GetClient(ctx, database)
	if err != nil {
		return
	}

	return SRem(ctx, client, key, value)
}

func (this_ *V8Service) LPush(ctx context.Context, database int, key string, value string) (err error) {

	client, err := this_.GetClient(ctx, database)
	if err != nil {
		return
	}

	return LPush(ctx, client, key, value)
}

func (this_ *V8Service) RPush(ctx context.Context, database int, key string, value string) (err error) {

	client, err := this_.GetClient(ctx, database)
	if err != nil {
		return
	}

	return RPush(ctx, client, key, value)
}

func (this_ *V8Service) LSet(ctx context.Context, database int, key string, index int64, value string) (err error) {

	client, err := this_.GetClient(ctx, database)
	if err != nil {
		return
	}

	return LSet(ctx, client, key, index, value)
}

func (this_ *V8Service) LRem(ctx context.Context, database int, key string, count int64, value string) (err error) {

	client, err := this_.GetClient(ctx, database)
	if err != nil {
		return
	}

	return LRem(ctx, client, key, count, value)
}

func (this_ *V8Service) HSet(ctx context.Context, database int, key string, field string, value string) (err error) {

	client, err := this_.GetClient(ctx, database)
	if err != nil {
		return
	}

	return HSet(ctx, client, key, field, value)
}

func (this_ *V8Service) HDel(ctx context.Context, database int, key string, field string) (err error) {

	client, err := this_.GetClient(ctx, database)
	if err != nil {
		return
	}

	return HDel(ctx, client, key, field)
}

func (this_ *V8Service) Del(ctx context.Context, database int, key string) (count int, err error) {

	client, err := this_.GetClient(ctx, database)
	if err != nil {
		return
	}

	return Del(ctx, client, key)
}

func (this_ *V8Service) DelPattern(ctx context.Context, database int, pattern string) (count int, err error) {

	client, err := this_.GetClient(ctx, database)
	if err != nil {
		return
	}

	return DelPattern(ctx, client, database, pattern)
}
