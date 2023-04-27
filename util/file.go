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
	rootDir string
)

func init() {
	rootDir, _ = os.Getwd()
	rootDir = FormatPath(rootDir)
}

// GetRootDir 获取当前程序根路径
func GetRootDir() string {
	return rootDir
}

// FormatPath 格式化路径
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
func GetAbsolutePath(path string) (absolutePath string) {
	var abs string
	abs, _ = filepath.Abs(path)

	absolutePath = filepath.ToSlash(abs)
	return
}

// PathExists 路径文件是否存在
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
func ReadFileString(filename string) (str string, err error) {
	bs, err := ReadFile(filename)
	if err != nil {
		return
	}
	str = string(bs)
	return
}

// WriteFile 写入文件内容
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
func WriteFileString(filename string, str string) (err error) {
	return WriteFile(filename, []byte(str))
}

// ReadLine 逐行读取文件
func ReadLine(filename string) (lines []string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	buf := bufio.NewReader(f)
	var line string
	for {
		line, err = buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF { //读取结束，会报EOF
				err = nil
				return
			}
			return nil, err
		}
		lines = append(lines, line)
	}
	return
}
