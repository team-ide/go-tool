package test

import (
	"fmt"
	"github.com/team-ide/cron"
	"github.com/team-ide/go-tool/util"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	cronHandler := cron.New(cron.WithSeconds())
	cronHandler.Start()

	_, err := cronHandler.AddFunc("0/5 * * * * ?", func() {
		fmt.Println("这是 cron 任务执行:", util.GetNowFormat())
	})
	if err != nil {
		panic(err)
	}
	//go func() {
	//	for {
	//		select {
	//		case <-time.After(time.Second * 5):
	//			fmt.Println("这是 time 任务执行:", util.GetNowFormat())
	//		}
	//	}
	//}()
	time.Sleep(time.Minute * 10)
}
