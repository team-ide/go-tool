package thrift

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"time"
)

type MethodParam struct {
	Name       string                 `json:"name"`
	Args       []interface{}          `json:"args"`
	ArgFields  []*Field               `json:"argFields"`
	Result     interface{}            `json:"result"`
	ResultMap  map[string]interface{} `json:"resultMap"`
	ResultType *FieldType             `json:"resultFiled"`
	ReadStart  time.Time              `json:"readStart"`
	ReadEnd    time.Time              `json:"readEnd"`
	WriteStart time.Time              `json:"writeStart"`
	WriteEnd   time.Time              `json:"writeEnd"`
}

func (this_ *MethodParam) String() string {
	if this_ == nil {
		return "<nil>"
	}
	return fmt.Sprintf(this_.Name+"(%+v)", *this_)
}

func (this_ *MethodParam) Read(ctx context.Context, inProtocol thrift.TProtocol) error {
	this_.ReadStart = time.Now()
	defer func() {
		this_.ReadEnd = time.Now()
	}()
	var err error
	//fmt.Println("MethodParam Read ResultType:", toJSON(this_.ResultType))
	this_.ResultMap, err = ReadStructFields(ctx, inProtocol, []*Field{
		{Num: 0, Name: "result", Type: this_.ResultType},
	})
	if this_.ResultMap != nil {
		this_.Result = this_.ResultMap["result"]
	}
	if err != nil {
		return err
	}

	return nil
}

func (this_ *MethodParam) Write(ctx context.Context, outProtocol thrift.TProtocol) error {
	this_.WriteStart = time.Now()
	defer func() {
		this_.WriteEnd = time.Now()
	}()

	value := map[string]interface{}{}
	for index, field := range this_.ArgFields {
		value[field.Name] = this_.Args[index]
	}

	err := WriteStructFields(ctx, outProtocol, this_.Name+"_args", this_.ArgFields, value)
	if err != nil {
		return err
	}

	return nil
}
