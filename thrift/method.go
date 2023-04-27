package thrift

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"time"
)

type MethodParam struct {
	Name            string        `json:"name"`
	Args            []interface{} `json:"args"`
	ArgFields       []*Field      `json:"argFields,omitempty"`
	Result          interface{}   `json:"result"`
	Exceptions      []interface{} `json:"exceptions"`
	ResultType      *FieldType    `json:"resultFiled,omitempty"`
	ExceptionFields []*Field      `json:"exceptionFields,omitempty"`
	ReadStart       time.Time     `json:"readStart"`
	ReadEnd         time.Time     `json:"readEnd"`
	WriteStart      time.Time     `json:"writeStart"`
	WriteEnd        time.Time     `json:"writeEnd"`
	Error           error         `json:"error"`
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
	var fields []*Field
	if this_.ResultType != nil {
		fields = append(fields, &Field{
			Num: 0, Name: "success", Type: this_.ResultType,
		})
	}
	for _, e := range this_.ExceptionFields {
		fields = append(fields, e)
	}
	//fmt.Println("MethodParam Read ResultType:", toJSON(this_.ResultType))
	resultMap, err := ReadStructFields(ctx, inProtocol, fields)
	if resultMap != nil {
		this_.Result = resultMap["success"]
		for _, e := range this_.ExceptionFields {
			if resultMap[e.Name] != nil {
				this_.Exceptions = append(this_.Exceptions, resultMap[e.Name])
			}
		}
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
