package thrift

import (
	"context"
	"errors"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"sync"
)

var (
	structCache     map[string]*Struct
	structCacheLock sync.Locker = &sync.RWMutex{}
)

func AddStruct(str *Struct) {
	structCacheLock.Lock()
	defer structCacheLock.Unlock()

	key := str.Name
	if str.Include != "" {
		key = str.Include + "." + str.Name
	}
	structCache[key] = str
}

func GetStruct(name string) *Struct {
	structCacheLock.Lock()
	defer structCacheLock.Unlock()

	return structCache[name]
}

type Struct struct {
	Include string   `json:"include"`
	Name    string   `json:"name"`
	Fields  []*Field `json:"fields"`
}

type Field struct {
	Num   int16       `json:"num"`
	Name  string      `json:"name"`
	Type  *FieldType  `json:"type"`
	Value interface{} `json:"value"`
}

type FieldType struct {
	Include      string       `json:"include"`
	TypeId       thrift.TType `json:"typeId"`
	Struct       *Struct      `json:"struct"`
	GenericTypes []*FieldType `json:"genericTypes"`
}

func WriteStructFields(ctx context.Context, protocol thrift.TProtocol, name string, fields []*Field, value map[string]interface{}) error {

	fmt.Println("WriteStructFields name:", name, ",fields:", fields, ",value:", value)
	if err := protocol.WriteStructBegin(ctx, name); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", name), err)
	}
	for _, filed := range fields {
		var v = value[filed.Name]
		if err := WriteStructField(ctx, protocol, filed, v); err != nil {
			return err
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

func WriteStructField(ctx context.Context, protocol thrift.TProtocol, field *Field, value interface{}) error {
	fmt.Println("WriteStructField field:", field, ",value:", value)
	var err error
	if err = protocol.WriteFieldBegin(ctx, field.Name, field.Type.TypeId, field.Num); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error", field), err)
	}
	err = WriteByType(ctx, protocol, field.Type, value)
	if err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write error: ", field), err)
	}
	if err = protocol.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error", field), err)
	}
	return nil
}

func WriteMap(ctx context.Context, protocol thrift.TProtocol, keyType *FieldType, valueType *FieldType, value map[string]interface{}) error {
	if value == nil {
		value = map[string]interface{}{}
	}
	size := len(value)
	if err := protocol.WriteMapBegin(ctx, keyType.TypeId, valueType.TypeId, size); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write map begin error: ", keyType), err)
	}
	for key, v := range value {
		if err := WriteByType(ctx, protocol, keyType, key); err != nil {
			return err
		}
		if err := WriteByType(ctx, protocol, valueType, v); err != nil {
			return err
		}
	}
	if err := protocol.WriteMapEnd(ctx); err != nil {
		return thrift.PrependError("write map end error: ", err)
	}
	return nil
}

func WriteSet(ctx context.Context, protocol thrift.TProtocol, setType *FieldType, value []interface{}) error {
	size := len(value)
	if err := protocol.WriteSetBegin(ctx, setType.TypeId, size); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write set begin error: ", setType), err)
	}
	for _, v := range value {
		if err := WriteByType(ctx, protocol, setType, v); err != nil {
			return err
		}
	}
	if err := protocol.WriteMapEnd(ctx); err != nil {
		return thrift.PrependError("write set end error: ", err)
	}
	return nil
}

func WriteList(ctx context.Context, protocol thrift.TProtocol, listType *FieldType, value []interface{}) error {
	size := len(value)
	if err := protocol.WriteListBegin(ctx, listType.TypeId, size); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write list begin error: ", listType), err)
	}
	for _, v := range value {
		if err := WriteByType(ctx, protocol, listType, v); err != nil {
			return err
		}
	}
	if err := protocol.WriteMapEnd(ctx); err != nil {
		return thrift.PrependError("write set list error: ", err)
	}
	return nil
}

func WriteByType(ctx context.Context, protocol thrift.TProtocol, fieldType *FieldType, value interface{}) (err error) {
	switch fieldType.TypeId {
	case thrift.BOOL:
		err = protocol.WriteBool(ctx, getBool(value))
	case thrift.BYTE:
		err = protocol.WriteByte(ctx, getByte(value))
	case thrift.DOUBLE:
		err = protocol.WriteDouble(ctx, getDouble(value))
	case thrift.I16:
		err = protocol.WriteI16(ctx, getInt16(value))
	case thrift.I32:
		err = protocol.WriteI32(ctx, getInt32(value))
	case thrift.I64:
		err = protocol.WriteI64(ctx, getInt64(value))
	case thrift.STRING:
		err = protocol.WriteString(ctx, getString(value))
	case thrift.STRUCT:
		err = WriteStructFields(ctx, protocol, fieldType.Struct.Name, fieldType.Struct.Fields, value.(map[string]interface{}))
	case thrift.MAP:
		err = WriteMap(ctx, protocol, fieldType.GenericTypes[0], fieldType.GenericTypes[1], value.(map[string]interface{}))
	case thrift.SET:
		err = WriteSet(ctx, protocol, fieldType.GenericTypes[0], value.([]interface{}))
	case thrift.LIST:
		err = WriteList(ctx, protocol, fieldType.GenericTypes[0], value.([]interface{}))
	default:
		return thrift.PrependError(fmt.Sprintf("%T type error: ", fieldType), errors.New("type unknown"))
	}
	return
}

func ReadStructFields(ctx context.Context, inProtocol thrift.TProtocol, fields []*Field) (map[string]interface{}, error) {
	fmt.Println("ReadStructFields fields:", fields)
	if _, err := inProtocol.ReadStructBegin(ctx); err != nil {
		return nil, thrift.PrependError(fmt.Sprintf("%T read error: ", fields), err)
	}

	fieldMap := map[int16]*Field{}
	for _, one := range fields {
		fieldMap[one.Num] = one
	}
	value := make(map[string]interface{})

	for {
		_, fieldTypeId, fieldId, err := inProtocol.ReadFieldBegin(ctx)
		if err != nil {
			return nil, thrift.PrependError(fmt.Sprintf("%T field %d read error: ", fields, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		field, ok := fieldMap[fieldId]

		fmt.Println("ReadStructFields find by fieldId:", fieldId, ",find:", ok, ",field:", field)
		if !ok {
			if err = inProtocol.Skip(ctx, fieldTypeId); err != nil {
				return nil, err
			}
			// 字段不存在
		}
		var v interface{}
		if v, err = ReadStructField(ctx, inProtocol, field, fieldTypeId); err != nil {
			return nil, err
		}
		value[field.Name] = v

		if err = inProtocol.ReadFieldEnd(ctx); err != nil {
			return nil, err
		}
	}
	if err := inProtocol.ReadStructEnd(ctx); err != nil {
		return nil, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", fields), err)
	}
	return value, nil
}

func ReadStructField(ctx context.Context, outProtocol thrift.TProtocol, field *Field, fieldTypeId thrift.TType) (interface{}, error) {
	fmt.Println("ReadStructField field:", field, ",fieldTypeId:", fieldTypeId)
	var v interface{}
	var err error
	if v, err = ReadByType(ctx, outProtocol, field.Type, fieldTypeId); err != nil {
		return nil, thrift.PrependError(fmt.Sprintf("error reading field %d:%s : ", field.Num, field.Name), err)
	}
	return v, nil
}

func ReadMap(ctx context.Context, protocol thrift.TProtocol, keyType *FieldType, valueType *FieldType) (interface{}, error) {
	res := map[interface{}]interface{}{}
	var keyTypeId thrift.TType
	var valueTypeId thrift.TType
	var size int
	var err error
	if keyTypeId, valueTypeId, size, err = protocol.ReadMapBegin(ctx); err != nil {
		return nil, thrift.PrependError(fmt.Sprintf("%T read map begin error: ", keyType), err)
	}
	for i := 0; i < size; i++ {
		var k interface{}
		if k, err = ReadByType(ctx, protocol, keyType, keyTypeId); err != nil {
			return nil, err
		}
		var v interface{}
		if v, err = ReadByType(ctx, protocol, valueType, valueTypeId); err != nil {
			return nil, err
		}
		res[k] = v
	}
	if err = protocol.ReadMapEnd(ctx); err != nil {
		return nil, thrift.PrependError("read map end error: ", err)
	}
	return res, nil
}

func ReadSet(ctx context.Context, protocol thrift.TProtocol, setType *FieldType) (interface{}, error) {
	var res []interface{}
	var setTypeId thrift.TType
	var size int
	var err error
	if setTypeId, size, err = protocol.ReadSetBegin(ctx); err != nil {
		return nil, thrift.PrependError(fmt.Sprintf("%T read set begin error: ", setType), err)
	}
	for i := 0; i < size; i++ {
		var v interface{}
		if v, err = ReadByType(ctx, protocol, setType, setTypeId); err != nil {
			return nil, err
		}
		res = append(res, v)
	}
	if err := protocol.WriteMapEnd(ctx); err != nil {
		return nil, thrift.PrependError("read set end error: ", err)
	}
	return res, nil
}

func ReadList(ctx context.Context, protocol thrift.TProtocol, listType *FieldType) (interface{}, error) {
	var res []interface{}
	var listTypeId thrift.TType
	var size int
	var err error
	if listTypeId, size, err = protocol.ReadListBegin(ctx); err != nil {
		return nil, thrift.PrependError(fmt.Sprintf("%T read list begin error: ", listType), err)
	}
	for i := 0; i < size; i++ {
		var v interface{}
		if v, err = ReadByType(ctx, protocol, listType, listTypeId); err != nil {
			return nil, err
		}
		res = append(res, v)
	}
	if err := protocol.WriteMapEnd(ctx); err != nil {
		return nil, thrift.PrependError("read set list error: ", err)
	}
	return nil, nil
}

func ReadByType(ctx context.Context, protocol thrift.TProtocol, fieldType *FieldType, fieldTypeId thrift.TType) (res interface{}, err error) {
	// 判断类型是否一致
	if fieldTypeId != fieldType.TypeId {
		if err = protocol.Skip(ctx, fieldTypeId); err != nil {
			return nil, err
		}
	}
	switch fieldType.TypeId {
	case thrift.BOOL:
		res, err = protocol.ReadBool(ctx)
	case thrift.BYTE:
		res, err = protocol.ReadByte(ctx)
	case thrift.DOUBLE:
		res, err = protocol.ReadDouble(ctx)
	case thrift.I16:
		res, err = protocol.ReadI16(ctx)
	case thrift.I32:
		res, err = protocol.ReadI32(ctx)
	case thrift.I64:
		res, err = protocol.ReadI64(ctx)
	case thrift.STRING:
		res, err = protocol.ReadString(ctx)
	case thrift.STRUCT:
		res, err = ReadStructFields(ctx, protocol, fieldType.Struct.Fields)
	case thrift.MAP:
		res, err = ReadMap(ctx, protocol, fieldType.GenericTypes[0], fieldType.GenericTypes[1])
	case thrift.SET:
		res, err = ReadSet(ctx, protocol, fieldType.GenericTypes[0])
	case thrift.LIST:
		res, err = ReadList(ctx, protocol, fieldType.GenericTypes[0])
	default:
		return nil, thrift.PrependError(fmt.Sprintf("%T type error: ", fieldType), errors.New("type unknown"))
	}
	return
}

func getBool(v interface{}) (res bool) {
	if v == nil {
		return
	}
	res, ok := v.(bool)
	if ok {
		return
	}
	return
}

func getByte(v interface{}) (res int8) {
	if v == nil {
		return
	}
	res, ok := v.(int8)
	if ok {
		return
	}
	return
}

func getDouble(v interface{}) (res float64) {
	if v == nil {
		return
	}
	res, ok := v.(float64)
	if ok {
		return
	}
	return
}

func getInt16(v interface{}) (res int16) {
	if v == nil {
		return
	}
	res, ok := v.(int16)
	if ok {
		return
	}
	return
}

func getInt32(v interface{}) (res int32) {
	if v == nil {
		return
	}
	res, ok := v.(int32)
	if ok {
		return
	}
	return
}

func getInt64(v interface{}) (res int64) {
	if v == nil {
		return
	}
	res, ok := v.(int64)
	if ok {
		return
	}
	return
}

func getString(v interface{}) (res string) {
	if v == nil {
		return
	}
	res, ok := v.(string)
	if ok {
		return
	}
	return
}
