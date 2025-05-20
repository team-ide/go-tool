package util

import (
	"hash/crc32"
	"sort"
	"sync"
	"time"
)

type ICache[K int | int64 | string, V any] interface {
	Get(key K) (res V, find bool)
	Gets(keys ...K) (res []V)
	PutIfAbsent(key string, v V)
	Put(key K, v V)
	Remove(key K)
	Removes(keys ...K)
	Size() int
	ItemsSize() int
}

// CacheItem 缓存项结构
type CacheItem[K int | int64 | string, V any] struct {
	Key         K
	Value       V
	WriteAt     int64  // 写入 时间 毫秒
	AccessAt    int64  // 访问 时间 毫秒
	AccessCount uint64 // 访问 次数
	Index       uint64 // 加入的顺序索引
}

type CacheOptions struct {
	// 缓存池大小，在缓存数量到达该大小时， 开始回收旧的数据
	MaximumSize int
	// 缓存 一段时间 未使用 从内存中删除  单位 毫秒
	ExpireAfterAccess int64
	// 缓存 在写入 一段时间后 不管 有没有使用 强制删除  单位 毫秒
	ExpireAfterWrite int64
	// 淘汰策略 定时器 设置 多长时间执行一次  默认 1 分钟
	EvictionPolicyTimer time.Duration
}

// Cache 临时缓存结构
type Cache[K int | int64 | string, V any] struct {
	items     map[K]*CacheItem[K, V]
	itemsLock sync.Locker
	size      int // 当前缓存大小

	ticker *time.Ticker // 用于定期清理的ticker

	// 缓存池大小，在缓存数量到达该大小时， 开始回收旧的数据
	maximumSize int
	// 缓存 一段时间 未使用 从内存中删除  单位 毫秒
	expireAfterAccess int64
	// 缓存 在写入 一段时间后 不管 有没有使用 强制删除  单位 毫秒
	expireAfterWrite int64
	// 淘汰策略 定时器 设置 多长时间执行一次  默认 1 分钟
	evictionPolicyTimer time.Duration

	itemIndex uint64

	onRemoves func(isEvictionPolicy bool, items ...*CacheItem[K, V])
}

func (this_ *Cache[K, V]) SetOnRemoves(onRemoves func(isEvictionPolicy bool, items ...*CacheItem[K, V])) {
	this_.onRemoves = onRemoves
}

// NewCache 创建 默认 缓存
func NewCache[K int | int64 | string, V any]() *Cache[K, V] {
	return NewCacheByOptions[K, V](nil)
}

// NewCacheByOptions 创建一个 带策略的 缓存
func NewCacheByOptions[K int | int64 | string, V any](opts *CacheOptions) *Cache[K, V] {
	if opts == nil {
		opts = &CacheOptions{}
	}
	tc := &Cache[K, V]{
		items:     make(map[K]*CacheItem[K, V]),
		itemsLock: &sync.Mutex{},
		size:      0,

		maximumSize:         opts.MaximumSize,
		expireAfterAccess:   opts.ExpireAfterAccess,
		expireAfterWrite:    opts.ExpireAfterWrite,
		evictionPolicyTimer: opts.EvictionPolicyTimer,
	}
	if tc.maximumSize > 0 || tc.expireAfterAccess > 0 || tc.expireAfterWrite > 0 {
		if tc.evictionPolicyTimer > 0 {
			tc.ticker = time.NewTicker(tc.evictionPolicyTimer)
		} else {
			// 每分钟检测一次
			tc.ticker = time.NewTicker(time.Minute)
		}
		go tc.startEvictionPolicy() // 启动清理协程
	}
	return tc
}

// Get 从缓存中获取一个项
func (this_ *Cache[K, V]) Get(key K) (res V, find bool) {
	this_.itemsLock.Lock()

	item, find := this_.items[key]

	this_.itemsLock.Unlock()

	if !find {
		return
	}
	res = item.Value
	item.AccessAt = time.Now().UnixMilli()
	return
}

// Gets 获取缓存多个项
func (this_ *Cache[K, V]) Gets(keys ...K) (res []V) {

	this_.itemsLock.Lock()

	accessAt := time.Now().UnixMilli()
	for _, key := range keys {
		item, find := this_.items[key]
		if find {
			res = append(res, item.Value)
			item.AccessAt = accessAt
			item.AccessCount++
		}
	}

	this_.itemsLock.Unlock()

	return
}

// GetAnRemove 从缓存中获取一个项
func (this_ *Cache[K, V]) GetAnRemove(key K) (res V, find bool) {
	this_.itemsLock.Lock()

	item, find := this_.items[key]
	if find {
		delete(this_.items, key)
		this_.size--
	}

	this_.itemsLock.Unlock()

	if !find {
		return
	}
	res = item.Value
	return
}

// GetOrLoad 从缓存中获取一个项 如果没有 则调用 load 加载
func (this_ *Cache[K, V]) GetOrLoad(key K, load func() (V, error)) (res V, err error) {
	this_.itemsLock.Lock()

	item, find := this_.items[key]

	this_.itemsLock.Unlock()

	if find {
		res = item.Value
		item.AccessAt = time.Now().UnixMilli()
		item.AccessCount++
		return
	}
	//if this_.outLog {
	//	Logger.Debug("缓存不存在 key [" + key + "] ，执行加载")
	//}
	value, err := load()
	if err != nil {
		return value, err
	}
	this_.Set(key, value)
	return value, err
}

// Set 从缓存中获取一个项 如果没有 则调用 load 加载
func (this_ *Cache[K, V]) Set(key K, value V) {
	this_.itemsLock.Lock()

	_, find := this_.items[key]
	this_.setNoLock(key, value)
	if !find {
		this_.size++
	}

	this_.itemsLock.Unlock()
}

func (this_ *Cache[K, V]) setNoLock(key K, value V) {
	item := &CacheItem[K, V]{
		Value:   value,
		WriteAt: time.Now().UnixMilli(),
	}
	item.Index = this_.itemIndex
	item.Key = key

	this_.itemIndex++
	this_.items[key] = item
}

func (this_ *Cache[K, V]) Put(key K, v V) {
	this_.Set(key, v)
}

// PutIfAbsent 不存在则新增
func (this_ *Cache[K, V]) PutIfAbsent(key K, v V) {
	this_.itemsLock.Lock()

	item, find := this_.items[key]
	if !find {
		this_.setNoLock(key, v)
		this_.size++
	} else {
		item.WriteAt = time.Now().UnixMilli()
	}

	this_.itemsLock.Unlock()
}

// Clear 清理
func (this_ *Cache[K, V]) Clear() {
	this_.itemsLock.Lock()

	this_.size = 0
	this_.itemIndex = 0
	this_.items = make(map[K]*CacheItem[K, V])

	this_.itemsLock.Unlock()

}

// Size 缓存 数量
func (this_ *Cache[K, V]) Size() int {
	return this_.size
}

// ItemsSize 缓存 数量
func (this_ *Cache[K, V]) ItemsSize() int {
	this_.itemsLock.Lock()

	s := len(this_.items)

	this_.itemsLock.Unlock()
	return s
}

// Remove 清理 缓存
func (this_ *Cache[K, V]) Remove(key K) {
	this_.itemsLock.Lock()

	this_.doRemoveNoLock(key)

	this_.itemsLock.Unlock()
}

// Removes 清理 缓存
func (this_ *Cache[K, V]) Removes(keys ...K) {
	if len(keys) == 0 {
		return
	}

	this_.itemsLock.Lock()

	this_.doRemoveNoLock(keys...)

	this_.itemsLock.Unlock()
}

func (this_ *Cache[K, V]) doRemoveNoLock(keys ...K) {
	if len(keys) == 0 {
		return
	}
	var ok bool
	for _, key := range keys {
		//if this_.outLog {
		//	Logger.Debug("删除 缓存 中的 key [" + removeKey + "]")
		//}
		_, ok = this_.items[key]
		if ok {
			delete(this_.items, key)
			this_.size--
		}
	}
	//if this_.outLog {
	//	Logger.Debug(fmt.Sprintf("删除 缓存 后 缓存大小 [%d]", this_.size))
	//}
}

// startEvictionPolicy 定期清理过期的缓存项
func (this_ *Cache[K, V]) startEvictionPolicy() {
	for {
		select {
		case <-this_.ticker.C:
			this_.cleanExpireAfter()
			this_.cleanMaximumSize()
		}
	}
}

type itemList[K int | int64 | string, V any] []*CacheItem[K, V]

func (this_ itemList[K, V]) Len() int { return len(this_) }
func (this_ itemList[K, V]) Swap(i, j int) {
	this_[i], this_[j] = this_[j], this_[i]
}
func (this_ itemList[K, V]) Less(i, j int) bool { return this_[i].AccessAt < this_[j].AccessAt }

// cleanLimit 清理超出 limit 的缓存
func (this_ *Cache[K, V]) cleanMaximumSize() {
	if this_.maximumSize <= 0 {
		return
	}

	this_.itemsLock.Lock()

	this_.doCleanMaximumSizeNoLock()

	this_.itemsLock.Unlock()
}

func (this_ *Cache[K, V]) doCleanMaximumSizeNoLock() {
	if this_.maximumSize <= 0 {
		return
	}
	removeKeySize := this_.size - this_.maximumSize
	if removeKeySize <= 0 {
		return
	}

	var items itemList[K, V]
	for _, item := range this_.items {
		items = append(items, item)
	}
	sort.Sort(items)
	removeItems := items[:removeKeySize]
	var removeKeys []K
	for _, item := range removeItems {
		removeKeys = append(removeKeys, item.Key)
	}
	if len(removeKeys) == 0 {
		return
	}

	this_.doRemoveNoLock(removeKeys...)

}

func (this_ *Cache[K, V]) cleanExpireAfter() {

	this_.itemsLock.Lock()

	this_.doCleanExpireAfterWriteNoLock()

	this_.doCleanExpireAfterAccessNoLock()

	this_.itemsLock.Unlock()

}

func (this_ *Cache[K, V]) doCleanExpireAfterWriteNoLock() {
	if this_.expireAfterWrite <= 0 {
		return
	}
	var now = time.Now().UnixMilli()
	// 过期的 写入 时间  清理 小于 这个时间 的 数据
	var cleanAfterWrite = now - this_.expireAfterWrite
	var removeKeys []K

	for key, item := range this_.items {
		// 清理 小于 过期的 写入 时间 的 数据
		if item.WriteAt <= cleanAfterWrite {
			removeKeys = append(removeKeys, key)
		}
	}

	if len(removeKeys) == 0 {
		return
	}
	this_.doRemoveNoLock(removeKeys...)
}

func (this_ *Cache[K, V]) doCleanExpireAfterAccessNoLock() {
	if this_.expireAfterAccess <= 0 {
		return
	}
	var now = time.Now().UnixMilli()
	// 过期的 访问 时间  清理 小于 这个时间 的 数据
	var cleanAfterAccess = now - this_.expireAfterAccess
	var removeKeys []K

	for key, item := range this_.items {
		// 清理 小于 过期的 访问 时间 的 数据
		if item.AccessAt <= cleanAfterAccess {
			removeKeys = append(removeKeys, key)
		}
	}
	if len(removeKeys) == 0 {
		return
	}
	this_.doRemoveNoLock(removeKeys...)
}

type CacheGroup[K int | int64 | string, V any] struct {
	firstCache *Cache[K, V]
	caches     []*Cache[K, V]
	cacheSize  uint32
	opts       *CacheOptions

	getCacheIndex func(key K) uint32
}

func NewCacheGroup[K int | int64 | string, V any](cacheSize uint) *CacheGroup[K, V] {
	g := &CacheGroup[K, V]{
		cacheSize: uint32(cacheSize),
	}
	g.init()
	return g
}

func NewCacheGroupByOption[K int | int64 | string, V any](cacheSize uint, opts *CacheOptions) *CacheGroup[K, V] {
	g := &CacheGroup[K, V]{
		cacheSize: uint32(cacheSize),
		opts:      opts,
	}
	g.init()
	return g
}

func (this_ *CacheGroup[K, V]) SetGetCacheIndex(getCacheIndex func(key K) uint32) {
	this_.getCacheIndex = getCacheIndex
}

func (this_ *CacheGroup[K, V]) init() {
	var caches []*Cache[K, V]
	var i uint32
	for i = 0; i < this_.cacheSize; i++ {
		var cache = NewCacheByOptions[K, V](this_.opts)
		caches = append(caches, cache)
		if i == 0 {
			this_.firstCache = cache
		}
	}
	this_.caches = caches
}

func (this_ *CacheGroup[K, V]) GetKeyHashCode(key K) uint32 {
	return this_.GetStringHashCode(GetStringValue(key))
}
func (this_ *CacheGroup[K, V]) GetStringHashCode(str string) uint32 {
	//h := fnv.New32a()
	//_, _ = h.Write([]byte(str))
	//return h.Sum32()
	return crc32.ChecksumIEEE([]byte(str))
}

func (this_ *CacheGroup[K, V]) GetCache(key K) *Cache[K, V] {
	if this_.cacheSize == 1 {
		return this_.firstCache
	}
	var index = this_.GetCacheIndex(key)
	return this_.caches[index]
}

func (this_ *CacheGroup[K, V]) GetCacheIndex(key K) uint32 {
	if this_.cacheSize == 1 {
		return 0
	}
	var index uint32
	if this_.getCacheIndex != nil {
		index = this_.getCacheIndex(key)
	} else {
		hashCode := this_.GetKeyHashCode(key)
		index = hashCode % this_.cacheSize
	}
	return index
}

// Gets 获取缓存多个项
func (this_ *CacheGroup[K, V]) Gets(keys ...K) (res []V) {
	if this_.cacheSize == 1 {
		return this_.firstCache.Gets(keys...)
	}
	indexKs := map[uint32][]K{}
	for _, key := range keys {
		index := this_.GetCacheIndex(key)
		indexKs[index] = append(indexKs[index], key)
	}
	for index, ks := range indexKs {
		list := this_.caches[index].Gets(ks...)
		res = append(res, list...)
	}
	return
}

func (this_ *CacheGroup[K, V]) Get(key K) (V, bool) {
	return this_.GetCache(key).Get(key)
}

// Put 设置值
func (this_ *CacheGroup[K, V]) Put(key K, v V) {
	this_.GetCache(key).Put(key, v)
}

// PutIfAbsent 不存在则新增
func (this_ *CacheGroup[K, V]) PutIfAbsent(key K, v V) {
	this_.GetCache(key).PutIfAbsent(key, v)
}

// Remove 删除
func (this_ *CacheGroup[K, V]) Remove(key K) {
	this_.GetCache(key).Remove(key)
}

// Removes 删除
func (this_ *CacheGroup[K, V]) Removes(keys ...K) {
	if this_.cacheSize == 1 {
		this_.firstCache.Removes(keys...)
		return
	}
	indexKs := map[uint32][]K{}
	for _, key := range keys {
		index := this_.GetCacheIndex(key)
		indexKs[index] = append(indexKs[index], key)
	}
	for index, ks := range indexKs {
		this_.caches[index].Removes(ks...)
	}
}

// Size 数量
func (this_ *CacheGroup[K, V]) Size() (res int) {
	for _, c := range this_.caches {
		res += c.Size()
	}
	return
}

// ItemsSize Items 数量
func (this_ *CacheGroup[K, V]) ItemsSize() (res int) {
	for _, c := range this_.caches {
		res += c.ItemsSize()
	}
	return
}

func (this_ *CacheGroup[K, V]) SetOnRemoves(onRemoves func(isEvictionPolicy bool, items ...*CacheItem[K, V])) {
	for _, c := range this_.caches {
		c.SetOnRemoves(onRemoves)
	}
}

// Clear 清理
func (this_ *CacheGroup[K, V]) Clear() {
	this_.init()
}
