package thrift

import (
	"errors"
	"github.com/team-ide/go-interpreter/thrift"
	"os"
	"strings"
)

func (this_ *Workspace) Load() (errs map[string]error) {
	errs = map[string]error{}

	this_.loadByDir(this_.formatDir, errs)

	return
}
func (this_ *Workspace) loadByDir(dir string, errs map[string]error) {
	fs, err := os.ReadDir(dir)
	if err != nil {
		err = errors.New("Workspace loadByDir ReadDir [" + dir + "] error:" + err.Error())
		errs[dir] = err
		return
	}
	for _, f := range fs {
		name := f.Name()
		if this_.IsIgnoreName(name) {
			continue
		}
		path := dir + "/" + f.Name()
		if f.IsDir() {
			if !this_.IsIncludeSubDir() {
				continue
			}
			this_.loadByDir(path, errs)
		} else {
			if !strings.HasSuffix(f.Name(), ".thrift") {
				continue
			}
			err = this_.loadByFilename(dir, path)
			if err != nil {
				err = errors.New("Workspace loadByFilename [" + path + "] error:" + err.Error())
				errs[path] = err
			}
		}
	}
	return
}

func (this_ *Workspace) loadByFilename(dir string, filename string) (err error) {
	//fmt.Println("loadByFilename dir:", dir, ",filename:", filename)
	bs, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	code := string(bs)
	tree, err := thrift.Parse(filename, code)
	if tree == nil {
		return
	}
	this_.SetTree(filename, tree)

	for _, node := range tree.Children {

		if stem, ok := node.(*thrift.StructStatement); ok {
			this_.SetStruct(filename, stem)
		} else if stem, ok := node.(*thrift.ExceptionStatement); ok {
			this_.SetException(filename, stem)
		} else if stem, ok := node.(*thrift.EnumStatement); ok {
			this_.SetEnum(filename, stem)
		} else if stem, ok := node.(*thrift.ServiceStatement); ok {
			this_.SetService(filename, stem)
		} else if stem, ok := node.(*thrift.IncludeStatement); ok {
			this_.SetIncludePath(dir, filename, stem)
		}
	}
	return
}
func (this_ *Workspace) Clean() {
	this_.treeCache.Clean()
	this_.structCache.Clean()
	this_.structCache_.Clean()
	this_.serviceCache.Clean()
	this_.serviceMethodCache.Clean()
	this_.enumCache.Clean()
	this_.exceptionCache.Clean()
	this_.includePathCache.Clean()

	this_.ServiceList = []*ServiceInfo{}

}

func (this_ *Workspace) Reload() {
	this_.Clean()
	this_.Load()
}
