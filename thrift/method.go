package thrift

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

type MethodParam struct {
	Name       string        `json:"name,omitempty"`
	Args       []interface{} `json:"args,omitempty"`
	ArgFields  []*Field      `json:"argFields,omitempty"`
	Result     interface{}   `json:"result,omitempty"`
	ResultType *FieldType    `json:"resultFiled,omitempty"`
}

func (this_ *MethodParam) String() string {
	if this_ == nil {
		return "<nil>"
	}
	return fmt.Sprintf(this_.Name+"(%+v)", *this_)
}

func (this_ *MethodParam) Read(ctx context.Context, inProtocol thrift.TProtocol) error {

	fmt.Println("MethodParam Read ResultType:", toJSON(this_.ResultType))
	value, err := ReadStructFields(ctx, inProtocol, []*Field{
		{Num: 0, Name: "result", Type: this_.ResultType},
	})
	if err != nil {
		return err
	}

	this_.Result = value["result"]

	return nil
}

func (this_ *MethodParam) Write(ctx context.Context, outProtocol thrift.TProtocol) error {
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
