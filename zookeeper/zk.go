package zookeeper

import (
	"errors"
	"fmt"
	"github.com/go-zookeeper/zk"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
	"net"
	"sort"
	"strings"
	"time"
)

var ZKLogger zk.Logger = &defaultLogger{}

type defaultLogger struct{}

var (
	logger *zap.Logger
)

func getLogger() *zap.Logger {
	if logger == nil {
		// 用于 输出 上层方法
		logger = util.NewLoggerByCallerSkip(1)
	}
	return logger
}

func (*defaultLogger) Printf(format string, args ...interface{}) {
	getLogger().Info(fmt.Sprintf("zookeeper log:"+format, args...))
}

// Service 注册处理器在线信息等
type Service struct {
	*Config
	zkConn   *zk.Conn        //zk连接
	zkEvent  <-chan zk.Event // zk事件通知管道
	isClosed bool
}

func (this_ *Service) init(sshClient *ssh.Client) (err error) {
	if this_.ConnectionTimeout == 0 {
		this_.ConnectionTimeout = 10000
	}
	if this_.SessionTimeout == 0 {
		this_.SessionTimeout = 60000
	}
	connectionTimeout := time.Millisecond * time.Duration(this_.ConnectionTimeout)
	sessionTimeout := time.Millisecond * time.Duration(this_.SessionTimeout)
	if sshClient != nil {
		this_.zkConn, this_.zkEvent, err = zk.Connect(this_.GetServers(), sessionTimeout, func(c *zk.Conn) {
			c.SetLogger(ZKLogger)
		}, zk.WithDialer(func(network, address string, timeout time.Duration) (net.Conn, error) {
			conn, e := sshClient.Dial(network, address)
			return &util.SSHChanConn{Conn: conn}, e
		}))
	} else {
		this_.zkConn, this_.zkEvent, err = zk.Connect(this_.GetServers(), sessionTimeout, func(c *zk.Conn) {
			c.SetLogger(ZKLogger)
		}, zk.WithDialer(func(network, address string, timeout time.Duration) (net.Conn, error) {
			conn, e := net.DialTimeout(network, address, connectionTimeout)
			return conn, e
		}))
	}
	if err != nil {
		util.Logger.Error("zk.Connect error", zap.Any("servers", this_.GetServers()), zap.Error(err))
		if this_.zkConn != nil {
			this_.zkConn.Close()
		}
		return
	}
	if this_.Username != "" || this_.Password != "" {
		err = this_.zkConn.AddAuth(this_.Username, []byte(this_.Password))
		if err != nil {
			util.Logger.Error("zk.Connect AddAuth error", zap.Any("servers", this_.GetServers()), zap.Error(err))
			if this_.zkConn != nil {
				this_.zkConn.Close()
			}
			return
		}
	}
	this_.isClosed = false
	return
}

func (this_ *Service) GetServers() []string {
	var servers []string
	if this_.Address == "" {
		return servers
	}
	if strings.Contains(this_.Address, ",") {
		servers = strings.Split(this_.Address, ",")
	} else if strings.Contains(this_.Address, ";") {
		servers = strings.Split(this_.Address, ";")
	} else {
		servers = []string{this_.Address}
	}
	return servers
}

func (this_ *Service) Close() {

	if this_ == nil {
		return
	}
	if this_.isClosed {
		return
	}
	this_.isClosed = true
	conn := this_.GetConn()
	if conn != nil {
		conn.Close()
	}
}

func (this_ *Service) GetConn() *zk.Conn {
	if this_ == nil {
		return nil
	}
	return this_.zkConn
}

type Info struct {
	Server    string   `json:"server"`
	SessionID int64    `json:"sessionID"`
	State     zk.State `json:"state"`
}

// Info ZK信息
func (this_ *Service) Info() (info *Info, err error) {
	info = &Info{}
	info.SessionID = this_.GetConn().SessionID()
	info.Server = this_.GetConn().Server()
	info.State = this_.GetConn().State()
	return
}

func (this_ *Service) Create(path string, data string) (err error) {
	if _, err = this_.GetConn().Create(path, []byte(data), 0, zk.WorldACL(zk.PermAll)); err != nil {
		err = errors.New("path [" + path + "] create error:" + err.Error())
		return
	}
	return
}

// CreateIfNotExists 一层层检查，如果不存在则创建
func (this_ *Service) CreateIfNotExists(path string, data string) (err error) {
	isExist, err := this_.Exists(path)
	if err != nil {
		return
	}
	if isExist {
		return
	}
	if strings.LastIndex(path, "/") > 0 {
		parentPath := path[0:strings.LastIndex(path, "/")]
		err = this_.CreateIfNotExists(parentPath, "")
		if err != nil {
			if err == zk.ErrNodeExists {
				err = nil
			} else {
				return
			}
		}
	}
	if _, err = this_.GetConn().Create(path, []byte(data), 0, zk.WorldACL(zk.PermAll)); err != nil {
		if err == zk.ErrNodeExists {
			err = nil
		} else {
			err = errors.New("path [" + path + "] create error:" + err.Error())
		}
	}
	return
}

func (this_ *Service) CreateEphemeral(path string, data string) (err error) {
	if _, err = this_.GetConn().Create(path, []byte(data), zk.FlagEphemeral, zk.WorldACL(zk.PermAll)); err != nil {
		err = errors.New("path [" + path + "] create error:" + err.Error())
		return
	}
	return
}

func (this_ *Service) CreateEphemeralIfNotExists(path string, data string) (err error) {
	isExist, err := this_.Exists(path)
	if err != nil {
		return
	}
	if isExist {
		return
	}
	if strings.LastIndex(path, "/") > 0 {
		parentPath := path[0:strings.LastIndex(path, "/")]
		err = this_.CreateIfNotExists(parentPath, "")
		if err != nil {
			if err == zk.ErrNodeExists {
				err = nil
			} else {
				return
			}
		}
	}
	if _, err = this_.GetConn().Create(path, []byte(data), zk.FlagEphemeral, zk.WorldACL(zk.PermAll)); err != nil {
		if err == zk.ErrNodeExists {
			err = nil
		} else {
			err = errors.New("path [" + path + "] create error:" + err.Error())
		}
	}
	return
}

// Set 修改节点数据
func (this_ *Service) Set(path string, data string) (err error) {
	isExist, state, err := this_.GetConn().Exists(path)
	if err != nil {
		return
	}
	if !isExist {
		err = errors.New("path:" + path + " not exists")
		return
	}
	if _, err = this_.GetConn().Set(path, []byte(data), state.Version); err != nil {
		err = errors.New("path [" + path + "] set error:" + err.Error())
		return
	}
	return
}

// Exists 判断节点是否存在
func (this_ *Service) Exists(path string) (isExist bool, err error) {
	isExist, _, err = this_.GetConn().Exists(path)
	if err != nil {
		err = errors.New("path [" + path + "] exists error:" + err.Error())
		return
	}
	return
}

type StatInfo struct {
	Czxid          int64 `json:"czxid,omitempty"`
	Mzxid          int64 `json:"mzxid,omitempty"`
	Ctime          int64 `json:"ctime,omitempty"`
	Mtime          int64 `json:"mtime,omitempty"`
	Version        int32 `json:"version,omitempty"`
	Cversion       int32 `json:"cversion,omitempty"`
	Aversion       int32 `json:"aversion,omitempty"`
	EphemeralOwner int64 `json:"ephemeralOwner,omitempty"`
	DataLength     int32 `json:"dataLength,omitempty"`
	NumChildren    int32 `json:"numChildren,omitempty"`
	Pzxid          int64 `json:"pzxid,omitempty"`
}

type NodeInfo struct {
	Path string    `json:"path"`
	Data string    `json:"data"`
	Stat *StatInfo `json:"stat"`
}

// Get 获取节点信息
func (this_ *Service) Get(path string) (data string, err error) {

	bs, _, err := this_.GetConn().Get(path)
	if err != nil {
		err = errors.New("path [" + path + "] get error:" + err.Error())
		return
	}
	data = string(bs)
	return
}

// GetInfo 获取节点信息
func (this_ *Service) GetInfo(path string) (info *NodeInfo, err error) {

	data, stat, err := this_.GetConn().Get(path)
	if err != nil {
		err = errors.New("path [" + path + "] get error:" + err.Error())
		return
	}
	info = &NodeInfo{}
	info.Data = string(data)

	if stat != nil {
		info.Stat = StatToInfo(stat)
	}
	return
}

func StatToInfo(stat *zk.Stat) (info *StatInfo) {
	info = &StatInfo{
		Czxid:          stat.Czxid,
		Mzxid:          stat.Mzxid,
		Ctime:          stat.Ctime,
		Mtime:          stat.Mtime,
		Version:        stat.Version,
		Cversion:       stat.Cversion,
		Aversion:       stat.Aversion,
		EphemeralOwner: stat.EphemeralOwner,
		DataLength:     stat.DataLength,
		NumChildren:    stat.NumChildren,
		Pzxid:          stat.Pzxid,
	}
	return
}

// Stat 获取节点状态
func (this_ *Service) Stat(path string) (info *StatInfo, err error) {
	_, stat, err := this_.GetConn().Exists(path)
	if err != nil {
		err = errors.New("path [" + path + "] exists error:" + err.Error())
		return
	}
	if stat != nil {
		info = StatToInfo(stat)
	}
	return
}

// GetChildren 查询子节点
func (this_ *Service) GetChildren(path string) (children []string, err error) {
	children, _, err = this_.GetConn().Children(path)
	if err != nil {
		err = errors.New("path [" + path + "] children error:" + err.Error())
		return
	}
	sort.Slice(children, func(i, j int) bool {
		return strings.ToLower(children[i]) < strings.ToLower(children[j]) //升序  即前面的值比后面的小 忽略大小写排序
	})
	return
}

// Delete 删除节点 如果有子节点，则子节点一同删除
func (this_ *Service) Delete(path string) (err error) {
	var isExist bool
	var stat *zk.Stat
	isExist, stat, err = this_.GetConn().Exists(path)
	if err != nil {
		err = errors.New("path [" + path + "] exists error:" + err.Error())
		return
	}
	if !isExist {
		return
	}
	var children []string
	children, _, err = this_.GetConn().Children(path)
	if err != nil {
		err = errors.New("path [" + path + "] children error:" + err.Error())
		return
	}
	if len(children) > 0 {
		for _, one := range children {
			err = this_.Delete(path + "/" + one)
			if err != nil {
				return
			}
		}
	}
	err = this_.GetConn().Delete(path, stat.Version)
	if err != nil {
		err = errors.New("path [" + path + "] delete error:" + err.Error())
		return
	}
	return
}
