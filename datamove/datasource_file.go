package datamove

import (
	"errors"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"os"
)

type DataSourceFile struct {
	FilePath string `json:"filePath"`

	readFile  *os.File
	writeFile *os.File
}

func (this_ *DataSourceFile) Stop(progress *DateMoveProgress) {
	this_.CloseReadFile()
	this_.CloseWriteFile()
}

func (this_ *DataSourceFile) ReadStart(progress *DateMoveProgress) (err error) {
	_, err = this_.GetReadFile()
	return
}

func (this_ *DataSourceFile) ReadEnd(progress *DateMoveProgress) (err error) {
	this_.CloseReadFile()
	return
}

func (this_ *DataSourceFile) WriteStart(progress *DateMoveProgress) (err error) {
	_, err = this_.GetWriteFile()
	return
}

func (this_ *DataSourceFile) WriteEnd(progress *DateMoveProgress) (err error) {
	this_.CloseWriteFile()
	return
}

func (this_ *DataSourceFile) GetReadFile() (file *os.File, err error) {
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

func (this_ *DataSourceFile) CloseReadFile() {
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

func (this_ *DataSourceFile) CloseWriteFile() {
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
