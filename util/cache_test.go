package util

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	var vType string
	cache := NewCache(vType)
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
	fmt.Println("耗时：", endTime.UnixMilli()-startTime.UnixMilli(), "毫秒")
}

func TestCacheGroup(t *testing.T) {
	var vType string
	cache := NewCacheGroup(100, vType)
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
	fmt.Println("耗时：", endTime.UnixMilli()-startTime.UnixMilli(), "毫秒")
}
