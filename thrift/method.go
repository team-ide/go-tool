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
	ReadStart       int64         `json:"readStart"`
	ReadEnd         int64         `json:"readEnd"`
	WriteStart      int64         `json:"writeStart"`
	WriteEnd        int64         `json:"writeEnd"`
	UseTime         int64         `json:"useTime"`
	Error           error         `json:"error"`
}

func (this_ *MethodParam) String() string {
	if this_ == nil {
		return "<nil>"
	}
	return fmt.Sprintf(this_.Name+"(%+v)", *this_)
}

func (this_ *MethodParam) Read(ctx context.Context, inProtocol thrift.TProtocol) error {
	this_.ReadStart = time.Now().UnixMilli()
	if this_.WriteEnd > 0 {
		this_.UseTime = this_.ReadStart - this_.WriteEnd
	}
	defer func() {
		this_.ReadEnd = time.Now().UnixMilli()
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
	this_.WriteStart = time.Now().UnixMilli()
	defer func() {
		this_.WriteEnd = time.Now().UnixMilli()
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
