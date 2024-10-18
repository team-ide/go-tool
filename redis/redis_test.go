package redis

import (
	"fmt"
	"testing"
)

func TestRedis(t *testing.T) {

	s, err := NewRedisService(&Config{
		Address: "192.168.0.85:11080",
		Auth:    "q7ZtCl^5S3",
	})
	if err != nil {
		panic(err)
		return
	}
	res, err := s.Scan("*online*")
	if err != nil {
		panic(err)
		return
	}
	for _, key := range res.KeyList {
		fmt.Println(key)
	}
}
