package util

import (
	"fmt"
	"reflect"
	"testing"
)

type TestS struct {
	Name       string `json:"name,omitempty" yaml:"name,omitempty" column:"xx" bson:"name"`
	TestSS     *TestSS
	TestSSList []*TestSS
	TestSSMap  map[string]*TestSS
}

type TestSS struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty" column:"xx" bson:"name"`
}

func TestStruct(t *testing.T) {
	cache := NewStructCache()
	fmt.Println(GetStringValue(cache.GetStructInfo(reflect.TypeOf(&TestS{}))))
	fmt.Println(GetStringValue(cache.GetStructInfo(reflect.TypeOf(TestS{}))))
}
