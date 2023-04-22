package thrift

import (
	"fmt"
	"github.com/team-ide/go-interpreter/node"
	"github.com/team-ide/go-interpreter/thrift"
	"github.com/team-ide/go-tool/util"
	"strings"
)

type ServiceInfo struct {
	Filename     string                   `json:"filename"`
	Name         string                   `json:"name"`
	RelativePath string                   `json:"relativePath"`
	ServiceNode  *thrift.ServiceStatement `json:"serviceNode"`
}

func NewWorkspace(dir string) *Workspace {
	formatDir := util.FormatPath(dir)
	res := &Workspace{
		dir:                dir,
		formatDir:          formatDir,
		treeCache:          util.NewSyncMap(),
		structCache:        util.NewSyncMap(),
		structCache_:       util.NewSyncMap(),
		serviceCache:       util.NewSyncMap(),
		serviceMethodCache: util.NewSyncMap(),
		enumCache:          util.NewSyncMap(),
		exceptionCache:     util.NewSyncMap(),
		includePathCache:   util.NewSyncMap(),
	}
	res.ignoreNames = []string{".git", ".idea", "node_modules"}
	return res
}

type Workspace struct {
	dir       string
	formatDir string

	ignoreNames   []string
	includeSubDir bool

	ServiceList        []*ServiceInfo
	treeCache          *util.SyncMap
	structCache        *util.SyncMap
	structCache_       *util.SyncMap
	serviceCache       *util.SyncMap
	serviceMethodCache *util.SyncMap
	enumCache          *util.SyncMap
	exceptionCache     *util.SyncMap
	includePathCache   *util.SyncMap
}

func (this_ *Workspace) IsIgnoreName(name string) bool {
	return util.StringIndexOf(this_.ignoreNames, name) >= 0
}

func (this_ *Workspace) AddIgnoreName(name string) {
	this_.ignoreNames = append(this_.ignoreNames, name)
}

func (this_ *Workspace) IncludeSubDir(includeSubDir bool) {
	this_.includeSubDir = includeSubDir
}

func (this_ *Workspace) IsIncludeSubDir() bool {
	return this_.includeSubDir
}

func (this_ *Workspace) GetDir() string {
	return this_.dir
}

func (this_ *Workspace) GetRelativePath(filename string) string {
	return strings.TrimLeft(filename, this_.formatDir)
}

func (this_ *Workspace) GetFormatDir() string {
	return this_.formatDir
}

func (this_ *Workspace) GetTree(filename string) *node.Tree {
	if res := this_.treeCache.Get(filename); res != nil {
		return res.(*node.Tree)
	}
	return nil
}

func (this_ *Workspace) SetTree(filename string, value *node.Tree) {
	fmt.Println("SetTree filename:", filename)
	this_.treeCache.Set(filename, value)
}

func (this_ *Workspace) GetStruct(filename string, name string) *thrift.StructStatement {
	if res := this_.structCache.Get(filename + "-" + name); res != nil {
		return res.(*thrift.StructStatement)
	}
	return nil
}

func (this_ *Workspace) SetStruct(filename string, value *thrift.StructStatement) {
	fmt.Println("SetStruct filename:", filename, ",name:", value.Name)
	this_.structCache.Set(filename+"-"+value.Name, value)
}

func (this_ *Workspace) GetService(filename string, name string) *thrift.ServiceStatement {
	if res := this_.serviceCache.Get(filename + "-" + name); res != nil {
		return res.(*thrift.ServiceStatement)
	}
	return nil
}

func (this_ *Workspace) SetService(filename string, value *thrift.ServiceStatement) {
	fmt.Println("SetService filename:", filename, ",name:", value.Name)
	this_.serviceCache.Set(filename+"-"+value.Name, value)
	for _, method := range value.Methods {
		this_.SetServiceMethod(filename, value.Name, method)
	}

	serviceInfo := &ServiceInfo{
		Filename:     filename,
		RelativePath: this_.GetRelativePath(filename),
		Name:         value.Name,
		ServiceNode:  value,
	}
	this_.ServiceList = append(this_.ServiceList, serviceInfo)
}

func (this_ *Workspace) GetServiceMethod(filename string, serviceName string, name string) *thrift.ServiceMethodNode {
	if res := this_.serviceMethodCache.Get(filename + "-" + serviceName + "-" + name); res != nil {
		return res.(*thrift.ServiceMethodNode)
	}
	return nil
}

func (this_ *Workspace) SetServiceMethod(filename string, serviceName string, value *thrift.ServiceMethodNode) {
	fmt.Println("SetServiceMethod filename:", filename, ",serviceName:", serviceName, ",name:", value.Name)
	this_.serviceMethodCache.Set(filename+"-"+serviceName+"-"+value.Name, value)
}

func (this_ *Workspace) GetEnum(filename string, name string) *thrift.EnumStatement {
	if res := this_.enumCache.Get(filename + "-" + name); res != nil {
		return res.(*thrift.EnumStatement)
	}
	return nil
}

func (this_ *Workspace) SetEnum(filename string, value *thrift.EnumStatement) {
	fmt.Println("SetEnum filename:", filename, ",name:", value.Name)
	this_.enumCache.Set(filename+"-"+value.Name, value)
}

func (this_ *Workspace) GetException(filename string, name string) *thrift.ExceptionStatement {
	if res := this_.exceptionCache.Get(filename + "-" + name); res != nil {
		return res.(*thrift.ExceptionStatement)
	}
	return nil
}

func (this_ *Workspace) SetException(filename string, value *thrift.ExceptionStatement) {
	fmt.Println("SetException filename:", filename, ",name:", value.Name)
	this_.exceptionCache.Set(filename+"-"+value.Name, value)
}

func (this_ *Workspace) GetIncludePath(filename string, name string) string {
	if res := this_.includePathCache.Get(filename); res != nil {
		data := res.(map[string]string)
		return data[name]
	}
	return ""
}

func (this_ *Workspace) SetIncludePath(dir string, filename string, value *thrift.IncludeStatement) {
	res := this_.includePathCache.Get(filename)
	if res == nil {
		res = make(map[string]string)
	}
	data := res.(map[string]string)
	path := util.FormatPath(dir + "/" + value.Include)
	name := strings.TrimRight(path, ".thrift")
	if strings.Index(name, "/") >= 0 {
		name = path[strings.LastIndex(path, "/")+1:]
	}

	data[name] = path
	fmt.Println("SetIncludePath filename:", filename, ",name:", name, ",path:", path)
	this_.includePathCache.Set(filename, data)
}
