package kafka

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/team-ide/go-tool/util"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Config kafka配置
type Config struct {
	Address  string `json:"address"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	CertPath string `json:"certPath,omitempty"`
}

// New 创建kafka服务
func New(config Config) (IService, error) {
	service := &Service{
		Config: config,
	}
	err := service.init()
	if err != nil {
		return nil, err
	}
	return service, nil
}

// Service 注册处理器在线信息等
type Service struct {
	Config
}

func (this_ *Service) init() (err error) {
	return
}

func (this_ *Service) Stop() {

}

func (this_ *Service) GetServers() []string {
	var servers []string
	if this_.Address == "" {
		return servers
	}
	if strings.Contains(this_.Address, ",") {
		servers = strings.Split(this_.Address, ",")
	} else if strings.Contains(this_.Address, ";") {
		servers = strings.Split(this_.Address, ";")
	} else {
		servers = []string{this_.Address}
	}
	return servers
}

func (this_ *Service) getClient() (saramaClient sarama.Client, err error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.MaxWaitTime = time.Second * 1

	if this_.Username != "" || this_.Password != "" {
		// sasl认证
		config.Net.SASL.Enable = true
		config.Net.SASL.User = this_.Username
		config.Net.SASL.Password = this_.Password
	}

	if this_.CertPath != "" {
		certPool := x509.NewCertPool()
		var pemCerts []byte
		pemCerts, err = util.ReadFile(this_.CertPath)
		if err != nil {
			return
		}

		if !certPool.AppendCertsFromPEM(pemCerts) {
			err = errors.New("证书[" + this_.CertPath + "]解析失败")
			return
		}
		TLSClientConfig := &tls.Config{
			InsecureSkipVerify: true,
		}
		TLSClientConfig.RootCAs = certPool
		config.Net.TLS.Enable = true
		config.Net.TLS.Config = TLSClientConfig
	}

	saramaClient, err = sarama.NewClient(this_.GetServers(), config)
	if err != nil {
		if saramaClient != nil {
			_ = saramaClient.Close()
		}
		return
	}
	return
}

type Info struct {
	Brokers []*BrokerInfo `json:"brokers"`
}
type BrokerInfo struct {
	Id        int32  `json:"id"`
	Addr      string `json:"addr"`
	Rack      string `json:"rack"`
	Connected bool   `json:"connected"`
}

func (this_ *Service) Info() (res *Info, err error) {
	var saramaClient sarama.Client
	saramaClient, err = this_.getClient()
	if err != nil {
		return
	}
	defer closeSaramaClient(saramaClient)
	brokers := saramaClient.Brokers()

	res = &Info{}
	for _, broker := range brokers {
		brokerInfo := &BrokerInfo{}
		brokerInfo.Id = broker.ID()
		brokerInfo.Addr = broker.Addr()
		brokerInfo.Rack = broker.Rack()
		brokerInfo.Connected, _ = broker.Connected()
	}

	return
}

func closeSaramaClient(saramaClient sarama.Client) {
	if saramaClient != nil {
		_ = saramaClient.Close()
	}
}
func closeClusterAdmin(clusterAdmin sarama.ClusterAdmin) {
	if clusterAdmin != nil {
		_ = clusterAdmin.Close()
	}
}

type TopicInfo struct {
	Topic      string            `json:"topic"`
	Partitions []*TopicPartition `json:"partitions"`
}

type TopicPartition struct {
	Partition int32   `json:"partition"`
	Offset    int64   `json:"offset"`
	Replicas  []int32 `json:"replicas"`
}

func (this_ *Service) GetTopics() (res []*TopicInfo, err error) {
	var saramaClient sarama.Client
	saramaClient, err = this_.getClient()
	if err != nil {
		return
	}
	defer closeSaramaClient(saramaClient)
	topics, err := saramaClient.Topics()
	if err != nil {
		return
	}

	sort.Strings(topics)

	for _, topic := range topics {
		info := &TopicInfo{
			Topic: topic,
		}
		ps, _ := saramaClient.Partitions(topic)
		for _, p := range ps {
			partition := &TopicPartition{
				Partition: p,
			}
			info.Partitions = append(info.Partitions, partition)
		}

		res = append(res, info)
	}

	return
}

func (this_ *Service) GetTopic(topic string, time int64) (res *TopicInfo, err error) {
	var saramaClient sarama.Client
	saramaClient, err = this_.getClient()
	if err != nil {
		return
	}
	defer closeSaramaClient(saramaClient)

	res = &TopicInfo{
		Topic: topic,
	}
	ps, _ := saramaClient.Partitions(topic)
	for _, p := range ps {
		partition := &TopicPartition{
			Partition: p,
		}
		partition.Offset, _ = saramaClient.GetOffset(topic, p, time)
		partition.Replicas, _ = saramaClient.Replicas(topic, p)
		res.Partitions = append(res.Partitions, partition)
	}

	return
}

func (this_ *Service) Pull(groupId string, topics []string, PullSize int, PullTimeout int, keyType, valueType string) (msgList []*Message, err error) {
	if PullSize <= 0 {
		PullSize = 10
	}
	if PullTimeout <= 0 {
		PullTimeout = 1000
	}
	var saramaClient sarama.Client
	saramaClient, err = this_.getClient()
	if err != nil {
		return
	}
	defer closeSaramaClient(saramaClient)
	group, err := sarama.NewConsumerGroupFromClient(groupId, saramaClient)
	if err != nil {
		return
	}
	handler := &consumerGroupHandler{
		size: PullSize,
	}
	go func() {
		ctx := context.Background()
		err = group.Consume(ctx, topics, handler)

		if err != nil {
			fmt.Println("group.Consume error:", err)
		}
	}()
	startTime := util.GetNowTime()
	for {
		time.Sleep(100 * time.Millisecond)
		nowTime := util.GetNowTime()
		if handler.appended || nowTime-startTime >= int64(PullTimeout) {
			break
		}
	}
	err = group.Close()
	if err != nil {
		fmt.Println("group.Close error:", err)
		return
	}
	for _, one := range handler.messages {
		var msg *Message
		msg, err = ConsumerMessageToMessage(keyType, valueType, one)
		if err != nil {
			return
		}
		msgList = append(msgList, msg)
	}
	return
}

type consumerGroupHandler struct {
	messages []*sarama.ConsumerMessage
	appended bool
	size     int
}

func (*consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (*consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (handler *consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	if sess == nil {
		return nil
	}
	chanMessages := claim.Messages()
	for msg := range chanMessages {
		handler.messages = append(handler.messages, msg)
		if len(handler.messages) >= handler.size {
			break
		}
	}
	handler.appended = true
	return nil
}

func (this_ *Service) MarkOffset(groupId string, topic string, partition int32, offset int64) (err error) {
	var saramaClient sarama.Client
	saramaClient, err = this_.getClient()
	if err != nil {
		return
	}
	defer closeSaramaClient(saramaClient)
	offsetManager, err := sarama.NewOffsetManagerFromClient(groupId, saramaClient)
	if err != nil {
		return
	}
	partitionOffsetManager, err := offsetManager.ManagePartition(topic, partition)
	if err != nil {
		return
	}
	partitionOffsetManager.MarkOffset(offset, "")
	err = offsetManager.Close()
	return
}

func (this_ *Service) ResetOffset(groupId string, topic string, partition int32, offset int64) (err error) {
	var saramaClient sarama.Client
	saramaClient, err = this_.getClient()
	if err != nil {
		return
	}
	defer closeSaramaClient(saramaClient)
	offsetManager, err := sarama.NewOffsetManagerFromClient(groupId, saramaClient)
	if err != nil {
		return
	}
	partitionOffsetManager, err := offsetManager.ManagePartition(topic, partition)
	if err != nil {
		return
	}
	defer func() {
		_ = offsetManager.Close()
	}()
	partitionOffsetManager.ResetOffset(offset, "")
	return
}

func (this_ *Service) CreatePartitions(topic string, count int32) (err error) {
	var saramaClient sarama.Client
	saramaClient, err = this_.getClient()
	if err != nil {
		return
	}
	defer closeSaramaClient(saramaClient)
	admin, err := sarama.NewClusterAdminFromClient(saramaClient)
	if err != nil {
		return
	}

	defer closeClusterAdmin(admin)

	err = admin.CreatePartitions(topic, count, nil, false)

	return
}

func (this_ *Service) CreateTopic(topic string, numPartitions int32, replicationFactor int16) (err error) {
	var saramaClient sarama.Client
	saramaClient, err = this_.getClient()
	if err != nil {
		return
	}
	defer closeSaramaClient(saramaClient)
	admin, err := sarama.NewClusterAdminFromClient(saramaClient)
	if err != nil {
		return
	}

	defer closeClusterAdmin(admin)
	if numPartitions <= 0 {
		numPartitions = 1
	}
	if replicationFactor <= 0 {
		replicationFactor = 1
	}
	detail := &sarama.TopicDetail{
		NumPartitions:     numPartitions,
		ReplicationFactor: replicationFactor,
	}
	err = admin.CreateTopic(topic, detail, false)

	return
}

func (this_ *Service) DeleteTopic(topic string) (err error) {
	var saramaClient sarama.Client
	saramaClient, err = this_.getClient()
	if err != nil {
		return
	}
	defer closeSaramaClient(saramaClient)
	admin, err := sarama.NewClusterAdminFromClient(saramaClient)
	if err != nil {
		return
	}

	defer closeClusterAdmin(admin)

	err = admin.DeleteTopic(topic)

	return
}

func (this_ *Service) DeleteConsumerGroup(groupId string) (err error) {
	var saramaClient sarama.Client
	saramaClient, err = this_.getClient()
	if err != nil {
		return
	}
	defer closeSaramaClient(saramaClient)
	admin, err := sarama.NewClusterAdminFromClient(saramaClient)
	if err != nil {
		return
	}

	defer closeClusterAdmin(admin)

	err = admin.DeleteConsumerGroup(groupId)

	return
}

func (this_ *Service) DeleteRecords(topic string, partitionOffsets map[int32]int64) (err error) {
	var saramaClient sarama.Client
	saramaClient, err = this_.getClient()
	if err != nil {
		return
	}
	defer closeSaramaClient(saramaClient)
	admin, err := sarama.NewClusterAdminFromClient(saramaClient)
	if err != nil {
		return
	}

	defer closeClusterAdmin(admin)

	err = admin.DeleteRecords(topic, partitionOffsets)

	return
}

// GetOffset 查询集群以获取
// 主题/分区组合上的给定时间（以毫秒为单位）。
// 对于最早的可用偏移，时间应该是OffsetOldest，
// OffsetNewest是下一次或某一时间将生成的消息的偏移量。
func (this_ *Service) GetOffset(topic string, partitionID int32, time int64) (offset int64, err error) {
	var saramaClient sarama.Client
	saramaClient, err = this_.getClient()
	if err != nil {
		return
	}
	defer closeSaramaClient(saramaClient)

	offset, err = saramaClient.GetOffset(topic, partitionID, time)

	return
}

func (this_ *Service) Partitions(topic string) (partitions []int32, err error) {
	var saramaClient sarama.Client
	saramaClient, err = this_.getClient()
	if err != nil {
		return
	}
	defer closeSaramaClient(saramaClient)

	partitions, err = saramaClient.Partitions(topic)

	return
}

// NewSyncProducer 创建生产者
func (this_ *Service) NewSyncProducer() (syncProducer sarama.SyncProducer, err error) {

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 3

	if this_.Username != "" || this_.Password != "" {
		// sasl认证
		config.Net.SASL.Enable = true
		config.Net.SASL.User = this_.Username
		config.Net.SASL.Password = this_.Password
	}

	if this_.CertPath != "" {
		certPool := x509.NewCertPool()
		var pemCerts []byte
		pemCerts, err = util.ReadFile(this_.CertPath)
		if err != nil {
			return
		}

		if !certPool.AppendCertsFromPEM(pemCerts) {
			err = errors.New("证书[" + this_.CertPath + "]解析失败")
			return
		}
		TLSClientConfig := &tls.Config{
			InsecureSkipVerify: true,
		}
		TLSClientConfig.RootCAs = certPool
		config.Net.TLS.Enable = true
		config.Net.TLS.Config = TLSClientConfig
	}

	syncProducer, err = sarama.NewSyncProducer(this_.GetServers(), config)
	if err != nil {
		if syncProducer != nil {
			_ = syncProducer.Close()
		}
		return
	}
	return
}

type ProducerMessage struct {
	*sarama.ProducerMessage
}

// Push 推送消息到kafka
func (this_ *Service) Push(msg *Message) (err error) {
	producerMessage, err := MessageToProducerMessage(msg)
	if err != nil {
		return
	}
	syncProducer, err := this_.NewSyncProducer()
	if err != nil {
		return
	}
	defer func() {
		_ = syncProducer.Close()
	}()

	_, _, err = syncProducer.SendMessage(producerMessage)
	return err
}

func MessageToProducerMessage(msg *Message) (producerMessage *sarama.ProducerMessage, err error) {
	var key sarama.Encoder
	var value sarama.Encoder
	if msg.Key != "" {
		if strings.ToLower(msg.KeyType) == "long" {
			longV, err := strconv.ParseInt(msg.Key, 10, 64)
			if err != nil {
				return nil, err
			}
			var bytes = make([]byte, 8)
			binary.BigEndian.PutUint64(bytes, uint64(longV))
			key = sarama.ByteEncoder(bytes)
		} else {
			key = sarama.ByteEncoder(msg.Key)
		}
	}
	if msg.Value != "" {
		if strings.ToLower(msg.ValueType) == "long" {
			longV, err := strconv.ParseInt(msg.Value, 10, 64)
			if err != nil {
				return nil, err
			}
			var bytes = make([]byte, 8)
			binary.BigEndian.PutUint64(bytes, uint64(longV))
			value = sarama.ByteEncoder(bytes)
		} else {
			value = sarama.ByteEncoder(msg.Value)
		}
	}

	producerMessage = &sarama.ProducerMessage{}
	producerMessage.Topic = msg.Topic
	producerMessage.Key = key
	producerMessage.Value = value
	if msg.Timestamp == nil || (*msg.Timestamp).IsZero() {
		producerMessage.Timestamp = time.Now()
	} else {
		producerMessage.Timestamp = *msg.Timestamp
	}
	if msg.Partition != nil {
		producerMessage.Partition = *msg.Partition
	}
	if msg.Offset != nil {
		producerMessage.Offset = *msg.Offset
	}
	if msg.Headers != nil {
		for _, one := range msg.Headers {
			producerMessage.Headers = append(producerMessage.Headers, sarama.RecordHeader{
				Key:   []byte(one.Key),
				Value: []byte(one.Value),
			})
		}
	}
	return
}

func ConsumerMessageToMessage(keyType string, valueType string, consumerMessage *sarama.ConsumerMessage) (msg *Message, err error) {
	var key string
	var value string

	if consumerMessage.Key != nil && len(consumerMessage.Key) > 0 {
		if len(consumerMessage.Key) == 8 {
			Uint64Key := binary.BigEndian.Uint64(consumerMessage.Key)
			int64Key := int64(Uint64Key)
			if int64Key >= 0 {
				key = strconv.FormatInt(int64Key, 10)
			}
		}
		if key == "" {
			key = string(consumerMessage.Key)
		}
	}
	if consumerMessage.Value != nil && len(consumerMessage.Value) > 0 {
		if len(consumerMessage.Value) == 8 {
			Uint64Value := binary.BigEndian.Uint64(consumerMessage.Value)
			int64Value := int64(Uint64Value)
			if int64Value >= 0 {
				value = strconv.FormatInt(int64Value, 10)
			}
		}
		if value == "" {
			value = string(consumerMessage.Value)
		}
	}
	msg = &Message{
		Key:       key,
		Value:     value,
		Topic:     consumerMessage.Topic,
		Partition: &consumerMessage.Partition,
		Offset:    &consumerMessage.Offset,
	}
	if consumerMessage.Headers != nil {
		for _, header := range consumerMessage.Headers {
			msg.Headers = append(msg.Headers, MessageHeader{Key: string(header.Key), Value: string(header.Value)})
		}
	}
	return
}

type MessageHeader struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type Message struct {
	KeyType   string          `json:"keyType,omitempty"`
	Key       string          `json:"key,omitempty"`
	ValueType string          `json:"valueType,omitempty"`
	Value     string          `json:"value,omitempty"`
	Topic     string          `json:"topic,omitempty"`
	Partition *int32          `json:"partition,omitempty"`
	Offset    *int64          `json:"offset,omitempty"`
	Headers   []MessageHeader `json:"headers,omitempty"`
	Timestamp *time.Time      `json:"timestamp,omitempty"`
}
