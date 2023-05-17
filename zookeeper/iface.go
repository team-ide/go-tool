package zookeeper

import "github.com/go-zookeeper/zk"

type IService interface {
	// Close 关闭 客户端
	Close()
	// GetConn 获取 zk Conn
	GetConn() *zk.Conn
	// Info 查看 zk 相关信息
	Info() (info *Info, err error)
	// Create 创建 永久 节点
	Create(path string, value string) (err error)
	// CreateIfNotExists 如果不存在 则创建 永久 节点 创建时候如果已存在不报错  如果 父节点不存在 则先创建父节点
	CreateIfNotExists(path string, value string) (err error)
	// CreateEphemeral 创建 临时 节点
	CreateEphemeral(path string, value string) (err error)
	// CreateEphemeralIfNotExists 如果不存在 则创建 临时 节点 创建时候如果已存在不报错 如果 父节点不存在 则先创建父节点
	CreateEphemeralIfNotExists(path string, value string) (err error)
	// Exists 查看节点是否存在
	Exists(path string) (isExist bool, err error)
	// Set 设置 节点 值
	Set(path string, value string) (err error)
	// Get 查看 节点 数据
	Get(path string) (value string, err error)
	// GetInfo 查看 节点 信息
	GetInfo(path string) (info *NodeInfo, err error)
	// Stat 节点 状态
	Stat(path string) (info *StatInfo, err error)
	// GetChildren 查询 子节点
	GetChildren(path string) (children []string, err error)
	// Delete 删除节点 如果 包含子节点 则先删除所有子节点
	Delete(path string) (err error)
	// WatchChildren 监听 子节点 只监听当前节点下的子节点 NodeEventError 监听异常 NodeEventStopped zk客户端关闭 NodeEventAdded 节点新增 NodeEventDeleted 节点删除 NodeEventNodeNotFound 节点不存在
	WatchChildren(path string, listen func(data *WatchChildrenListenData) (finish bool)) (err error)
}
