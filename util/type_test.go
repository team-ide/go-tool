package util

import (
	"fmt"
	"reflect"
	"testing"
)

func TestType(t *testing.T) {
	testType(1, 111)
	testType(new(int), 111)

	testType(1, nil)
	testType(new(int), nil)

	testType(int64(1), 111)
	testType(new(int64), 111)

	testType(int64(1), nil)
	testType(new(int64), nil)

	testType(float64(1), 111)
	testType(new(float64), 111)

	testType(float64(1), nil)
	testType(new(float64), nil)

	testType("", 111)
	testType(new(string), 111)

	testType("", nil)
	testType(new(string), nil)

	testType(TestT{}, map[string]interface{}{
		"name": "这是名字",
	})
	testType(new(TestT), map[string]interface{}{
		"name": "这是名字",
	})

	testType(TestT{}, nil)
	testType(new(TestT), nil)

}

type TestT struct {
	Name string `json:"name"`
}

func testType(tV interface{}, data interface{}) {
	fmt.Println("------tV:", tV, ",data:", data, " start------")
	tVType := reflect.TypeOf(tV)
	v, err := GetValueByType(tVType, data)
	if err != nil {
		panic(err)
	}
	fmt.Println("value:", v, ",type:", reflect.TypeOf(v))
	if tVType.Kind() == reflect.Ptr && v != nil {
		fmt.Println("value Elem:", reflect.ValueOf(v).Elem().Interface())
	}
}
