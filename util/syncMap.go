package util

import "sync"

func NewSyncMap() *SyncMap {
	res := &SyncMap{
		cache:  make(map[interface{}]interface{}),
		locker: &sync.RWMutex{},
	}

	return res
}

type SyncMap struct {
	cache  map[interface{}]interface{}
	locker sync.Locker
}

func (this_ *SyncMap) Set(key interface{}, value interface{}) {
	this_.locker.Lock()
	defer this_.locker.Unlock()

	this_.cache[key] = value
}

func (this_ *SyncMap) Get(key interface{}) interface{} {
	this_.locker.Lock()
	defer this_.locker.Unlock()

	return this_.cache[key]
}

func (this_ *SyncMap) Clean() {
	this_.locker.Lock()
	defer this_.locker.Unlock()

	this_.cache = make(map[interface{}]interface{})
}

func (this_ *SyncMap) Len() int {
	this_.locker.Lock()
	defer this_.locker.Unlock()

	return len(this_.cache)
}
