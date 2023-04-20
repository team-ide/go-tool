package thrift

import (
	"context"
	"errors"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

type Struct struct {
	Include string  `json:"include"`
	Name    string  `json:"name"`
	Fields  []Field `json:"fields"`
	Value   map[string]interface{}
}

func (this_ *Struct) String() string {
	if this_ == nil {
		return "<nil>"
	}
	return fmt.Sprintf(this_.Name+"(%+v)", *this_)
}

func (this_ *Struct) Read(ctx context.Context, protocol thrift.TProtocol) error {
	value, err := ReadStruct(ctx, protocol, this_)
	if err != nil {
		return err
	}

	this_.Value = value

	return nil
}

type Field struct {
	Num   int16       `json:"num"`
	Name  string      `json:"name"`
	Type  FieldType   `json:"type"`
	Value interface{} `json:"value"`
}

type FieldType struct {
	Include      string      `json:"include"`
	Type         string      `json:"type"`
	GenericTypes []FieldType `json:"genericTypes"`
}

func (this_ *Field) Read(ctx context.Context, protocol thrift.TProtocol, fieldTypeId thrift.TType, value map[string]interface{}) error {
	// 判断类型是否一致
	if fieldTypeId != thrift.I16 {
		if err := protocol.Skip(ctx, fieldTypeId); err != nil {
			return err
		}
	}
	if fieldTypeId == thrift.LIST {

	} else if fieldTypeId == thrift.STRUCT {

	} else if fieldTypeId == thrift.MAP {

	}
	if v, err := protocol.ReadByte(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("error reading field %d:%s : ", this_.Num, this_.Name), err)
	} else {
		value[this_.Name] = v
	}
	return nil
}

func ReadStruct(ctx context.Context, protocol thrift.TProtocol, Struct *Struct) (map[string]interface{}, error) {
	if _, err := protocol.ReadStructBegin(ctx); err != nil {
		return nil, thrift.PrependError(fmt.Sprintf("%T read error: ", Struct), err)
	}

	fieldMap := map[int16]Field{}
	for _, one := range Struct.Fields {
		fieldMap[one.Num] = one
	}
	value := make(map[string]interface{})

	for {
		_, fieldTypeId, fieldId, err := protocol.ReadFieldBegin(ctx)
		if err != nil {
			return nil, thrift.PrependError(fmt.Sprintf("%T field %d read error: ", Struct, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		field, ok := fieldMap[fieldId]
		if !ok {
			if err := protocol.Skip(ctx, fieldTypeId); err != nil {
				return nil, err
			}
		}
		if err := field.Read(ctx, protocol, fieldTypeId, value); err != nil {
			return nil, err
		}

		if err := protocol.ReadFieldEnd(ctx); err != nil {
			return nil, err
		}
	}
	if err := protocol.ReadStructEnd(ctx); err != nil {
		return nil, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", Struct), err)
	}
	return value, nil
}

func WriteStruct(ctx context.Context, protocol thrift.TProtocol, Struct *Struct) error {
	if err := protocol.WriteStructBegin(ctx, "Request"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", Struct), err)
	}
	if Struct != nil {
		for _, filed := range Struct.Fields {
			var v interface{}
			if err := filed.Write(ctx, protocol, v); err != nil {
				return err
			}
		}
	}
	if err := protocol.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := protocol.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (this_ *Field) Write(ctx context.Context, protocol thrift.TProtocol, value interface{}) error {
	if err := protocol.WriteFieldBegin(ctx, this_.Name, thrift.BYTE, this_.Num); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:field1: ", this_), err)
	}
	if this_.Type.Type == "int8" {
		if err := protocol.WriteByte(ctx, value.(int8)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.field1 (1) field write error: ", this_), err)
		}
	} else {
		return thrift.PrependError(fmt.Sprintf("%T.field1 (1) field type error: ", this_), errors.New("type unknown"))
	}
	if err := protocol.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:field1: ", this_), err)
	}
	return nil
}
