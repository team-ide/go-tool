package kafka

import "github.com/Shopify/sarama"

type IService interface {
	Stop()
	Info() (res *Info, err error)
	GetTopics() (res []*TopicInfo, err error)
	GetTopic(topic string, time int64) (res *TopicInfo, err error)
	Pull(groupId string, topics []string, PullSize int, PullTimeout int, keyType, valueType string) (msgList []*Message, err error)
	MarkOffset(groupId string, topic string, partition int32, offset int64) (err error)
	ResetOffset(groupId string, topic string, partition int32, offset int64) (err error)
	CreatePartitions(topic string, count int32) (err error)
	CreateTopic(topic string, numPartitions int32, replicationFactor int16) (err error)
	DeleteTopic(topic string) (err error)
	DeleteConsumerGroup(groupId string) (err error)
	DeleteRecords(topic string, partitionOffsets map[int32]int64) (err error)
	NewSyncProducer() (syncProducer sarama.SyncProducer, err error)
	Push(msg *Message) (err error)
	GetOffset(topic string, partitionID int32, time int64) (offset int64, err error)
	Partitions(topic string) (partitions []int32, err error)
	ListConsumerGroups() (res []*Group, err error)
	DescribeConsumerGroups(groups []string) (res []*GroupDescription, err error)
	DeleteConsumerGroupOffset(group string, topic string, partition int32) (err error)
	ListConsumerGroupOffsets(group string, topicPartitions map[string][]int32) (res *OffsetFetchResponse, err error)
	RemoveMemberFromConsumerGroup(groupId string, groupInstanceIds []string) (res *LeaveGroupResponse, err error)
	GetClient() (res sarama.Client, err error)
}
