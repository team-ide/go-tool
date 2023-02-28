package zookeeper

import (
	"fmt"
	"sync"
	"testing"
)

func TestWatchChildren(t *testing.T) {
	zkService, err := New(Config{Address: "127.0.0.1:2181"})
	if err != nil {
		panic(err)
	}

	wait := &sync.WaitGroup{}
	wait.Add(1)
	err = zkService.WatchChildren("/test/node1",
		func(data *WatchChildrenListenData) (finish bool) {
			defer func() {
				if finish {
					wait.Done()
				}
			}()
			if data.Event == NodeEventNodeNotFound {
				finish = true
			}
			fmt.Println("listen event:", data.Event, ",child:", data.Child, ",err:", data.Err)
			return
		},
	)
	if err != nil {
		panic(err)
	}

	wait.Wait()
}
