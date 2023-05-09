package redis

import (
	"fmt"
	"testing"
)

func TestRedis(t *testing.T) {
	config := &Config{
		Address: "127.0.0.1:6379",
	}
	service, err := New(config)
	if err != nil {
		panic("redis new error:" + err.Error())
	}

	v, notFound, err := service.Get(nil, "ss")
	if err != nil {
		panic("redis Get error:" + err.Error())
	}
	fmt.Println("redis Get value:", v, ",notFound:", notFound)

	v, notFound, err = service.HGet(nil, "ss", "ss")
	if err != nil {
		panic("redis HGet error:" + err.Error())
	}
	fmt.Println("redis HGet value:", v, ",notFound:", notFound)
}
