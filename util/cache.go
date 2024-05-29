package util

import (
	"hash/fnv"
	"sync"
)

type Cache[T any] struct {
	obj  T
	data map[string]T
	lock sync.Mutex
}

func NewCache[T any](obj T) *Cache[T] {
	return &Cache[T]{
		obj:  obj,
		data: make(map[string]T),
	}
}

func (this_ *Cache[T]) Get(key string) (T, bool) {
	this_.lock.Lock()
	v, find := this_.data[key]
	this_.lock.Unlock()
	return v, find
}

// Put 设置值
func (this_ *Cache[T]) Put(key string, v T) {
	this_.lock.Lock()
	this_.data[key] = v
	this_.lock.Unlock()
}

// PutIfAbsent 不存在则新增
func (this_ *Cache[T]) PutIfAbsent(key string, v T) {
	this_.lock.Lock()
	_, find := this_.data[key]
	if !find {
		this_.data[key] = v
	}
	this_.lock.Unlock()
}

// Delete 删除
func (this_ *Cache[T]) Delete(keys ...string) {
	this_.lock.Lock()
	for _, key := range keys {
		delete(this_.data, key)
	}
	this_.lock.Unlock()
}

// Clear 清理
func (this_ *Cache[T]) Clear() {
	this_.lock.Lock()
	this_.data = make(map[string]T)
	this_.lock.Unlock()
}

type CacheGroup[T any] struct {
	obj       T
	caches    []*Cache[T]
	cacheSize uint32
}

func NewCacheGroup[T any](cacheSize int, obj T) *CacheGroup[T] {
	g := &CacheGroup[T]{
		obj:       obj,
		cacheSize: uint32(cacheSize),
	}
	g.init()
	return g
}

func (this_ *CacheGroup[T]) init() {
	var caches []*Cache[T]
	var i uint32
	for i = 0; i < this_.cacheSize; i++ {
		var cache = NewCache(this_.obj)
		caches = append(caches, cache)
	}
	this_.caches = caches
}
func (this_ *CacheGroup[T]) GetStringHashCode(str string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(str))
	return h.Sum32()
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
