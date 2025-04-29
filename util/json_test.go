package util

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type TestObj struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Price float32 `json:"price"`
	Other string  `json:"other"`
}

func TestJson(t *testing.T) {

	obj := TestObj{}
	obj.Age = 12
	obj.Price = 1.23
	obj.Name = "张三"
	obj.Other = "aaa<d>dd<>asd?asd&ddd"

	s, err := ObjToJson(obj)
	if err != nil {
		panic(err)
	}
	fmt.Println("obj to json:", s)

	data := map[string]any{}
	err = ObjToObjByJson(obj, &data)
	if err != nil {
		panic(err)
	}
	fmt.Println("obj to map:", data)
	for k, v := range data {
		fmt.Println("key:", k, ",value:", v, ",type:", reflect.TypeOf(v))
	}
	data = map[string]any{}
	err = json.Unmarshal([]byte(s), &data)
	if err != nil {
		panic(err)
	}
	fmt.Println("obj to map use json:", data)
	for k, v := range data {
		fmt.Println("key:", k, ",value:", v, ",type:", reflect.TypeOf(v))
	}
}
