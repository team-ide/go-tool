package zookeeper

type IService interface {
	Stop()
	Create(path string, data []byte, mode int32) (err error)
	Info() (info *Info, err error)
	Set(path string, data []byte) (err error)
	CreateIfNotExists(path string, data []byte) (err error)
	Exists(path string) (isExist bool, err error)
	Get(path string) (info *NodeInfo, err error)
	Stat(path string) (info *StatInfo, err error)
	GetChildren(path string) (children []string, err error)
	Delete(path string) (err error)
	WatchChildren(path string, listen func(data *WatchChildrenListenData) (finish bool)) (err error)
}
