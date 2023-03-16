package webproxy

import (
	"sync"
	"testing"
)

func TestServer(t *testing.T) {
	Start()
	wait := sync.WaitGroup{}
	wait.Add(1)
	go func() {
		//time.Sleep(time.Second * 5)
		//wait.Done()
	}()
	wait.Wait()
}
