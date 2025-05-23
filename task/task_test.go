package task

import (
	"fmt"
	"github.com/team-ide/go-tool/metric"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestTask(t *testing.T) {
	task, err := New(&Options{
		Key:      "xx",
		Worker:   1000,
		Duration: 10,
		Executor: &testExecutor{},
	})
	if err != nil {
		panic(err)
	}
	task.Metric.SetOnCount(func() {
		outMetric(task.Metric)
	})
	util.Logger.Info("task run start", zap.Any("task", task))
	task.Run()
	util.Logger.Info("task run end", zap.Any("task", task))

	fmt.Println("-----总统计------")
	s, _ := util.ObjToJson(task.Metric.GetCount())
	fmt.Println(s)

	fmt.Println("-----秒统计 开始------")
	cs := task.Metric.GetSecondCounts()
	for _, c := range cs {
		fmt.Println("秒时间：", util.TimeFormat(time.UnixMilli(c.StartTime/int64(time.Millisecond)), "2006-01-02 15:04:05"))
		s, _ = util.ObjToJson(c)
		fmt.Println(s)
	}
}

func outMetric(m *metric.Metric) {
	var text string
	count := m.GetCount()
	text = metric.MarkdownTable([]*metric.Count{count}, nil)
	fmt.Println("-----总统计 信息------")
	fmt.Println(text)

	fmt.Println("-----秒统计 信息------")
	cs := m.GetSecondCounts()
	text = metric.MarkdownTable(cs, nil)
	fmt.Println(text)
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
	num := util.RandomInt(1, 5)
	time.Sleep(time.Millisecond * time.Duration(num))
	param.Extend.(map[string]interface{})["executeNum"] = num
	return
}

func (this_ *testExecutor) After(param *ExecutorParam) (err error) {
	//util.Logger.Info("test After", zap.Any("param", param))
	return
}
