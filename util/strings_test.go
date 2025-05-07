package util

import (
	"fmt"
	"testing"
)

func TestGetString(t *testing.T) {
	fmt.Println(GetStringValue("111"))
	var s = new(string)
	*s = "111"
	fmt.Println(GetStringValue(s))
	var i = new(int)
	*i = 111
	fmt.Println(GetStringValue(i))
}
