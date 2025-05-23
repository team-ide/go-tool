package util

import (
	"fmt"
	"testing"
)

func TestBase1(t *testing.T) {

	var aa any
	aa = new(int)

	var b = To[*int](aa)
	fmt.Println(*b)

	var ss = "123456"
	bb, err := StringTo[int64](ss)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bb)

}
