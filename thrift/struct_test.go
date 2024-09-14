package thrift

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/team-ide/go-tool/util"
	"testing"
)

func TestStruct(t *testing.T) {
	s := &Struct{
		Fields: []*Field{
			{Name: "name", Num: 1, Type: GetFieldType(thrift.STRING)},
			{Name: "userId", Num: 2, Type: GetFieldType(thrift.I64)},
		},
	}
	s.SetData(map[string]any{
		"name":    "张三",
		"userId":  1000,
		"userId1": 1000,
	})

	bs, err := Serialize(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(bs))
	fmt.Println("Serialize data:", util.GetStringValue(s.GetData()))
	s = &Struct{
		Fields: []*Field{
			{Name: "name", Num: 1, Type: GetFieldType(thrift.STRING)},
			{Name: "userId", Num: 2, Type: GetFieldType(thrift.I64)},
		},
	}
	err = Deserialize(s, bs)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deserialize data:", util.GetStringValue(s.GetData()))

}
