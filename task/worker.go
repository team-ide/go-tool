package task

import (
	"errors"
	"fmt"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"time"
)

type Worker struct {
	*Task
	WorkerIndex int `json:"workerIndex"` // 执行线程编号
}

func NewWorker(workerIndex int, task *Task) (worker *Worker) {
	worker = &Worker{
		Task:        task,
		WorkerIndex: workerIndex,
	}
	return
}

func (this_ *Worker) work() {
	defer func() {
		this_.waitGroup.Done()
	}()
	for {
		index := this_.getNext()
		// 索引 小于0 表示结束
		if index < 0 {
			break
		}
		param := &ExecutorParam{
			Index:       index,
			WorkerIndex: this_.WorkerIndex,
		}
		this_.runExecutor(param)
	}
}

func (this_ *Worker) executorDo(param *ExecutorParam, counter *int, start *time.Time, end *time.Time, do func(param *ExecutorParam) error) (err error) {
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

func (this_ *Worker) runExecutor(param *ExecutorParam) {
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

		if !param.ExecuteEndTime.IsZero() {
			allUse := param.EndTime.UnixNano() - param.StartTime.UnixNano()
			executeUse := param.ExecuteEndTime.UnixNano() - param.ExecuteStartTime.UnixNano()
			item := this_.Metric.NewItem(param.WorkerIndex, param.StartTime)
			item.Extend = param.Extend
			item.Loss(allUse - executeUse)
			item.End(param.EndTime, param.Error)
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
