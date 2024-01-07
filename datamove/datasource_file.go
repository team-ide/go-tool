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
	DataSourceBase
	FilePath     string `json:"filePath"`
	RowSeparator byte   `json:"rowSeparator"` // 行 分隔符 默认 `\n`
	ColSeparator string `json:"colSeparator"` // 列 分隔符 默认 `,`

	readFile  *os.File
	writeFile *os.File
}

func (this_ DataSourceFile) Stop(progress *DateMoveProgress) {
	this_.CloseReadFile()
	this_.CloseWriteFile()
}

func (this_ DataSourceFile) ReadStart(progress *DateMoveProgress) (err error) {
	_, err = this_.GetReadFile()
	return
}

func (this_ DataSourceFile) ReadEnd(progress *DateMoveProgress) (err error) {
	this_.CloseReadFile()
	return
}

func (this_ DataSourceFile) WriteStart(progress *DateMoveProgress) (err error) {
	_, err = this_.GetWriteFile()
	return
}

func (this_ DataSourceFile) WriteEnd(progress *DateMoveProgress) (err error) {
	this_.CloseWriteFile()
	return
}

func (this_ DataSourceFile) GetReadFile() (file *os.File, err error) {
	file = this_.readFile
	if file != nil {
		return
	}
	file, err = os.Open(this_.FilePath)
	if err != nil {
		err = errors.New("open file [" + this_.FilePath + "] error:" + err.Error())
		return
	}
	this_.readFile = file
	return
}

func (this_ DataSourceFile) CloseReadFile() {
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

func (this_ DataSourceFile) GetWriteFile() (file *os.File, err error) {
	file = this_.writeFile
	if file != nil {
		return
	}
	file, err = os.Create(this_.FilePath)
	if err != nil {
		err = errors.New("create file [" + this_.FilePath + "] error:" + err.Error())
		return
	}
	this_.writeFile = file
	return
}

func (this_ DataSourceFile) CloseWriteFile() {
	file := this_.writeFile
	this_.writeFile = nil
	if file != nil {
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

func (this_ DataSourceFile) GetRowSeparator() byte {
	if this_.RowSeparator == 0 {
		return '\n'
	}
	return this_.RowSeparator
}

func (this_ DataSourceFile) GetColSeparator() string {
	if this_.ColSeparator == "" {
		return ","
	}
	return this_.ColSeparator
}

func (this_ DataSourceFile) ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	rowSeparator := this_.GetRowSeparator()
	if i := bytes.IndexByte(data, rowSeparator); i >= 0 {
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

func (this_ DataSourceFile) ReadLineCount() (lineCount int64, err error) {
	file, err := os.Open(this_.FilePath)
	if err != nil {
		err = errors.New("open file [" + this_.FilePath + "] error:" + err.Error())
		return
	}
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

func (this_ DataSourceFile) ReadTitles() (titles []string, err error) {
	file, err := os.Open(this_.FilePath)
	if err != nil {
		err = errors.New("open file [" + this_.FilePath + "] error:" + err.Error())
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(this_.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		cols := strings.Split(line, this_.GetColSeparator())
		for _, c := range cols {
			if this_.ShouldTrimSpace {
				c = strings.TrimSpace(c)
			}
			titles = append(titles, c)
		}
		break
	}
	util.Logger.Info("file data source read titles", zap.Any("titles", titles))
	return
}
