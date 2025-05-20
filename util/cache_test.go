package util

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	opts := &CacheOptions{
		MaximumSize:         990000,    // 缓存大小限制
		ExpireAfterWrite:    1000 * 10, // 缓存项存活时间 毫秒
		EvictionPolicyTimer: time.Second,
	}
	//cache := NewCacheByOptions[string, string](opts)
	cache := NewCacheGroupByOption[string, string](10, opts)
	//cache := NewCacheGroup(100)
	var threadSize = 100
	g := sync.WaitGroup{}
	g.Add(threadSize)
	var callStop bool
	for i := 0; i < threadSize; i++ {
		go func() {
			for !callStop {
				key := fmt.Sprintf("%d", RandomInt(100, 1000000))

				//key := fmt.Sprintf("%d", num)
				cache.Get(key)
				cache.PutIfAbsent(key, key)
				cache.Put(key, key)
			}
			g.Done()
		}()
	}
	go func() {
		time.Sleep(60 * time.Second)
		callStop = true
	}()
	g.Wait()
	fmt.Println("cache size:", cache.Size(), ",items len:", cache.ItemsSize())
}

func TestCacheUse(t *testing.T) {
	testCache("不分组", NewCache[string, string]())
	testCache(fmt.Sprintf("分组数量：%d", 1), NewCacheGroup[string, string](1))
	testCache(fmt.Sprintf("分组数量：%d", 2), NewCacheGroup[string, string](2))
	testCache(fmt.Sprintf("分组数量：%d", 10), NewCacheGroup[string, string](10))
	testCache(fmt.Sprintf("分组数量：%d", 20), NewCacheGroup[string, string](20))
	testCache(fmt.Sprintf("分组数量：%d", 30), NewCacheGroup[string, string](30))
	testCache(fmt.Sprintf("分组数量：%d", 40), NewCacheGroup[string, string](40))
	testCache(fmt.Sprintf("分组数量：%d", 50), NewCacheGroup[string, string](50))
	testCache(fmt.Sprintf("分组数量：%d", 60), NewCacheGroup[string, string](60))
	testCache(fmt.Sprintf("分组数量：%d", 70), NewCacheGroup[string, string](70))
	testCache(fmt.Sprintf("分组数量：%d", 80), NewCacheGroup[string, string](80))
	testCache(fmt.Sprintf("分组数量：%d", 90), NewCacheGroup[string, string](90))
	testCache(fmt.Sprintf("分组数量：%d", 100), NewCacheGroup[string, string](100))
}

func testCache(name string, cache ICache[string, string]) {
	//cache := NewCacheGroup(100)
	var size = 100000
	var threadSize = 100
	g := sync.WaitGroup{}
	g.Add(threadSize)
	startTime := time.Now()
	for i := 0; i < threadSize; i++ {
		go func() {
			for n := 0; n < size; n++ {
				key := fmt.Sprintf("%d", RandomInt(100, 1000000))
				//key := fmt.Sprintf("%d", num)
				cache.Get(key)
				cache.PutIfAbsent(key, key)
				cache.Put(key, key)
			}
			g.Done()
		}()
	}
	g.Wait()
	endTime := time.Now()
	useNano := endTime.UnixNano() - startTime.UnixNano()
	count := size * threadSize
	fmt.Println(name, "，耗时：", useNano/1000000, "毫秒", "，执行：", count, "次，平均：", int64(useNano)/int64(count), "纳秒/次", ",cache size:", cache.Size(), ",items len:", cache.ItemsSize())
}
