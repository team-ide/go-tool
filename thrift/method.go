package thrift

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

type MethodParam struct {
	Name        string        `json:"name"`
	Args        []interface{} `json:"args"`
	ArgFields   []*Field      `json:"argFields"`
	Result      interface{}   `json:"result"`
	ResultFiled *Field        `json:"resultFiled"`
}

func (this_ *MethodParam) String() string {
	if this_ == nil {
		return "<nil>"
	}
	return fmt.Sprintf(this_.Name+"(%+v)", *this_)
}

func (this_ *MethodParam) Read(ctx context.Context, inProtocol thrift.TProtocol) error {
	value, err := ReadStructFields(ctx, inProtocol, []*Field{this_.ResultFiled})
	if err != nil {
		return err
	}

	this_.Result = value

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
