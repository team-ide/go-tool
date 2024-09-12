package util

import (
	"fmt"
	"hash/crc32"
	"strings"
	"sync"
	"time"
)

// CacheItem 缓存项结构
type CacheItem[T any] struct {
	value    T
	expireAt int64 // 过期时间 毫秒
}

// Cache 临时缓存结构
type Cache[T any] struct {
	keys      []string
	items     map[string]*CacheItem[T]
	itemsLock sync.Locker
	size      int          // 当前缓存大小
	limit     int          // 缓存大小限制
	ttl       int64        // 缓存项存活时间 毫秒
	ticker    *time.Ticker // 用于定期清理的ticker
	outLog    bool
}

// NewCache 创建 默认 缓存
func NewCache[T any](obj T) *Cache[T] {
	return NewCacheByOptions(nil, obj)
}

type CacheOptions struct {
	Limit  int   // 缓存大小限制
	Ttl    int64 // 缓存项存活时间 毫秒
	OutLog bool
}

// NewCacheByOptions 创建一个 带策略的 缓存
func NewCacheByOptions[T any](opts *CacheOptions, obj T) *Cache[T] {
	if opts == nil {
		opts = &CacheOptions{}
	}
	tc := &Cache[T]{
		items:     make(map[string]*CacheItem[T]),
		itemsLock: &sync.Mutex{},
		size:      0,
		limit:     opts.Limit,
		ttl:       opts.Ttl,
	}
	if opts.Limit > 0 || opts.Ttl > 0 {
		// 每分钟检测一次
		tc.ticker = time.NewTicker(time.Minute)
		go tc.cleanup() // 启动清理协程
	}
	return tc
}

// Get 从缓存中获取一个项
func (this_ *Cache[T]) Get(key string) (res T, find bool) {
	this_.itemsLock.Lock()
	item, find := this_.items[key]
	this_.itemsLock.Unlock()
	if !find {
		return
	}
	res = item.value
	return
}

// GetAnRemove 从缓存中获取一个项
func (this_ *Cache[T]) GetAnRemove(key string) (res T, find bool) {
	this_.itemsLock.Lock()
	item, find := this_.items[key]
	this_.itemsLock.Unlock()
	if !find {
		return
	}
	delete(this_.items, key)
	res = item.value
	return
}

// GetOrLoad 从缓存中获取一个项 如果没有 则调用 load 加载
func (this_ *Cache[T]) GetOrLoad(key string, load func() (T, error)) (res T, err error) {
	this_.itemsLock.Lock()
	item, find := this_.items[key]
	this_.itemsLock.Unlock()
	if find {
		res = item.value
		return
	}
	if this_.outLog {
		Logger.Debug("缓存不存在 key [" + key + "] ，执行加载")
	}
	value, err := load()
	if err != nil {
		return value, err
	}
	this_.Set(key, value)
	return value, err
}

// Set 从缓存中获取一个项 如果没有 则调用 load 加载
func (this_ *Cache[T]) Set(key string, value T) {
	this_.itemsLock.Lock()
	defer this_.itemsLock.Unlock()

	this_.items[key] = &CacheItem[T]{
		value:    value,
		expireAt: time.Now().UnixMilli() + this_.ttl,
	}
	this_.keys = append(this_.keys, key)
	this_.size++
}

func (this_ *Cache[T]) Put(key string, v T) {
	this_.Set(key, v)
}

// PutIfAbsent 不存在则新增
func (this_ *Cache[T]) PutIfAbsent(key string, v T) {
	this_.itemsLock.Lock()
	_, find := this_.items[key]
	this_.itemsLock.Unlock()
	if !find {
		this_.Set(key, v)
	}
}

// Delete 删除
func (this_ *Cache[T]) Delete(keys ...string) {
	this_.Remove(keys...)
}

// Clear 清理
func (this_ *Cache[T]) Clear() {
	this_.itemsLock.Lock()
	defer this_.itemsLock.Unlock()

	this_.keys = []string{}
	this_.size = 0
	this_.items = make(map[string]*CacheItem[T])

}

// Remove 清理 缓存
func (this_ *Cache[T]) Remove(removeKeys ...string) {
	if len(removeKeys) == 0 {
		return
	}

	this_.itemsLock.Lock()
	defer this_.itemsLock.Unlock()

	removeKeyStr := "," + strings.Join(removeKeys, ",") + ","
	var keys []string
	for _, k := range this_.keys {
		if strings.Contains(removeKeyStr, ","+k+",") {

			if this_.outLog {
				Logger.Debug("删除 keys 中的 key [" + k + "]")
			}
			continue
		}
		keys = append(keys, k)
	}
	for _, removeKey := range removeKeys {
		if this_.outLog {
			Logger.Debug("删除 缓存 中的 key [" + removeKey + "]")
		}
		delete(this_.items, removeKey)
	}
	this_.keys = keys
	this_.size = len(keys)
	if this_.outLog {
		Logger.Debug(fmt.Sprintf("删除 缓存 后 缓存大小 [%d]", this_.size))
	}
}

// cleanLimit 清理超出 limit 的缓存
func (this_ *Cache[T]) cleanLimit() {
	if this_.limit <= 0 {
		return
	}
	this_.itemsLock.Lock()

	var removeKeys []string
	keySize := len(this_.keys)
	removeKeySize := keySize - this_.limit
	if removeKeySize > 0 {
		if this_.outLog {
			Logger.Debug(fmt.Sprintf("缓存数量 [%d] 超出最大数量 [%d]", keySize, this_.limit))
		}
		removeKeys = this_.keys[0:removeKeySize]
	}
	this_.itemsLock.Unlock()

	if len(removeKeys) > 0 {
		this_.Remove(removeKeys...)
	}
}

// cleanup 定期清理过期的缓存项
func (this_ *Cache[T]) cleanup() {
	for {
		select {
		case <-this_.ticker.C:

			var removeKeys []string
			var now = time.Now().UnixMilli()

			func() {
				this_.itemsLock.Lock()
				defer this_.itemsLock.Unlock()
				for key, item := range this_.items {
					// 过期时间 小于当前时间 则清理缓存
					if item.expireAt <= now {
						//vlog.Debug("缓存过期的 key [%s]", key)
						removeKeys = append(removeKeys, key)
					}
				}
			}()

			if this_.outLog {
				Logger.Debug(fmt.Sprintf("缓存过期的 key %s", removeKeys))
			}

			if len(removeKeys) > 0 {
				this_.Remove(removeKeys...)
			}
			if this_.limit > 0 && this_.size > this_.limit {
				this_.cleanLimit()
			}
		}
	}
}

type CacheGroup[T any] struct {
	obj       T
	caches    []*Cache[T]
	cacheSize uint32
	opts      *CacheOptions
}

func NewCacheGroup[T any](cacheSize int, obj T) *CacheGroup[T] {
	g := &CacheGroup[T]{
		obj:       obj,
		cacheSize: uint32(cacheSize),
	}
	g.init()
	return g
}

func NewCacheGroupByOption[T any](cacheSize int, opts *CacheOptions, obj T) *CacheGroup[T] {
	g := &CacheGroup[T]{
		obj:       obj,
		cacheSize: uint32(cacheSize),
		opts:      opts,
	}
	g.init()
	return g
}

func (this_ *CacheGroup[T]) init() {
	var caches []*Cache[T]
	var i uint32
	for i = 0; i < this_.cacheSize; i++ {
		var cache = NewCacheByOptions(this_.opts, this_.obj)
		caches = append(caches, cache)
	}
	this_.caches = caches
}

func (this_ *CacheGroup[T]) GetStringHashCode(str string) uint32 {
	//h := fnv.New32a()
	//_, _ = h.Write([]byte(str))
	//return h.Sum32()
	return crc32.ChecksumIEEE([]byte(str))
}

func (this_ *CacheGroup[T]) GetCache(key string) *Cache[T] {
	hashCode := this_.GetStringHashCode(key)
	return this_.GetCacheByHash(hashCode)
}

func (this_ *CacheGroup[T]) GetCacheByHash(hashCode uint32) *Cache[T] {
	index := hashCode % this_.cacheSize
	return this_.caches[index]
}

func (this_ *CacheGroup[T]) Get(key string) (T, bool) {
	return this_.GetCache(key).Get(key)
}

// Put 设置值
func (this_ *CacheGroup[T]) Put(key string, v T) {
	this_.GetCache(key).Put(key, v)
}

// PutIfAbsent 不存在则新增
func (this_ *CacheGroup[T]) PutIfAbsent(key string, v T) {
	this_.GetCache(key).PutIfAbsent(key, v)
}

// Delete 删除
func (this_ *CacheGroup[T]) Delete(keys ...string) {
	ks := map[uint32][]string{}
	for _, key := range keys {
		c := this_.GetStringHashCode(key)
		ks[c] = append(ks[c], key)
	}
	for c, k := range ks {
		this_.GetCacheByHash(c).Delete(k...)
	}
}

// Clear 清理
func (this_ *CacheGroup[T]) Clear() {
	this_.init()
}
