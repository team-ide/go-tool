package task

import (
	"time"
)

type Executor interface {
	Before(param *ExecutorParam) (err error)  // 每次执行之前调用
	Execute(param *ExecutorParam) (err error) // 执行
	After(param *ExecutorParam) (err error)   // 每次执行之后调用
}

type ExecutorParam struct {
	Index       int         `json:"index"`       // 索引号 从 0 开始
	WorkerIndex int         `json:"workerIndex"` // 执行线程编号
	Extend      interface{} `json:"extend"`      // 扩展 业务自己数据可以放在此处

	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`

	BeforeStartTime time.Time `json:"beforeStartTime"`
	BeforeEndTime   time.Time `json:"beforeEndTime"`

	ExecuteStartTime time.Time `json:"executeStartTime"`
	ExecuteEndTime   time.Time `json:"executeEndTime"`

	AfterStartTime time.Time `json:"afterStartTime"`
	AfterEndTime   time.Time `json:"afterEndTime"`

	Error error `json:"error"`

	isStop bool // isStop 是否需要停止
}

func (this_ *ExecutorParam) Stop() {
	this_.isStop = true
}

func (this_ *ExecutorParam) IsStopped() bool {

	return this_.isStop

}
