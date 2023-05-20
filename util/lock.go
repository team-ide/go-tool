package util

import "sync"

var (
	lockMapLock = &sync.Mutex{}
	lockMap     = make(map[string]sync.Locker)
)

// GetLock 获取一个Locker，如果不存在，则新建
// obj = GetLock("user:1")
// obj.Lock()
// obj.Unlock()
func GetLock(key string) (lock sync.Locker) {

	lockMapLock.Lock()

	defer lockMapLock.Unlock()

	var ok bool
	lock, ok = lockMap[key]
	if ok {
		return
	}
	lock = &sync.Mutex{}
	lockMap[key] = lock
	return lock
}

// LockByKey 根据Key进行同步锁
// LockByKey("user:1")
func LockByKey(key string) {
	locker := GetLock(key)
	locker.Lock()
}

// UnlockByKey 根据Key进行解锁同步锁
// UnlockByKey("user:1")
func UnlockByKey(key string) {
	locker := GetLock(key)
	locker.Unlock()
}
