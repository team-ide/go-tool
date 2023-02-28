package zookeeper

import (
	"errors"
	"fmt"
	"github.com/go-zookeeper/zk"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"sync"
)

type NodeEvent string

var NodeEventStopped = newNodeEvent("node event zk is stopped")
var NodeEventError = newNodeEvent("node event node listen error")
var NodeEventAdded = newNodeEvent("node event node added")
var NodeEventDeleted = newNodeEvent("node event node deleted")
var NodeEventNodeNotFound = newNodeEvent("node event node not found")

func newNodeEvent(str string) *NodeEvent {
	event := NodeEvent(str)
	return &event
}

// WatchChildren 监听子节点  子节点 新增 删除
func (this_ *ZKService) WatchChildren(path string, listen func(data *WatchChildrenListenData) (finish bool)) (err error) {
	if listen == nil {
		err = errors.New("listen is null")
		return
	}
	isExists, err := this_.Exists(path)
	if err != nil {
		return
	}
	if !isExists {
		err = errors.New("path [" + path + "] not exists")
		return
	}
	cache := &watchChildrenCache{
		ZKService:     this_,
		Path:          path,
		listen:        listen,
		childrenLock:  &sync.Mutex{},
		childrenCache: make(map[string]bool),
	}
	go cache.doChildrenW()
	return
}

type WatchChildrenListenData struct {
	Path  string
	Event *NodeEvent
	Child string
	Err   error
}

type watchChildrenCache struct {
	Path     string
	Children []string
	listen   func(data *WatchChildrenListenData) (finish bool)
	*ZKService
	childrenLock  sync.Locker
	childrenCache map[string]bool
}

func (this_ *watchChildrenCache) getListenData(event *NodeEvent, child string, err error) (data *WatchChildrenListenData) {
	data = &WatchChildrenListenData{
		Event: event,
		Child: child,
		Path:  this_.Path,
		Err:   err,
	}

	return
}

func (this_ *watchChildrenCache) change(newChildren []string) (finish bool) {
	this_.childrenLock.Lock()
	defer this_.childrenLock.Unlock()

	var newChildrenCache = make(map[string]bool)
	for _, child := range newChildren {
		newChildrenCache[child] = true
		// 查看在历史中是否存在
		_, find := this_.childrenCache[child]
		if find {
			continue
		} else {
			this_.childrenCache[child] = true
			this_.Children = append(this_.Children, child)
			// 节点新增通知
			finish = this_.listen(this_.getListenData(NodeEventAdded, child, nil))
			if finish {
				return
			}
		}
	}
	var newList []string
	for _, child := range this_.Children {
		// 查看在新节点中是否存在
		_, find := newChildrenCache[child]
		if find {
			newList = append(newList, child)
			continue
		} else {
			// 删除节点缓存
			delete(this_.childrenCache, child)
			// 节点新增通知
			finish = this_.listen(this_.getListenData(NodeEventDeleted, child, nil))
			if finish {
				return
			}
		}
	}
	this_.Children = newList

	return
}

func (this_ *watchChildrenCache) doChildrenW() {

	// 标记是否已完成 如果为 true 则不再监听
	var finish = false
	var err error
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
		}
		if err != nil {
			var event = NodeEventError
			if err == zk.ErrNoNode {
				event = NodeEventNodeNotFound
			}
			util.Logger.Error("zk doChildrenW error", zap.Any("path", this_.Path), zap.Any("error", err))
			finish = this_.listen(this_.getListenData(event, "", err))
		} else {
			if this_.isStop {
				err = errors.New("zk is stopped")
				util.Logger.Error("zk stopped", zap.Any("path", this_.Path), zap.Any("error", err))
				finish = this_.listen(this_.getListenData(NodeEventStopped, "", err))
			}
		}
		// 如果 未结束 则继续监听
		if !finish {
			this_.doChildrenW()
		}
	}()

	// 如果已停止 则跳出监听
	if this_.isStop {
		return
	}
	children, _, event, err := this_.GetConn().ChildrenW(this_.Path)
	if err != nil {
		return
	}
	finish = this_.change(children)
	<-event

}
