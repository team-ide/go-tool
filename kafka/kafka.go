package kafka

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"errors"
	"github.com/Shopify/sarama"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Service 注册处理器在线信息等
type Service struct {
	*Config
}

func (this_ *Service) init() (err error) {
	return
}

func (this_ *Service) Close() {
	if this_ == nil {
		return
	}

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

	sort.Slice(topics, func(i, j int) bool {
		return strings.ToLower(topics[i]) < strings.ToLower(topics[j]) //升序  即前面的值比后面的小 忽略大小写排序
	})

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
	defer func() {
		e := group.Close()
		if e != nil {
			util.Logger.Error("group close error", zap.Error(e))
		}
	}()
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(PullTimeout))
	handler := &consumerGroupHandler{
		size:   PullSize,
		cancel: cancel,
	}
	util.Logger.Info("kafka pull start", zap.Any("topics", topics), zap.Any("groupId", groupId), zap.Any("timeout", PullTimeout))
	err = group.Consume(ctx, topics, handler)
	util.Logger.Info("kafka pull end", zap.Any("topics", topics), zap.Any("groupId", groupId), zap.Any("timeout", PullTimeout))
	if err != nil {
		util.Logger.Error("group consume error", zap.Error(err))
		return
	}

	for _, one := range handler.messages {
		var msg *Message
		msg = ConsumerMessageToMessage(keyType, valueType, one)
		msgList = append(msgList, msg)
	}
	return
}

type consumerGroupHandler struct {
	messages []*sarama.ConsumerMessage
	cancel   context.CancelFunc
	size     int
}

func (*consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (*consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (handler *consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	if sess == nil || claim == nil {
		return nil
	}
	chanMessages := claim.Messages()
	for msg := range chanMessages {
		handler.messages = append(handler.messages, msg)
		if len(handler.messages) >= handler.size {
			break
		}
	}
	if len(handler.messages) >= handler.size {
		if handler.cancel != nil {
			handler.cancel()
		}
	}
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

type Group struct {
	GroupId string `json:"groupId"`
	Cluster string `json:"cluster"`
}

func (this_ *Service) ListConsumerGroups() (res []*Group, err error) {
	var saramaClient sarama.Client
	saramaClient, err = this_.getClient()
	if err != nil {
		return
	}
	defer closeSaramaClient(saramaClient)
	manager, err := sarama.NewClusterAdminFromClient(saramaClient)
	if err != nil {
		return
	}
	defer closeClusterAdmin(manager)

	data, err := manager.ListConsumerGroups()
	if err != nil {
		return
	}
	if data != nil {
		for groupId, cluster := range data {
			res = append(res, &Group{
				GroupId: groupId,
				Cluster: cluster,
			})
		}
	}

	return
}

type GroupDescription struct {
	// Version defines the protocol version to use for encode and decode
	Version int16 `json:"version"`
	// Err contains the describe error as the KError type.
	Err sarama.KError `json:"err"`
	// ErrorCode contains the describe error, or 0 if there was no error.
	ErrorCode int16 `json:"errorCode"`
	// GroupId contains the group ID string.
	GroupId string `json:"groupId"`
	// State contains the group state string, or the empty string.
	State string `json:"state"`
	// ProtocolType contains the group protocol type, or the empty string.
	ProtocolType string `json:"protocolType"`
	// Protocol contains the group protocol data, or the empty string.
	Protocol string `json:"protocol"`
	// Members contains the group members.
	Members map[string]*GroupMemberDescription `json:"members"`
	// AuthorizedOperations contains a 32-bit bitfield to represent authorized
	// operations for this group.
	AuthorizedOperations int32 `json:"authorizedOperations"`
}

type GroupMemberDescription struct {
	// Version defines the protocol version to use for encode and decode
	Version int16 `json:"version"`
	// MemberId contains the member ID assigned by the group coordinator.
	MemberId string `json:"memberId"`
	// GroupInstanceId contains the unique identifier of the consumer instance
	// provided by end user.
	GroupInstanceId *string `json:"groupInstanceId"`
	// ClientId contains the client ID used in the member's latest join group
	// request.
	ClientId string `json:"clientId"`
	// ClientHost contains the client host.
	ClientHost string `json:"clientHost"`
	// MemberMetadata contains the metadata corresponding to the current group
	// protocol in use.
	MemberMetadata []byte `json:"memberMetadata"`
	// MemberAssignment contains the current assignment provided by the group
	// leader.
	MemberAssignment []byte `json:"memberAssignment"`
}

func (this_ *Service) DescribeConsumerGroups(groups []string) (res []*GroupDescription, err error) {
	var saramaClient sarama.Client
	saramaClient, err = this_.getClient()
	if err != nil {
		return
	}
	defer closeSaramaClient(saramaClient)
	manager, err := sarama.NewClusterAdminFromClient(saramaClient)
	if err != nil {
		return
	}
	defer closeClusterAdmin(manager)

	list, err := manager.DescribeConsumerGroups(groups)
	if err != nil {
		return
	}
	for _, one := range list {
		d := &GroupDescription{
			Version:              one.Version,
			GroupId:              one.GroupId,
			Err:                  one.Err,
			ErrorCode:            one.ErrorCode,
			State:                one.State,
			ProtocolType:         one.ProtocolType,
			Protocol:             one.Protocol,
			AuthorizedOperations: one.AuthorizedOperations,
			Members:              map[string]*GroupMemberDescription{},
		}
		if one.Members != nil {
			for key, v := range one.Members {
				d.Members[key] = &GroupMemberDescription{
					Version:          v.Version,
					MemberId:         v.MemberId,
					GroupInstanceId:  v.GroupInstanceId,
					ClientId:         v.ClientId,
					ClientHost:       v.ClientHost,
					MemberMetadata:   v.MemberMetadata,
					MemberAssignment: v.MemberAssignment,
				}
			}
		}
		res = append(res, d)
	}

	return
}

type LeaveGroupResponse struct {
	Version      int16            `json:"version"`
	ThrottleTime int32            `json:"throttleTime"`
	Err          sarama.KError    `json:"err"`
	Members      []MemberResponse `json:"members"`
}

type MemberResponse struct {
	MemberId        string        `json:"memberId"`
	GroupInstanceId *string       `json:"groupInstanceId"`
	Err             sarama.KError `json:"err"`
}

func (this_ *Service) RemoveMemberFromConsumerGroup(groupId string, groupInstanceIds []string) (res *LeaveGroupResponse, err error) {
	var saramaClient sarama.Client
	saramaClient, err = this_.getClient()
	if err != nil {
		return
	}
	defer closeSaramaClient(saramaClient)
	manager, err := sarama.NewClusterAdminFromClient(saramaClient)
	if err != nil {
		return
	}
	defer closeClusterAdmin(manager)

	one, err := manager.RemoveMemberFromConsumerGroup(groupId, groupInstanceIds)
	if err != nil {
		return
	}
	if one != nil {
		res = &LeaveGroupResponse{
			Version:      one.Version,
			ThrottleTime: one.ThrottleTime,
			Err:          one.Err,
		}
		for _, v := range one.Members {
			res.Members = append(res.Members, MemberResponse{
				MemberId:        v.MemberId,
				GroupInstanceId: v.GroupInstanceId,
				Err:             v.Err,
			})
		}
	}

	return
}

func (this_ *Service) DeleteConsumerGroupOffset(group string, topic string, partition int32) (err error) {
	var saramaClient sarama.Client
	saramaClient, err = this_.getClient()
	if err != nil {
		return
	}
	defer closeSaramaClient(saramaClient)
	manager, err := sarama.NewClusterAdminFromClient(saramaClient)
	if err != nil {
		return
	}
	defer closeClusterAdmin(manager)

	err = manager.DeleteConsumerGroupOffset(group, topic, partition)

	return
}

type OffsetFetchResponse struct {
	Version        int16                                          `json:"version"`
	ThrottleTimeMs int32                                          `json:"throttleTimeMs"`
	Blocks         map[string]map[int32]*OffsetFetchResponseBlock `json:"blocks"`
	Err            sarama.KError                                  `json:"err"`
}

type OffsetFetchResponseBlock struct {
	Offset      int64         `json:"offset"`
	LeaderEpoch int32         `json:"leaderEpoch"`
	Metadata    string        `json:"metadata"`
	Err         sarama.KError `json:"err"`
}

func (this_ *Service) ListConsumerGroupOffsets(group string, topicPartitions map[string][]int32) (res *OffsetFetchResponse, err error) {
	var saramaClient sarama.Client
	saramaClient, err = this_.getClient()
	if err != nil {
		return
	}
	defer closeSaramaClient(saramaClient)
	manager, err := sarama.NewClusterAdminFromClient(saramaClient)
	if err != nil {
		return
	}
	defer closeClusterAdmin(manager)

	one, err := manager.ListConsumerGroupOffsets(group, topicPartitions)
	if one != nil {
		res = &OffsetFetchResponse{
			Version:        one.Version,
			ThrottleTimeMs: one.ThrottleTimeMs,
			Err:            one.Err,
			Blocks:         map[string]map[int32]*OffsetFetchResponseBlock{},
		}
		for k, v := range one.Blocks {
			s := map[int32]*OffsetFetchResponseBlock{}
			if v != nil {
				for k_, v_ := range v {
					s[k_] = &OffsetFetchResponseBlock{
						Offset:      v_.Offset,
						LeaderEpoch: v_.LeaderEpoch,
						Metadata:    v_.Metadata,
						Err:         v_.Err,
					}
				}
			}
			res.Blocks[k] = s
		}
	}

	return
}

func (this_ *Service) GetClient() (res sarama.Client, err error) {
	res, err = this_.getClient()
	if err != nil {
		return
	}

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

type TopicMetadata struct {
	// Version defines the protocol version to use for encode and decode
	Version int16 `json:"version"`
	// Err contains the topic error, or 0 if there was no error.
	Err sarama.KError `json:"err"`
	// Name contains the topic name.
	Name string `json:"name"`
	// IsInternal contains a True if the topic is internal.
	IsInternal bool `json:"isInternal"`
	// Partitions contains each partition in the topic.
	Partitions []*PartitionMetadata `json:"partitions"`
}

type PartitionMetadata struct {
	// Version defines the protocol version to use for encode and decode
	Version int16 `json:"version"`
	// Err contains the partition error, or 0 if there was no error.
	Err sarama.KError `json:"err"`
	// ID contains the partition index.
	ID int32 `json:"ID"`
	// Leader contains the ID of the leader broker.
	Leader int32 `json:"leader"`
	// LeaderEpoch contains the leader epoch of this partition.
	LeaderEpoch int32 `json:"leaderEpoch"`
	// Replicas contains the set of all nodes that host this partition.
	Replicas []int32 `json:"replicas"`
	// Isr contains the set of nodes that are in sync with the leader for this partition.
	Isr []int32 `json:"isr"`
	// OfflineReplicas contains the set of offline replicas of this partition.
	OfflineReplicas []int32 `json:"offlineReplicas"`
}

func (this_ *Service) DescribeTopics(topics []string) (res []*TopicMetadata, err error) {
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

	list, err := admin.DescribeTopics(topics)
	if err != nil {
		return
	}

	for _, one := range list {
		d := &TopicMetadata{
			Version:    one.Version,
			Err:        one.Err,
			Name:       one.Name,
			IsInternal: one.IsInternal,
		}
		for _, v := range one.Partitions {
			d.Partitions = append(d.Partitions, &PartitionMetadata{
				Version:         v.Version,
				Err:             v.Err,
				ID:              v.ID,
				Leader:          v.Leader,
				LeaderEpoch:     v.LeaderEpoch,
				Replicas:        v.Replicas,
				Isr:             v.Isr,
				OfflineReplicas: v.OfflineReplicas,
			})
		}
		res = append(res, d)
	}

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

func ConsumerMessageToMessage(keyType string, valueType string, consumerMessage *sarama.ConsumerMessage) (msg *Message) {
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
