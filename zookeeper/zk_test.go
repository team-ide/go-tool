package zookeeper

import (
	"fmt"
	"testing"
)

func TestZK(t *testing.T) {
	zk, err := New(&Config{
		Address: "192.168.0.85:11100",
	})
	if err != nil {
		panic(err)
	}
	Exists, err := zk.Exists("/xx")
	if err != nil {
		panic(err)
	}
	fmt.Println("Exists:", Exists)

}
