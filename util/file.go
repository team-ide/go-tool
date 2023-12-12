package util

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var (
	rootDir = getRootDir()
)

func getRootDir() string {
	dir, _ := os.Getwd()
	dir = FormatPath(dir)
	return dir
}

// GetRootDir 获取当前程序根路径
func GetRootDir() string {
	return rootDir
}

// FormatPath 格式化路径
// FormatPath("/x/x/xxx\xx\xx")
func FormatPath(path string) string {

	var abs string
	abs, err := filepath.Abs(path)
	if err != nil {
		return path
	}
	res := filepath.ToSlash(abs)
	return res
}

// GetAbsolutePath 获取路径觉得路径
// GetAbsolutePath("/x/x/xxx\xx\xx")
func GetAbsolutePath(path string) (absolutePath string) {
	var abs string
	abs, _ = filepath.Abs(path)

	absolutePath = filepath.ToSlash(abs)
	return
}

// PathExists 路径文件是否存在
// PathExists("/x/x/xxx\xx\xx")
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// LoadDirFiles 加载目录下文件 读取文件内容（key为文件名为相对路径）
// LoadDirFiles("/x/x/xxx\xx\xx")
func LoadDirFiles(dir string) (fileMap map[string][]byte, err error) {
	fileMap = map[string][]byte{}
	var exist bool
	exist, err = PathExists(dir)
	if err != nil {
		return
	}
	if !exist {
		return
	}

	formatDir := FormatPath(dir)
	//获取当前目录下的所有文件或目录信息
	err = filepath.Walk(formatDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {

		} else {
			var abs = FormatPath(path)
			name := strings.TrimPrefix(abs, formatDir)
			name = strings.TrimPrefix(name, "/")
			var f *os.File
			f, err = os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()
			var bs []byte
			bs, err = io.ReadAll(f)
			if err != nil {
				return err
			}
			fileMap[name] = bs
		}
		return nil
	})
	return
}

// LoadDirFilenames 加载目录下文件（文件名为相对路径）
// LoadDirFilenames("/x/x/xxx\xx\xx")
func LoadDirFilenames(dir string) (filenames []string, err error) {
	var exist bool
	exist, err = PathExists(dir)
	if err != nil {
		return
	}
	if !exist {
		return
	}
	formatDir := FormatPath(dir)
	err = filepath.Walk(formatDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {

		} else {
			var abs = FormatPath(path)
			name := strings.TrimPrefix(abs, formatDir)
			name = strings.TrimPrefix(name, "/")
			filenames = append(filenames, name)
		}
		return nil
	})
	sort.Slice(filenames, func(i, j int) bool {
		return strings.ToLower(filenames[i]) < strings.ToLower(filenames[j]) //升序  即前面的值比后面的小 忽略大小写排序
	})
	return
}

// ReadFile 读取文件内容 返回 []byte
// ReadFile("/x/x/xxx\xx\xx")
func ReadFile(filename string) (bs []byte, err error) {
	var f *os.File
	var exists bool
	exists, err = PathExists(filename)
	if err != nil {
		return
	}
	if !exists {
		return
	} else {
		f, err = os.Open(filename)
	}
	if err != nil {
		return
	}
	defer f.Close()
	bs, err = io.ReadAll(f)
	if err != nil {
		return
	}
	return
}

// ReadFileString 读取文件内容 返回字符串
// ReadFileString("/x/x/xxx\xx\xx")
func ReadFileString(filename string) (str string, err error) {
	bs, err := ReadFile(filename)
	if err != nil {
		return
	}
	str = string(bs)
	return
}

// StringToBytes 字符串转为 []byte
// StringToBytes("这是文本")
func StringToBytes(str string) []byte {
	return []byte(str)
}

// WriteFile 写入文件内容,
// WriteFile("/x/x/xxx\xx\xx", StringToBytes("这是文本"))
func WriteFile(filename string, bs []byte) (err error) {
	f, err := os.Create(filename)
	if err != nil {
		return
	}
	defer func() { _ = f.Close() }()
	_, err = f.Write(bs)
	if err != nil {
		return
	}
	return
}

// WriteFileString 写入文件内容
// WriteFileString("/x/x/xxx\xx\xx", "这是文本")
func WriteFileString(filename string, str string) (err error) {
	return WriteFile(filename, []byte(str))
}

// ReadLine 逐行读取文件
// ReadLine("/x/x/xxx\xx\xx")
func ReadLine(filename string) (lines []string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer func() { _ = f.Close() }()
	buf := bufio.NewReader(f)
	var line string
	for {
		line, err = buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF { //读取结束，会报EOF
				err = nil
				break
			}
			return nil, err
		}
		lines = append(lines, line)
	}
	return
}

// IsSubPath child是否是parent子路径
// IsSubPath("/a/b", "/a/b/c")
func IsSubPath(parent, child string) (isSub bool, err error) {
	parentPath, err := filepath.Abs(parent)
	if err != nil {
		return
	}
	parentPath = filepath.ToSlash(parentPath)
	if !strings.HasSuffix(parentPath, "/") {
		parentPath += "/"
	}
	childPath, err := filepath.Abs(child)
	if err != nil {
		return
	}
	childPath = filepath.ToSlash(childPath)
	isSub = strings.HasPrefix(childPath, parentPath)
	return
}

type DirInfo struct {
	Path          string         `json:"path"`
	Name          string         `json:"name"`
	FileSize      int64          `json:"fileSize,omitempty"`
	AllFileSize   int64          `json:"allFileSize,omitempty"`
	DirNumber     int            `json:"dirNumber,omitempty"`
	FileNumber    int            `json:"fileNumber,omitempty"`
	AllFileNumber int            `json:"allFileNumber,omitempty"`
	AllDirNumber  int            `json:"allDirNumber,omitempty"`
	FileInfos     []*DirFileInfo `json:"fileInfos,omitempty"`
	DirInfos      []*DirInfo     `json:"dirInfos,omitempty"`
}
type DirFileInfo struct {
	Path     string `json:"path"`
	Name     string `json:"name"`
	FileSize int64  `json:"fileSize"`
}

// LoadDirInfo 加载目录信息，目录下文件目录数量大小等，可以配置加载所有子目录
// LoadDirInfo("/x/x/", true)
func LoadDirInfo(dir string, loadSubDir bool) (dirInfo *DirInfo, err error) {

	var path = FormatPath(dir)
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	info, err := os.Stat(path)
	if err != nil {
		return
	}
	if !info.IsDir() {
		return
	}
	dirInfo = &DirInfo{}
	dirInfo.Name = info.Name()
	dirInfo.Path = path

	fs, err := os.ReadDir(path)
	if err != nil {
		return
	}
	for _, f := range fs {
		fPath := path + f.Name()
		if f.IsDir() {
			dirInfo.DirNumber += 1
			if !loadSubDir {
				continue
			}
			subDir, _ := LoadDirInfo(fPath, loadSubDir)
			if subDir != nil {
				dirInfo.DirInfos = append(dirInfo.DirInfos, subDir)
				dirInfo.AllFileNumber += subDir.AllFileNumber
				dirInfo.AllDirNumber += subDir.AllDirNumber
				dirInfo.AllFileSize += subDir.AllFileSize
			}
		} else {
			fInfo, _ := f.Info()
			if fInfo == nil {
				continue
			}
			fileInfo := &DirFileInfo{}
			dirInfo.FileInfos = append(dirInfo.FileInfos, fileInfo)
			fileInfo.FileSize = fInfo.Size()
			fileInfo.Path = fPath
			fileInfo.Name = f.Name()
			dirInfo.FileNumber += 1
			dirInfo.FileSize += fileInfo.FileSize
			dirInfo.AllFileNumber += 1
			dirInfo.AllFileSize += fileInfo.FileSize
		}
	}

	return
}
