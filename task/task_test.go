package task

import (
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestTask(t *testing.T) {
	task, err := New(&Options{
		Key:       "xx",
		Worker:    1,
		Frequency: 1,
		Executor:  &testExecutor{},
	})
	if err != nil {
		panic(err)
	}
	util.Logger.Info("task run start", zap.Any("task", task))
	task.Run()
	util.Logger.Info("task run end", zap.Any("task", task))
}

type testExecutor struct {
}

func (this_ *testExecutor) Before(param *ExecutorParam) (err error) {
	util.Logger.Info("test Before", zap.Any("param", param))
	return
}

func (this_ *testExecutor) Execute(param *ExecutorParam) (err error) {
	util.Logger.Info("test Execute", zap.Any("param", param))
	time.Sleep(5 * time.Second)
	return
}

func (this_ *testExecutor) After(param *ExecutorParam) (err error) {
	util.Logger.Info("test After", zap.Any("param", param))
	return
}
