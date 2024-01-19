package datamove

import (
	"errors"
	"fmt"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"sync"
)

func DateMove(progress *Progress, from DataSource, to DataSource) (err error) {
	progress.dataMoveStop = nil
	defer func() {
		if x := recover(); x != nil {
			err = errors.New(fmt.Sprintf("%s", x))
		}
		progress.dataMoveStop = nil
	}()
	var dataChan = make(chan *Data, 1)

	var stopWait sync.WaitGroup
	stopWait.Add(2)
	go func() {
		defer stopWait.Done()
		e := startRead(progress, from, dataChan)
		if e != nil {
			err = e
		}
	}()
	go func() {
		defer stopWait.Done()
		e := startWrite(progress, to, dataChan)
		if e != nil {
			err = e
		}
	}()
	util.Logger.Info("data move wait start")
	stopWait.Wait()
	util.Logger.Info("data move wait end")
	return
}

func startRead(progress *Progress, from DataSource, dataChan chan *Data) (err error) {
	util.Logger.Info("read data source start")
	defer func() {
		if x := recover(); x != nil {
			err = errors.New(fmt.Sprintf("%s", x))
		}
		close(dataChan)

		if e := from.ReadEnd(progress); e != nil {
			if err == nil {
				err = e
			}
		}
		from.Stop(progress)
		util.Logger.Info("read data source end")
	}()
	if err = from.ReadStart(progress); err != nil {
		return
	}
	err = from.Read(progress, dataChan)
	return
}

func startWrite(progress *Progress, to DataSource, dataChan chan *Data) (err error) {
	util.Logger.Info("write data source start")
	defer func() {
		if x := recover(); x != nil {
			err = errors.New(fmt.Sprintf("%s", x))
		}

		if e := to.WriteEnd(progress); e != nil {
			if err == nil {
				err = e
			}
		}
		to.Stop(progress)
		// 防止管道未消费
		for {
			_, ok := <-dataChan
			if !ok {
				break
			}
		}
		util.Logger.Info("write data source end")
	}()
	if err = to.WriteStart(progress); err != nil {
		return
	}
	for {
		data, ok := <-dataChan
		if !ok {
			break
		}
		if err != nil || data == nil {
			continue
		}

		util.Logger.Info("write data source start", zap.Any("total", data.Total))
		err = to.Write(progress, data)
		util.Logger.Info("write data source end", zap.Any("total", data.Total))
		if err != nil {
			if !progress.ErrorContinue || errors.Is(err, ErrorStopped) {
				dataMoveStop := new(bool)
				*dataMoveStop = true
				progress.dataMoveStop = dataMoveStop
				continue
			}
			err = nil
		}
	}
	return
}
