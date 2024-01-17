package datamove

import (
	"context"
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/team-ide/go-tool/kafka"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"time"
)

func NewDataSourceKafka() *DataSourceKafka {
	return &DataSourceKafka{
		DataSourceBase: &DataSourceBase{},
	}
}

type DataSourceKafka struct {
	*DataSourceBase
	TopicName        string `json:"topicName"`
	TopicGroupName   string `json:"topicGroupName"`
	TopicKey         string `json:"topicKey"`
	TopicValue       string `json:"topicValue"`
	TopicValueByData bool   `json:"topicValueByData"`
	PullWait         int64  `json:"pullWait"`

	Service kafka.IService

	lastData *Data
}

func (this_ *DataSourceKafka) Stop(progress *Progress) {

}

func (this_ *DataSourceKafka) ReadStart(progress *Progress) (err error) {

	return
}

func (this_ *DataSourceKafka) Read(progress *Progress, dataChan chan *Data) (err error) {

	client, err := this_.Service.GetClient()
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	defer func() { _ = client.Close() }()
	group, err := sarama.NewConsumerGroupFromClient(this_.TopicGroupName, client)
	if err != nil {
		return
	}

	var topics []string
	topics = append(topics, this_.TopicName)
	ctx, cancel := context.WithCancel(context.Background())
	handler := &consumerGroupHandler{
		DataSourceKafka: this_,
		Progress:        progress,
		dataChan:        dataChan,
		cancel:          cancel,
	}
	var isConsumeEnd bool
	go func() {
		defer func() { _ = group.Close() }()
		util.Logger.Info("kafka pull start", zap.Any("topics", topics), zap.Any("groupId", this_.TopicGroupName))
		err = group.Consume(ctx, topics, handler)
		isConsumeEnd = true
		util.Logger.Info("kafka pull end", zap.Any("topics", topics), zap.Any("groupId", this_.TopicGroupName))
		if err != nil {
			util.Logger.Error("group consume error", zap.Error(err))
		}
	}()
	var pullWait = this_.PullWait
	if pullWait <= 0 {
		pullWait = 10
	}
	pullWait = pullWait * 1000
	var isTimeout bool
	for {
		time.Sleep(time.Second)
		if isConsumeEnd {
			break
		}
		if isTimeout {
			continue
		}
		nowTime := time.Now().UnixMilli()
		nowWait := nowTime - handler.lastTime
		if nowWait > pullWait {
			if cancel != nil {
				cancel()
			}
			isTimeout = true
		}
	}
	if cancel != nil {
		cancel()
	}
	return
}

type consumerGroupHandler struct {
	*DataSourceKafka
	*Progress
	dataChan chan *Data
	lastTime int64
	cancel   context.CancelFunc
}

func (*consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (*consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (this_ *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	if session == nil || claim == nil {
		return nil
	}
	chanMessages := claim.Messages()
	var err error
	for {
		this_.lastTime = time.Now().UnixMilli()
		if this_.ShouldStop() {
			break
		}
		select {
		case msg := <-chanMessages:
			if msg == nil {
				break
			}
			err = this_.onMsg(msg)
		}
		if err != nil {
			break
		}
	}
	if this_.cancel != nil {
		this_.cancel()
	}
	if this_.lastData != nil && this_.lastData.Total > 0 {
		this_.dataChan <- this_.lastData
	}
	return nil
}

func (this_ *consumerGroupHandler) onMsg(msg *sarama.ConsumerMessage) (err error) {
	pageSize := this_.Progress.BatchNumber
	data := map[string]interface{}{}

	for _, h := range msg.Headers {
		name := string(h.Key)
		data[name] = string(h.Value)
	}
	if msg.Key != nil {
		data[this_.TopicKey] = string(msg.Key)
	}
	if msg.Value != nil {
		data[this_.TopicValue] = string(msg.Value)
		if msg.Value != nil {
			_ = util.JSONDecodeUseNumber(msg.Value, &data)
		}
	}

	values, e := this_.DataToValues(this_.Progress, data)
	if e != nil {
		this_.Progress.ReadCount.AddError(1, e)
		if !this_.Progress.ErrorContinue {
			err = e
			return
		}
	} else {
		if this_.lastData == nil {
			this_.lastData = &Data{
				DataType: DataTypeCols,
			}
		}
		this_.lastData.ColsList = append(this_.lastData.ColsList, values)
		this_.lastData.Total++
		this_.Progress.ReadCount.AddSuccess(1)
		if this_.lastData.Total >= pageSize {
			this_.dataChan <- this_.lastData
			this_.lastData = &Data{
				DataType: DataTypeCols,
			}
		}
	}
	return
}

func (this_ *DataSourceKafka) ReadEnd(progress *Progress) (err error) {
	return
}

func (this_ *DataSourceKafka) WriteStart(progress *Progress) (err error) {

	return
}

func (this_ *DataSourceKafka) Write(progress *Progress, data *Data) (err error) {

	switch data.DataType {
	case DataTypeCols:
		data.Total = int64(len(data.ColsList))
		if data.Total > 0 {
			for _, cols := range data.ColsList {
				d, e := this_.ValuesToData(progress, cols)
				if e != nil {
					progress.WriteCount.AddError(1, e)
					if !progress.ErrorContinue {
						err = e
						return
					}
				} else {
					var key string
					var value string
					if this_.TopicKey != "" {
						key = util.GetStringValue(d[this_.TopicKey])
					}
					if this_.TopicValueByData {
						value = util.GetStringValue(d)
					} else if this_.TopicValue != "" {
						value = util.GetStringValue(d[this_.TopicValue])
					}
					msg := &kafka.Message{}
					msg.Key = key
					msg.Value = value
					e = this_.Service.Push(msg)
					if e != nil {
						progress.WriteCount.AddError(1, e)
						if !progress.ErrorContinue {
							err = e
							return
						}
					} else {
						progress.WriteCount.AddSuccess(1)
					}
				}

			}

		}
		break
	default:
		err = errors.New(fmt.Sprintf("当前数据类型[%d]，不支持写入", data.DataType))
		return
	}
	return

}

func (this_ *DataSourceKafka) WriteEnd(progress *Progress) (err error) {
	return
}
