package metric

import (
	"fmt"
	"github.com/team-ide/go-tool/util"
	"sync"
	"testing"
	"time"
)

func TestMetric(t *testing.T) {
	metric := NewMetric()

	wait := sync.WaitGroup{}

	worker := 1000
	wait.Add(worker)
	var isStop bool

	for i := 0; i < worker; i++ {
		go func(workerIndex int) {
			workerMetric := metric.NewWorkerMetric(workerIndex)
			for !isStop {
				//s := time.Now().UnixMilli()
				item := workerMetric.NewItem(time.Now().UnixNano())

				Wait(time.Microsecond * 1)
				// 随机等待一段时间 模拟其它准备耗时
				startTime := time.Now().UnixNano()
				//fmt.Println("worker index:", workerIndex, ",data index:", dataIndex)
				// 随机等等一段时间 模拟执行耗时
				num := util.RandomInt(1, 5)
				Wait(time.Microsecond * time.Duration(num))
				var err error = nil
				useTime := int(time.Now().UnixNano() - startTime)
				if util.RandomInt(1, 50) == 5 {
					useTime = 0
				}
				item.End(useTime, time.Now().UnixNano(), err)
				//e := time.Now().UnixMilli()
				//fmt.Println("num：", num, " 执行耗时：", useTime/1000000, "ms 总耗时：", e-s, "ms")
			}

			wait.Done()

		}(i)
	}
	go func() {
		time.Sleep(time.Minute * 10)
		isStop = true
	}()

	metric.SetOnCount(func() {
		outMetric(metric)
	})
	metric.StartCount()
	wait.Wait()
	metric.StopCount()

}

func outMetric(metric *Metric) {
	var text string
	count := metric.GetCount()
	text = MarkdownTable([]*Count{count}, nil)
	fmt.Println("-----总统计 信息------")
	fmt.Println(text)

	fmt.Println("-----秒统计 信息------")
	cs := metric.GetSecondCounts()
	text = MarkdownTable(cs, nil)
	fmt.Println(text)
}

func Wait(d time.Duration) {
	t1 := time.NewTicker(d)
	<-t1.C
	t1.Stop()
}
