package datamove

import (
	"errors"
	"fmt"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"time"
)

func DateMove(param *Param, from DataSource, to DataSource, onProgress func(progress *DateMoveProgress)) (err error) {
	progress := &DateMoveProgress{
		Param:    param,
		Read:     &DateMoveProgressInfo{},
		Write:    &DateMoveProgressInfo{},
		callback: onProgress,
	}
	defer func() {
		if x := recover(); x != nil {
			err = errors.New(fmt.Sprintf("%s", x))
		}
		onProgress(progress)
	}()
	progress.dataChan = make(chan *Data, 1)

	progress.stopWait.Add(2)
	go func() {
		defer progress.stopWait.Done()
		e := startRead(progress, from)
		if e != nil {
			err = e
		}
	}()
	go func() {
		defer progress.stopWait.Done()
		e := startWrite(progress, to)
		if e != nil {
			err = e
		}
	}()
	util.Logger.Info("data move wait start")
	progress.stopWait.Wait()
	util.Logger.Info("data move wait end")
	return
}

func startRead(progress *DateMoveProgress, from DataSource) (err error) {
	util.Logger.Info("read data source start")
	defer func() {
		if x := recover(); x != nil {
			err = errors.New(fmt.Sprintf("%s", x))
		}
		progress.isEnd = true
		close(progress.dataChan)

		progress.Read.EndTime = time.Now()
		if e := from.ReadEnd(progress); e != nil {
			if err == nil {
				err = e
			}
		}
		from.Stop(progress)
		util.Logger.Info("read data source end")
	}()
	progress.Read.StartTime = time.Now()
	if err = from.ReadStart(progress); err != nil {
		return
	}
	err = from.Read(progress, progress.dataChan)
	return
}

func startWrite(progress *DateMoveProgress, to DataSource) (err error) {
	util.Logger.Info("write data source start")
	defer func() {
		if x := recover(); x != nil {
			err = errors.New(fmt.Sprintf("%s", x))
		}
		progress.isEnd = true

		progress.Write.EndTime = time.Now()
		if e := to.WriteEnd(progress); e != nil {
			if err == nil {
				err = e
			}
		}
		to.Stop(progress)
		// 防止管道未消费
		for {
			_, ok := <-progress.dataChan
			if !ok {
				break
			}
		}
		util.Logger.Info("write data source end")
	}()
	progress.Write.StartTime = time.Now()
	if err = to.WriteStart(progress); err != nil {
		return
	}
	for {
		data, ok := <-progress.dataChan
		if !ok {
			break
		}
		if err != nil || data == nil || data.Total <= 0 {
			continue
		}

		util.Logger.Info("write data source", zap.Any("total", data.Total))
		err = to.Write(progress, data)
		progress.callback(progress)
		if err != nil {
			progress.Write.Errors = append(progress.Write.Errors, err.Error())
			if !progress.ErrorContinue || errors.Is(err, ErrorStopped) {
				progress.isEnd = true
				continue
			}
			err = nil
		}
	}
	return
}
