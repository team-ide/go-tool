package util

import (
	"hash/fnv"
	"sync"
)

type Cache struct {
	data map[string]interface{}
	lock sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func (this_ *Cache) Get(key string) (interface{}, bool) {
	this_.lock.Lock()
	v, find := this_.data[key]
	this_.lock.Unlock()
	return v, find
}

// Put 设置值
func (this_ *Cache) Put(key string, v interface{}) {
	this_.lock.Lock()
	this_.data[key] = v
	this_.lock.Unlock()
}

// PutIfAbsent 不存在则新增
func (this_ *Cache) PutIfAbsent(key string, v interface{}) {
	this_.lock.Lock()
	_, find := this_.data[key]
	if !find {
		this_.data[key] = v
	}
	this_.lock.Unlock()
}

// Delete 删除
func (this_ *Cache) Delete(keys ...string) {
	this_.lock.Lock()
	for _, key := range keys {
		delete(this_.data, key)
	}
	this_.lock.Unlock()
}

// Clear 清理
func (this_ *Cache) Clear() {
	this_.lock.Lock()
	this_.data = make(map[string]interface{})
	this_.lock.Unlock()
}

type CacheGroup struct {
	caches    []*Cache
	cacheSize uint32
}

func NewCacheGroup(cacheSize int) *CacheGroup {
	g := &CacheGroup{
		cacheSize: uint32(cacheSize),
	}
	g.init()
	return g
}

func (this_ *CacheGroup) init() {
	var caches []*Cache
	var i uint32
	for i = 0; i < this_.cacheSize; i++ {
		var cache = NewCache()
		caches = append(caches, cache)
	}
	this_.caches = caches
}
func (this_ *CacheGroup) GetStringHashCode(str string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(str))
	return h.Sum32()
}
func (this_ *CacheGroup) GetCache(key string) *Cache {
	hashCode := this_.GetStringHashCode(key)
	return this_.GetCacheByHash(hashCode)
}
func (this_ *CacheGroup) GetCacheByHash(hashCode uint32) *Cache {
	index := hashCode % this_.cacheSize
	return this_.caches[index]
}

func (this_ *CacheGroup) Get(key string) (interface{}, bool) {
	return this_.GetCache(key).Get(key)
}

// Put 设置值
func (this_ *CacheGroup) Put(key string, v interface{}) {
	this_.GetCache(key).Put(key, v)
}

// PutIfAbsent 不存在则新增
func (this_ *CacheGroup) PutIfAbsent(key string, v interface{}) {
	this_.GetCache(key).PutIfAbsent(key, v)
}

// Delete 删除
func (this_ *CacheGroup) Delete(keys ...string) {
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
func (this_ *CacheGroup) Clear() {
	this_.init()
}
