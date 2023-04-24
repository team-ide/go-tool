package task

import (
	"errors"
	"fmt"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"sync"
	"time"
)

func New(options *Options) (task *Task, err error) {
	if options == nil {
		err = errors.New("options 配置项不能为空")
		return
	}

	if options.Key == "" {
		err = errors.New("任务 Key 不能为空")
		return
	}
	if options.Executor == nil {
		err = errors.New("任务 执行器 不能为空")
		return
	}
	if options.Worker <= 0 {
		err = errors.New("任务 工作线程 必须大于 0")
		return
	}
	if options.Frequency <= 0 && options.Duration <= 0 {
		err = errors.New("任务 必须配置 执行时长 或 执行次数")
		return
	}
	if options.Duration > 0 {
		options.durationMilli = int64(options.Duration * 60 * 1000)
	}
	task = &Task{
		Options:       options,
		nextLocker:    &sync.Mutex{},
		counterLocker: &sync.Mutex{},
		waitGroup:     &sync.WaitGroup{},
	}

	return
}

type Task struct {
	*Options

	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Errors    []error   `json:"errors"`

	OnStart func() `json:"-"` // DoStart 执行开始
	OnEnd   func() `json:"-"` // DoEnd 执行结束
	OnStop  func() `json:"-"`

	isStop  bool // isStop 是否需要停止
	IsStart bool `json:"isStart"` // IsStart 是否启动
	IsEnd   bool `json:"isEnd"`   // IsEnd 是否结束

	getNextCount int // 调用 getNext 次数

	nextLocker    sync.Locker
	counterLocker sync.Locker

	ExecutorBeforeCount  int `json:"executorBeforeCount"`
	ExecutorExecuteCount int `json:"executorExecuteCount"`
	ExecutorAfterCount   int `json:"executorAfterCount"`
	ExecutorSuccessCount int `json:"executorSuccessCount"`
	ExecutorErrorCount   int `json:"executorErrorCount"`

	waitGroup *sync.WaitGroup
}

func (this_ *Task) Run() {
	if this_.IsStart {
		return
	}

	util.Logger.Info("任务执行 [Start]", zap.Any("Key", this_.Key))
	this_.IsStart = true
	this_.IsEnd = false
	this_.StartTime = time.Now()

	defer func() {
		this_.runAfter()
		this_.EndTime = time.Now()
		this_.IsEnd = true
	}()

	if !this_.runBefore() {
		return
	}

	this_.runDo()

}

func (this_ *Task) runBefore() bool {

	defer func() {
		if e := recover(); e != nil {
			err := errors.New(fmt.Sprintf("任务执行 [runBefore] 异常:%s", e))
			this_.Errors = append(this_.Errors, err)
			util.Logger.Error("runBefore error", zap.Error(err))
		}
	}()
	if this_.IsStopped() {
		return false
	}
	util.Logger.Info("任务执行 [runBefore]", zap.Any("Key", this_.Key))

	if this_.OnStart != nil {
		this_.OnStart()
	}

	return true
}

func (this_ *Task) runDo() {

	defer func() {
		if e := recover(); e != nil {
			err := errors.New(fmt.Sprintf("任务执行 [runDo] 异常:%s", e))
			this_.Errors = append(this_.Errors, err)
			util.Logger.Error("runDo error", zap.Error(err))
		}
	}()
	util.Logger.Info("任务执行 [runDo]", zap.Any("Key", this_.Key))

	this_.waitGroup.Add(this_.Worker)
	for i := 0; i < this_.Worker; i++ {
		go func(workerIndex int) {
			defer func() {
				this_.waitGroup.Done()
			}()
			this_.workerWork(workerIndex)
		}(i)
	}

	this_.waitGroup.Wait()

}

func (this_ *Task) workerWork(workerIndex int) {
	index := this_.getNext()
	// 索引 小于0 表示结束
	if index < 0 {
		return
	}

	param := &ExecutorParam{
		Index:       index,
		WorkerIndex: workerIndex,
	}
	this_.runExecutor(param)
}

func (this_ *Task) executorDo(param *ExecutorParam, counter *int, start *time.Time, end *time.Time, do func(param *ExecutorParam) error) (err error) {
	defer func() {
		*end = time.Now()
		this_.counterLocker.Lock()
		defer this_.counterLocker.Unlock()

		*counter++
	}()
	*start = time.Now()
	err = do(param)
	return
}

func (this_ *Task) runExecutor(param *ExecutorParam) {
	var err error
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("任务执行 [runExecutor] 异常:%s", e))
			util.Logger.Error("runExecutor error", zap.Error(err))
		}

		param.EndTime = time.Now()
		param.Error = err

		this_.counterLocker.Lock()
		defer this_.counterLocker.Unlock()

		if err != nil {
			this_.ExecutorErrorCount++
		} else {
			this_.ExecutorSuccessCount++
		}
	}()

	param.StartTime = time.Now()

	err = this_.executorDo(param, &this_.ExecutorBeforeCount, &param.BeforeStartTime, &param.BeforeEndTime, this_.Executor.Before)

	if err != nil {
		return
	}

	err = this_.executorDo(param, &this_.ExecutorExecuteCount, &param.ExecuteStartTime, &param.ExecuteEndTime, this_.Executor.Execute)

	if err != nil {
		return
	}

	err = this_.executorDo(param, &this_.ExecutorAfterCount, &param.AfterStartTime, &param.AfterEndTime, this_.Executor.After)

	if err != nil {
		return
	}
}

func (this_ *Task) getNext() (index int) {
	this_.nextLocker.Lock()
	defer this_.nextLocker.Unlock()

	index = -1 // 返回 -1 表示结束

	// 如果已经停止
	if this_.IsStopped() {
		return
	}
	// 如果 设置 执行次数 则 判断
	if this_.Frequency > 0 {
		// getNext 计数 大于等于 总次数 则 不在执行
		if this_.getNextCount >= this_.Frequency {
			return
		}

	} else if this_.Duration > 0 { // 如果设置 执行时长 则判断
		nowMilli := time.Now().UnixMilli()
		startMilli := this_.StartTime.UnixMilli()

		// 执行 毫秒 大于等于 总时长 则 不在执行
		if nowMilli-startMilli > this_.durationMilli {
			return
		}

	}
	this_.getNextCount++

	index = this_.getNextCount - 1

	return
}

func (this_ *Task) runAfter() {

	defer func() {
		if e := recover(); e != nil {
			err := errors.New(fmt.Sprintf("任务执行 [runAfter] 异常:%s", e))
			this_.Errors = append(this_.Errors, err)
			util.Logger.Error("runAfter error", zap.Error(err))
		}
	}()
	util.Logger.Info("任务执行 [runAfter]", zap.Any("Key", this_.Key))

	if this_.OnEnd != nil {
		this_.OnEnd()
	}

}

func (this_ *Task) Stop() {
	defer func() {
		if e := recover(); e != nil {
			err := errors.New(fmt.Sprintf("任务执行 [Stop] 异常:%s", e))
			this_.Errors = append(this_.Errors, err)
			util.Logger.Error("Stop error", zap.Error(err))
		}
	}()
	this_.isStop = true
	util.Logger.Info("任务执行 [Stop]", zap.Any("Key", this_.Key))
	if this_.OnStop != nil {
		this_.OnStop()
	}

}

func (this_ *Task) IsStopped() bool {

	return this_.isStop

}
