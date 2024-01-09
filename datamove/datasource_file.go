package datamove

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"os"
	"strings"
)

type DataSourceFile struct {
	*DataSourceBase
	FilePath string `json:"filePath"`

	readFile  *os.File
	writeFile *os.File
	bufWriter *bufio.Writer
}

func (this_ *DataSourceFile) Stop(progress *Progress) {
	//fmt.Println("stop data source file:" + this_.FilePath)
	this_.CloseReadFile()
	this_.CloseWriteFile()
}

func (this_ *DataSourceFile) ReadStart(progress *Progress) (err error) {
	_, err = this_.GetReadFile()
	return
}

func (this_ *DataSourceFile) ReadEnd(progress *Progress) (err error) {
	this_.CloseReadFile()
	return
}

func (this_ *DataSourceFile) WriteStart(progress *Progress) (err error) {
	_, err = this_.GetWriteFile()
	return
}

func (this_ *DataSourceFile) WriteEnd(progress *Progress) (err error) {
	this_.CloseWriteFile()
	return
}

func (this_ *DataSourceFile) GetReadFile() (file *os.File, err error) {
	file = this_.readFile
	if file != nil {
		return
	}
	//fmt.Println("open read file:" + this_.FilePath)
	file, err = os.Open(this_.FilePath)
	if err != nil {
		err = errors.New("open file [" + this_.FilePath + "] error:" + err.Error())
		return
	}
	this_.readFile = file
	return
}

func (this_ *DataSourceFile) CloseReadFile() {
	//fmt.Println("close read file:" + this_.FilePath)
	file := this_.readFile
	this_.readFile = nil
	if file != nil {
		err := file.Close()
		if err != nil {
			util.Logger.Error("close read file ["+this_.FilePath+"] error", zap.Error(err))
			return
		}
	}
	return
}

func (this_ *DataSourceFile) GetWriteFile() (file *os.File, err error) {
	//fmt.Println("get write file:"+this_.FilePath, this_.writeFile != nil)
	file = this_.writeFile
	if file != nil {
		return
	}

	ex, err := util.PathExists(this_.FilePath)
	if err != nil {
		return
	}
	if !ex {
		file, err = os.Create(this_.FilePath)
		if err != nil {
			err = errors.New("create file [" + this_.FilePath + "] error:" + err.Error())
			return
		}
		_ = file.Close()
	}
	//fmt.Println("open write file:" + this_.FilePath)
	// 打开 只写 创建 追加
	file, err = os.OpenFile(this_.FilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		err = errors.New("create file [" + this_.FilePath + "] error:" + err.Error())
		return
	}
	this_.writeFile = file
	return
}

func (this_ *DataSourceFile) GetWriteBuf() (bufWriter *bufio.Writer, err error) {
	bufWriter = this_.bufWriter
	if bufWriter != nil {
		return
	}
	file, err := this_.GetWriteFile()
	if err != nil {
		return
	}

	bufWriter = bufio.NewWriter(file)
	this_.bufWriter = bufWriter
	return
}

func (this_ *DataSourceFile) CloseWriteFile() {
	bufWriter := this_.bufWriter
	if bufWriter != nil {
		this_.bufWriter = nil
		_ = bufWriter.Flush()
	}
	//fmt.Println("close write file:" + this_.FilePath)
	file := this_.writeFile
	if file != nil {
		this_.writeFile = nil
		err := file.Close()
		if err != nil {
			util.Logger.Error("close write file ["+this_.FilePath+"] error", zap.Error(err))
			return
		}
	}
	return
}

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

func (this_ *DataSourceFile) ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, dropCR(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

func (this_ *DataSourceFile) ReadLineCount() (lineCount int64, err error) {
	file, err := os.Open(this_.FilePath)
	if err != nil {
		err = errors.New("open file [" + this_.FilePath + "] error:" + err.Error())
		return
	}
	defer func() { _ = file.Close() }()
	scanner := bufio.NewScanner(file)
	scanner.Split(this_.ScanLines)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		lineCount++
	}
	util.Logger.Info("txt data source read line count", zap.Any("lineCount", lineCount))
	return
}

func (this_ *DataSourceFile) ReadTitles(progress *Progress) (titles []string, err error) {
	file, err := os.Open(this_.FilePath)
	if err != nil {
		err = errors.New("open file [" + this_.FilePath + "] error:" + err.Error())
		return
	}
	defer func() { _ = file.Close() }()
	scanner := bufio.NewScanner(file)
	scanner.Split(this_.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		cols := strings.Split(line, progress.GetColSeparator())
		for _, c := range cols {
			if progress.ShouldTrimSpace {
				c = strings.TrimSpace(c)
			}
			titles = append(titles, c)
		}
		break
	}
	util.Logger.Info("file data source read titles", zap.Any("titles", titles))
	return
}
