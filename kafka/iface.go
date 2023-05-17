package kafka

import "github.com/Shopify/sarama"

type IService interface {
	// Close 关闭 kafka 客户端
	Close()
	// Info 查看 kafka 信息
	Info() (res *Info, err error)
	// GetTopics 获取主题
	GetTopics() (res []*TopicInfo, err error)
	// GetTopic 获取主题
	GetTopic(topic string, time int64) (res *TopicInfo, err error)
	// Pull 拉取消息
	Pull(groupId string, topics []string, PullSize int, PullTimeout int, keyType, valueType string) (msgList []*Message, err error)
	// MarkOffset 提交 位置
	MarkOffset(groupId string, topic string, partition int32, offset int64) (err error)
	// ResetOffset 重置 位置
	ResetOffset(groupId string, topic string, partition int32, offset int64) (err error)
	// CreatePartitions 创建 主题 分区
	CreatePartitions(topic string, count int32) (err error)
	// CreateTopic 创建主题
	CreateTopic(topic string, numPartitions int32, replicationFactor int16) (err error)
	// DeleteTopic 删除 主题
	DeleteTopic(topic string) (err error)
	// DeleteConsumerGroup 删除 某个 消费组
	DeleteConsumerGroup(groupId string) (err error)
	// DeleteRecords 删除 主题 数据
	DeleteRecords(topic string, partitionOffsets map[int32]int64) (err error)
	// NewSyncProducer 创建 提供者
	NewSyncProducer() (syncProducer sarama.SyncProducer, err error)
	// Push 推送
	Push(msg *Message) (err error)
	// GetOffset 获取 主题 某个 分区 最新 位置
	GetOffset(topic string, partitionID int32, time int64) (offset int64, err error)
	// Partitions 获取 主题 分区
	Partitions(topic string) (partitions []int32, err error)
	// ListConsumerGroups 查询 所有 消费组
	ListConsumerGroups() (res []*Group, err error)
	// DescribeConsumerGroups 查询 消费组 明细
	DescribeConsumerGroups(groups []string) (res []*GroupDescription, err error)
	// DeleteConsumerGroupOffset 删除 消费组 某个主题 分区
	DeleteConsumerGroupOffset(group string, topic string, partition int32) (err error)
	// ListConsumerGroupOffsets 查询 消费组 主题分区 信息
	ListConsumerGroupOffsets(group string, topicPartitions map[string][]int32) (res *OffsetFetchResponse, err error)
	// RemoveMemberFromConsumerGroup 删除 消费组 成员
	RemoveMemberFromConsumerGroup(groupId string, groupInstanceIds []string) (res *LeaveGroupResponse, err error)
	// DescribeTopics 主题 元数据
	DescribeTopics(topics []string) (res []*TopicMetadata, err error)
	// GetClient 获取 kafka 客户端
	GetClient() (res sarama.Client, err error)
}
