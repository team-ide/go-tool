package task

import (
	"encoding/json"
	"fmt"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestTask(t *testing.T) {
	task, err := New(&Options{
		Key:       "xx",
		Worker:    50,
		Frequency: 1000,
		Executor:  &testExecutor{},
	})
	if err != nil {
		panic(err)
	}
	util.Logger.Info("task run start", zap.Any("task", task))
	task.Run()
	util.Logger.Info("task run end", zap.Any("task", task))

	fmt.Println("-----总统计------")
	bs, _ := json.Marshal(task.Metric.GetCount())
	fmt.Println(string(bs))

	fmt.Println("-----秒统计 开始------")
	cs := task.Metric.GetSecondCounts()
	for _, c := range cs {
		fmt.Println("秒时间：", util.TimeFormat(time.UnixMilli(c.StartTime/int64(time.Millisecond)), "2006-01-02 15:04:05"))
		bs, _ := json.Marshal(c)
		fmt.Println(string(bs))
	}
}

type testExecutor struct {
}

func (this_ *testExecutor) Before(param *ExecutorParam) (err error) {
	num := util.RandomInt(10, 50)
	time.Sleep(time.Millisecond * time.Duration(num))
	param.Extend = map[string]interface{}{
		"beforeNum": num,
	}
	//util.Logger.Info("test Before", zap.Any("param", param))
	return
}

func (this_ *testExecutor) Execute(param *ExecutorParam) (err error) {
	//util.Logger.Info("test Execute", zap.Any("param", param))
	num := util.RandomInt(50, 200)
	time.Sleep(time.Millisecond * time.Duration(num))
	param.Extend.(map[string]interface{})["executeNum"] = num
	return
}

func (this_ *testExecutor) After(param *ExecutorParam) (err error) {
	//util.Logger.Info("test After", zap.Any("param", param))
	return
}
