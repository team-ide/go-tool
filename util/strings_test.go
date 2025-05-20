package util

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetString(t *testing.T) {
	var s = new(string)
	*s = "111"
	bs, _ := json.Marshal(*s)
	fmt.Println(string(bs))
	fmt.Println(fmt.Sprintf("%v", s))
	fmt.Println(*s)
	fmt.Println(GetStringValue("111"))
	fmt.Println(GetStringValue(s))
	var i = new(int)
	*i = 111
	fmt.Println(GetStringValue(i))
}
