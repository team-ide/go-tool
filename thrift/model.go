package thrift

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/team-ide/go-tool/util"
	"strconv"
)

func toJSON(v interface{}) (res string) {
	res = util.GetStringValue(v)
	return
}

type Struct struct {
	Name   string   `json:"name"`
	Fields []*Field `json:"fields"`
}

type Field struct {
	Num  int16      `json:"num"`
	Name string     `json:"name"`
	Type *FieldType `json:"type"`
}

type FieldType struct {
	TypeId        thrift.TType `json:"typeId"`
	StructInclude string       `json:"structInclude"`
	StructName    string       `json:"structName"`
	structObj     *Struct
	SetType       *FieldType `json:"setType"`
	ListType      *FieldType `json:"listType"`
	MapKeyType    *FieldType `json:"mapKeyType"`
	MapValueType  *FieldType `json:"mapValueType"`
}

func WriteStructFields(ctx context.Context, protocol thrift.TProtocol, name string, fields []*Field, value map[string]interface{}) error {

	fmt.Println("WriteStructFields name:", name, ",fields:", toJSON(fields), ",value:", toJSON(value))
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
	fmt.Println("WriteStructField field:", toJSON(field), ",value:", toJSON(value))
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

func WriteMap(ctx context.Context, protocol thrift.TProtocol, keyType *FieldType, valueType *FieldType, value map[interface{}]interface{}) error {
	if value == nil {
		value = map[interface{}]interface{}{}
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
		data, ok := value.(map[string]interface{})
		if !ok {
			data = map[string]interface{}{}
			bs, _ := json.Marshal(value)
			_ = json.Unmarshal(bs, &data)
		}
		err = WriteStructFields(ctx, protocol, fieldType.structObj.Name, fieldType.structObj.Fields, data)
	case thrift.MAP:
		data, ok := value.(map[interface{}]interface{})
		if !ok {
			data = map[interface{}]interface{}{}
			bs, _ := json.Marshal(value)
			_ = json.Unmarshal(bs, &data)
		}
		err = WriteMap(ctx, protocol, fieldType.MapKeyType, fieldType.MapValueType, data)
	case thrift.SET:
		data, ok := value.([]interface{})
		if !ok {
			data = []interface{}{}
			bs, _ := json.Marshal(value)
			_ = json.Unmarshal(bs, &data)
		}
		err = WriteSet(ctx, protocol, fieldType.SetType, data)
	case thrift.LIST:
		data, ok := value.([]interface{})
		if !ok {
			data = []interface{}{}
			bs, _ := json.Marshal(value)
			_ = json.Unmarshal(bs, &data)
		}
		err = WriteList(ctx, protocol, fieldType.ListType, value.([]interface{}))
	default:
		return thrift.PrependError(fmt.Sprintf("%T type error: ", fieldType), errors.New("type unknown"))
	}
	return
}

func ReadStructFields(ctx context.Context, inProtocol thrift.TProtocol, fields []*Field) (map[string]interface{}, error) {
	fmt.Println("ReadStructFields fields:", toJSON(fields))
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

		//fmt.Println("ReadStructFields find by fieldId:", fieldId, ",find:", ok, ",field:", toJSON(field))
		if !ok {
			if err = inProtocol.Skip(ctx, fieldTypeId); err != nil {
				return nil, err
			}

			fmt.Println("ReadStructFields fields:", toJSON(fields))
			// 字段不存在
			return nil, errors.New(fmt.Sprintf("ReadStructFields field %d %d not found", fieldId, fieldTypeId))
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
	fmt.Println("ReadStructField field:", toJSON(field), ",fieldTypeId:", fieldTypeId)
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
	//if fieldTypeId != fieldType.TypeId {
	//	if err = protocol.Skip(ctx, fieldTypeId); err != nil {
	//		return nil, err
	//	}
	//}
	switch fieldTypeId {
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
		res, err = ReadStructFields(ctx, protocol, fieldType.structObj.Fields)
	case thrift.MAP:
		res, err = ReadMap(ctx, protocol, fieldType.MapKeyType, fieldType.MapValueType)
	case thrift.SET:
		res, err = ReadSet(ctx, protocol, fieldType.SetType)
	case thrift.LIST:
		res, err = ReadList(ctx, protocol, fieldType.ListType)
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
	str := util.GetStringValue(v)
	i64, _ := strconv.ParseInt(str, 10, 64)
	res = int8(i64)
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
	str := util.GetStringValue(v)
	i64, _ := strconv.ParseFloat(str, 64)
	res = i64
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
	str := util.GetStringValue(v)
	i64, _ := strconv.ParseInt(str, 10, 64)
	res = int16(i64)
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
	str := util.GetStringValue(v)
	i64, _ := strconv.ParseInt(str, 10, 64)
	res = int32(i64)
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
	str := util.GetStringValue(v)
	i64, _ := strconv.ParseInt(str, 10, 64)
	res = i64
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
	res = util.GetStringValue(v)
	return
}
