package metric

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/team-ide/go-tool/util"
	"sync"
	"testing"
	"time"
)

func TestMetric(t *testing.T) {

	metric := NewMetric()

	wait := sync.WaitGroup{}

	worker := 10
	wait.Add(worker)

	for i := 0; i < worker; i++ {
		go func(workerIndex int) {

			for dataIndex := 0; dataIndex < 100; dataIndex++ {
				num := util.RandomInt(50, 200)
				loss := util.RandomInt(10, 20)

				item := metric.NewItem(workerIndex, time.Now())

				fmt.Println("worker index:", workerIndex, ",data index:", dataIndex)
				time.Sleep(time.Millisecond * time.Duration(num))
				var err error = nil
				if num%120 == 0 {
					err = errors.New("error execute")
				}
				item.Loss(1000000 * loss)
				item.End(time.Now(), err)
			}

			wait.Done()

		}(i)
	}

	metric.StartCount()

	wait.Wait()
	metric.StopCount()

	count := metric.GetCount()
	fmt.Println("-----总统计------")
	bs, _ := json.Marshal(count)
	fmt.Println(string(bs))

	fmt.Println("-----分钟统计 开始------")
	cs := metric.GetMinuteCounts()
	for _, c := range cs {
		fmt.Println("分钟时间：", util.TimeFormat(time.UnixMilli(c.StartTime/int64(time.Millisecond)), "2006-01-02 15:04"))
		bs, _ := json.Marshal(c)
		fmt.Println(string(bs))
	}
	fmt.Println("-----秒统计 开始------")
	cs = metric.GetSecondCounts()
	for _, c := range cs {
		fmt.Println("秒时间：", util.TimeFormat(time.UnixMilli(c.StartTime/int64(time.Millisecond)), "2006-01-02 15:04:05"))
		bs, _ := json.Marshal(c)
		fmt.Println(string(bs))
	}

}
