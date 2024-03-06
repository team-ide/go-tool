package redis

import (
	"fmt"
	"strconv"
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
	ss := string([]byte{172, 237, 0, 5, 116, 0, 35, 103, 114, 111, 117, 112, 58, 109, 101, 109, 98, 101, 114, 115, 58, 52, 54, 49, 49, 54, 56, 54, 48, 50, 57, 49, 54, 52, 56, 48, 54, 51, 56, 55, 58, 48})
	//s.Set("\\xac\\xed\\x00\\x05t\\x00#group:members:4611686029164806387:0", "xx")
	s.Set(ss, "xx")
	res, err := s.Keys("*group:members:4611686029164806387*")
	if err != nil {
		panic(err)
		return
	}
	for _, key := range res.KeyList {
		//k := strconv.QuoteToASCII(key.Key)
		//k = k[1 : len(k)-1]
		//s.Del(key.Key)
		fmt.Println(strconv.QuoteToASCII(key.Key))
		fmt.Println(strconv.Quote(key.Key))
		//s.Del(k)
		return
	}
}
