package zookeeper

import "github.com/go-zookeeper/zk"

type IService interface {
	Stop()
	GetConn() *zk.Conn
	Create(path string, value string) (err error)
	CreateByMode(path string, value string, mode int32) (err error)
	Info() (info *Info, err error)
	Set(path string, value string) (err error)
	CreateIfNotExists(path string, value string) (err error)
	Exists(path string) (isExist bool, err error)
	Get(path string) (value string, err error)
	GetInfo(path string) (info *NodeInfo, err error)
	Stat(path string) (info *StatInfo, err error)
	GetChildren(path string) (children []string, err error)
	Delete(path string) (err error)
	WatchChildren(path string, listen func(data *WatchChildrenListenData) (finish bool)) (err error)
}
