package datamove

import (
	"errors"
	"fmt"
	"github.com/team-ide/go-tool/elasticsearch"
	"github.com/team-ide/go-tool/kafka"
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

func (this_ *Executor) datasourceToSql(from DataSource) (err error) {
	util.Logger.Info("datasource to sql start")
	to := NewDataSourceSql()
	to.ParamModel = this_.To.GetDialectParam()
	to.OwnerName = this_.To.OwnerName
	to.TableName = this_.To.TableName
	to.ColumnList = this_.To.ColumnList
	to.DialectType = this_.To.DialectType
	to.FilePath = this_.getFilePath("", this_.To.GetFileName(), "sql")
	err = DateMove(this_.Progress, from, to)
	if err != nil {
		util.Logger.Error("datasource to sql error", zap.Error(err))
		return
	}
	util.Logger.Info("datasource to sql end")
	return
}

func (this_ *Executor) datasourceToTxt(from DataSource) (err error) {
	util.Logger.Info("datasource to text start")
	to := NewDataSourceTxt()
	to.ColSeparator = this_.To.ColSeparator
	to.ReplaceCol = this_.To.ReplaceCol
	to.ReplaceLine = this_.To.ReplaceLine
	to.ShouldTrimSpace = this_.To.ShouldTrimSpace
	to.ColumnList = this_.To.ColumnList
	to.FilePath = this_.getFilePath("", this_.To.GetFileName(), this_.To.GetTxtFileType())
	err = DateMove(this_.Progress, from, to)
	if err != nil {
		util.Logger.Error("datasource to text error", zap.Error(err))
		return
	}
	util.Logger.Info("datasource to text end")
	return
}

func (this_ *Executor) datasourceToExcel(from DataSource) (err error) {
	util.Logger.Info("datasource to excel start")
	to := NewDataSourceExcel()
	to.ColumnList = this_.To.ColumnList
	to.ShouldTrimSpace = this_.To.ShouldTrimSpace
	to.FilePath = this_.getFilePath("", this_.To.GetFileName(), "xlsx")
	err = DateMove(this_.Progress, from, to)
	if err != nil {
		util.Logger.Error("datasource to excel error", zap.Error(err))
		return
	}
	util.Logger.Info("datasource to excel end")
	return
}

func (this_ *Executor) datasourceToDb(from DataSource) (err error) {
	util.Logger.Info("datasource to excel start")
	to := NewDataSourceDb()
	to.ColumnList = this_.To.ColumnList
	to.OwnerName = this_.To.OwnerName
	to.TableName = this_.To.TableName
	to.ParamModel = this_.To.GetDialectParam()
	to.Service, err = this_.newDbService(*this_.To.DbConfig, this_.To.Username, this_.To.Password, this_.To.OwnerName)
	if err != nil {
		return
	}
	defer func() { _ = to.Service.GetDb().Close() }()
	err = DateMove(this_.Progress, from, to)
	if err != nil {
		util.Logger.Error("datasource to excel error", zap.Error(err))
		return
	}
	util.Logger.Info("datasource to excel end")
	return
}

func (this_ *Executor) datasourceToEs(from DataSource) (err error) {
	util.Logger.Info("datasource to es start")
	to := NewDataSourceEs()
	to.ColumnList = this_.To.ColumnList
	to.IndexName = this_.To.IndexName
	to.IndexIdName = this_.To.IndexIdName
	to.IndexIdScript = this_.To.IndexIdScript
	to.Service, err = elasticsearch.New(this_.To.EsConfig)
	if err != nil {
		util.Logger.Error("elasticsearch client new error", zap.Error(err))
		return
	}
	err = DateMove(this_.Progress, from, to)
	if err != nil {
		util.Logger.Error("datasource to es error", zap.Error(err))
		return
	}
	util.Logger.Info("datasource to es end")
	return
}

func (this_ *Executor) datasourceToKafka(from DataSource) (err error) {
	util.Logger.Info("datasource to kafka start")
	to := NewDataSourceKafka()
	to.ColumnList = this_.To.ColumnList
	to.TopicName = this_.To.TopicName
	to.TopicGroupName = this_.To.TopicGroupName
	to.TopicKey = this_.To.TopicKey
	to.TopicValue = this_.To.TopicValue
	to.Service, err = kafka.New(this_.To.KafkaConfig)
	if err != nil {
		util.Logger.Error("kafka client new error", zap.Error(err))
		return
	}
	err = DateMove(this_.Progress, from, to)
	if err != nil {
		util.Logger.Error("datasource to kafka error", zap.Error(err))
		return
	}
	util.Logger.Info("datasource to kafka end")
	return
}

func (this_ *Executor) datasourceToRedis(from DataSource) (err error) {
	util.Logger.Info("datasource to redis start")
	to := NewDataSourceKafka()
	to.ColumnList = this_.To.ColumnList
	to.TopicName = this_.To.TopicName
	to.TopicGroupName = this_.To.TopicGroupName
	to.TopicKey = this_.To.TopicKey
	to.TopicValue = this_.To.TopicValue
	to.Service, err = kafka.New(this_.To.KafkaConfig)
	if err != nil {
		util.Logger.Error("kafka client new error", zap.Error(err))
		return
	}
	err = DateMove(this_.Progress, from, to)
	if err != nil {
		util.Logger.Error("datasource to redis error", zap.Error(err))
		return
	}
	util.Logger.Info("datasource to redis end")
	return
}
