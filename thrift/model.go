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

func toJSON(v any) (res string) {
	res = util.GetStringValue(v)
	return
}

type Struct struct {
	Name   string   `json:"name"`
	Fields []*Field `json:"fields"`
	data   map[string]any
}

func (this_ *Struct) GetData() map[string]any {
	if this_ == nil {
		return nil
	}
	return this_.data
}

func (this_ *Struct) SetData(data map[string]any) {
	if this_ == nil {
		return
	}
	this_.data = data
}

func (this_ *Struct) String() string {
	if this_ == nil {
		return "<nil>"
	}
	return fmt.Sprintf(this_.Name+"(%+v)", *this_)
}

func (this_ *Struct) Read(ctx context.Context, inProtocol thrift.TProtocol) (err error) {
	this_.data, err = ReadStructFields(ctx, inProtocol, this_.Fields)
	if err != nil {
		return err
	}
	return nil
}

func (this_ *Struct) Write(ctx context.Context, outProtocol thrift.TProtocol) error {
	err := WriteStructFields(ctx, outProtocol, this_.Name, this_.Fields, this_.data)
	if err != nil {
		return err
	}

	return nil
}

var (
	TSerializer   = thrift.NewTSerializer()
	TDeserializer = thrift.NewTDeserializer()
)

func Serialize(s *Struct) (bs []byte, err error) {
	return TSerializer.Write(context.Background(), s)
}
func Deserialize(s *Struct, bs []byte) (err error) {
	return TDeserializer.Read(context.Background(), s, bs)
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

func GetFieldType(typeId thrift.TType) *FieldType {
	return &FieldType{TypeId: typeId}
}

func WriteStructFields(ctx context.Context, protocol thrift.TProtocol, name string, fields []*Field, value map[string]any) error {

	//fmt.Println("WriteStructFields name:", name, ",fields:", toJSON(fields), ",value:", toJSON(value))
	if err := protocol.WriteStructBegin(ctx, name); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", name), err)
	}
	for _, filed := range fields {
		v, ok := value[filed.Name]
		if !ok {
			continue
		}
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

func WriteStructField(ctx context.Context, protocol thrift.TProtocol, field *Field, value any) error {
	//fmt.Println("WriteStructField field:", toJSON(field), ",value:", toJSON(value))
	var err error
	typeId := field.Type.TypeId
	if typeId == BINARY {
		typeId = thrift.STRING
	}
	if err = protocol.WriteFieldBegin(ctx, field.Name, typeId, field.Num); err != nil {
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

func WriteMap(ctx context.Context, protocol thrift.TProtocol, keyType *FieldType, valueType *FieldType, value map[any]any) error {
	if value == nil {
		value = map[any]any{}
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

func WriteSet(ctx context.Context, protocol thrift.TProtocol, setType *FieldType, value []any) error {
	size := len(value)
	if err := protocol.WriteSetBegin(ctx, setType.TypeId, size); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write set begin error: ", setType), err)
	}
	for _, v := range value {
		if err := WriteByType(ctx, protocol, setType, v); err != nil {
			return err
		}
	}
	if err := protocol.WriteSetEnd(ctx); err != nil {
		return thrift.PrependError("write set end error: ", err)
	}
	return nil
}

func WriteList(ctx context.Context, protocol thrift.TProtocol, listType *FieldType, value []any) error {
	size := len(value)
	if err := protocol.WriteListBegin(ctx, listType.TypeId, size); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write list begin error: ", listType), err)
	}
	for _, v := range value {
		if err := WriteByType(ctx, protocol, listType, v); err != nil {
			return err
		}
	}
	if err := protocol.WriteListEnd(ctx); err != nil {
		return thrift.PrependError("write list end error: ", err)
	}
	return nil
}

var (
	BINARY thrift.TType = 18 // 兼容 thrift.BINARY 类型 即  []byte
)

func WriteByType(ctx context.Context, protocol thrift.TProtocol, fieldType *FieldType, value any) (err error) {

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
	case BINARY: // 兼容 thrift.BINARY 类型 即  []byte
		err = protocol.WriteString(ctx, getString(value))
	case thrift.STRUCT:
		data, ok := value.(map[string]any)
		if !ok {
			data = map[string]any{}
			bs, _ := util.ObjToJsonBytes(value)
			_ = util.JSONDecodeUseNumber(bs, &data)
		}
		err = WriteStructFields(ctx, protocol, fieldType.structObj.Name, fieldType.structObj.Fields, data)
	case thrift.MAP:
		data, ok := value.(map[any]any)
		if !ok {
			data = map[any]any{}
			bs, _ := util.ObjToJsonBytes(value)

			strMap := map[string]any{}
			if e := util.JSONDecodeUseNumber(bs, &strMap); e == nil {
				for k, v := range strMap {
					data[k] = v
				}
			} else {
				intMap := map[int64]any{}
				if e = util.JSONDecodeUseNumber(bs, &intMap); e == nil {
					for k, v := range intMap {
						data[k] = v
					}
				}
			}

		}
		err = WriteMap(ctx, protocol, fieldType.MapKeyType, fieldType.MapValueType, data)
	case thrift.SET:
		data, ok := value.([]any)
		if !ok {
			data = []any{}
			bs, _ := util.ObjToJsonBytes(value)
			_ = util.JSONDecodeUseNumber(bs, &data)
		}
		err = WriteSet(ctx, protocol, fieldType.SetType, data)
	case thrift.LIST:
		data, ok := value.([]any)
		if !ok {
			data = []any{}
			bs, _ := util.ObjToJsonBytes(value)
			_ = util.JSONDecodeUseNumber(bs, &data)
		}
		err = WriteList(ctx, protocol, fieldType.ListType, data)
	default:
		return thrift.PrependError(fmt.Sprintf("%T type error: ", fieldType), errors.New("type unknown"))
	}
	return
}

func ReadStructFields(ctx context.Context, inProtocol thrift.TProtocol, fields []*Field) (map[string]any, error) {
	//fmt.Println("ReadStructFields fields:", toJSON(fields))
	if _, err := inProtocol.ReadStructBegin(ctx); err != nil {
		return nil, thrift.PrependError(fmt.Sprintf("%T read error: ", fields), err)
	}

	fieldMap := map[int16]*Field{}
	for _, one := range fields {
		fieldMap[one.Num] = one
	}
	value := make(map[string]any)

	for {
		_, fieldTypeId, fieldId, err := inProtocol.ReadFieldBegin(ctx)
		//fmt.Println("read fieldTypeId:", fieldTypeId, ",fieldId:", fieldId)
		if err != nil {
			return nil, thrift.PrependError(fmt.Sprintf("%T field %d read error: ", fields, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		field, ok := fieldMap[fieldId]

		//fmt.Println("ReadStructFields find by fieldId:", fieldId, ",find:", ok, ",field:", toJSON(field))
		if !ok { // 如果字段不存在，也继续解析
			//

			//if err = inProtocol.Skip(ctx, fieldTypeId); err != nil {
			//	return nil, err
			//}
			//
			////fmt.Println("ReadStructFields fields:", toJSON(fields))
			//// 字段不存在
			//return nil, errors.New(fmt.Sprintf("ReadStructFields field %d %d not found", fieldId, fieldTypeId))
		}
		var v any
		if v, err = ReadStructField(ctx, inProtocol, field, fieldId, fieldTypeId); err != nil {
			return nil, err
		}
		if field == nil {
			name := fmt.Sprintf("field-unknow-%d", fieldId)
			value[name] = v
		} else {
			value[field.Name] = v
		}

		if err = inProtocol.ReadFieldEnd(ctx); err != nil {
			return nil, err
		}
	}
	if err := inProtocol.ReadStructEnd(ctx); err != nil {
		return nil, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", fields), err)
	}
	return value, nil
}

func ReadStructField(ctx context.Context, outProtocol thrift.TProtocol, field *Field, fieldId int16, fieldTypeId thrift.TType) (any, error) {
	//fmt.Println("ReadStructField field:", toJSON(field), ",fieldTypeId:", fieldTypeId)
	var v any
	var err error
	var fieldType *FieldType
	if field != nil {
		fieldType = field.Type
	}
	if v, err = ReadByType(ctx, outProtocol, fieldType, fieldTypeId); err != nil {
		return nil, thrift.PrependError(fmt.Sprintf("error reading field %d:%d : ", fieldId, fieldTypeId), err)
	}
	return v, nil
}

func ReadMap(ctx context.Context, protocol thrift.TProtocol, keyType *FieldType, valueType *FieldType) (res any, err error) {
	var keyTypeId thrift.TType
	var valueTypeId thrift.TType
	var size int
	if keyTypeId, valueTypeId, size, err = protocol.ReadMapBegin(ctx); err != nil {
		return nil, thrift.PrependError(fmt.Sprintf("%T read map begin error: ", keyType), err)
	}
	var doubleMap = map[float64]any{}
	var intMap = map[int64]any{}
	var stringMap = map[string]any{}
	switch keyTypeId {
	case thrift.DOUBLE:
		res = doubleMap
	case thrift.BYTE, thrift.I16, thrift.I32, thrift.I64:
		res = intMap
	default:
		res = stringMap
	}
	for i := 0; i < size; i++ {
		var k any
		if k, err = ReadByType(ctx, protocol, keyType, keyTypeId); err != nil {
			return nil, err
		}
		var v any
		if v, err = ReadByType(ctx, protocol, valueType, valueTypeId); err != nil {
			return nil, err
		}

		switch keyTypeId {
		case thrift.DOUBLE:
			doubleMap[util.StringToFloat64(util.GetStringValue(k))] = v
		case thrift.BYTE, thrift.I16, thrift.I32, thrift.I64:
			intMap[util.StringToInt64(util.GetStringValue(k))] = v
		default:
			stringMap[util.GetStringValue(k)] = v
		}
	}
	if err = protocol.ReadMapEnd(ctx); err != nil {
		return nil, thrift.PrependError("read map end error: ", err)
	}
	return res, nil
}

func ReadSet(ctx context.Context, protocol thrift.TProtocol, setType *FieldType) (any, error) {
	var res = make([]any, 0)
	var setTypeId thrift.TType
	var size int
	var err error
	if setTypeId, size, err = protocol.ReadSetBegin(ctx); err != nil {
		return nil, thrift.PrependError(fmt.Sprintf("%T read set begin error: ", setType), err)
	}
	for i := 0; i < size; i++ {
		var v any
		if v, err = ReadByType(ctx, protocol, setType, setTypeId); err != nil {
			return nil, err
		}
		res = append(res, v)
	}
	if err := protocol.ReadSetEnd(ctx); err != nil {
		return nil, thrift.PrependError("read set end error: ", err)
	}
	return res, nil
}

func ReadList(ctx context.Context, protocol thrift.TProtocol, listType *FieldType) (any, error) {
	var res = make([]any, 0)
	var listTypeId thrift.TType
	var size int
	var err error
	if listTypeId, size, err = protocol.ReadListBegin(ctx); err != nil {
		return nil, thrift.PrependError(fmt.Sprintf("%T read list begin error: ", listType), err)
	}
	//fmt.Println("ReadList listTypeId:", listTypeId, ",size:", size)
	for i := 0; i < size; i++ {
		var v any
		if v, err = ReadByType(ctx, protocol, listType, listTypeId); err != nil {
			return nil, err
		}
		res = append(res, v)
	}
	if err := protocol.ReadListEnd(ctx); err != nil {
		return nil, thrift.PrependError("read list end error: ", err)
	}
	return res, nil
}

func ReadByType(ctx context.Context, protocol thrift.TProtocol, fieldType *FieldType, fieldTypeId thrift.TType) (res any, err error) {
	// 判断类型是否一致
	//if fieldTypeId != fieldType.TypeId {
	//	if err = protocol.Skip(ctx, fieldTypeId); err != nil {
	//		return nil, err
	//	}
	//}
	switch fieldTypeId {
	case thrift.VOID:
	case thrift.STOP:
	case thrift.BOOL:
		res, err = protocol.ReadBool(ctx)
	//case thrift.UUID:
	//	res, err = protocol.ReadUUID(ctx)
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
	case BINARY:
		res, err = protocol.ReadString(ctx)
		if res != nil {
			res = []byte(getString(res))
		} else {
			res = []byte{}
		}
	case thrift.STRUCT:
		var fields []*Field
		if fieldType != nil && fieldType.structObj != nil {
			fields = fieldType.structObj.Fields
		}
		res, err = ReadStructFields(ctx, protocol, fields)
	case thrift.MAP:
		var keyType *FieldType
		var valueType *FieldType
		if fieldType != nil {
			keyType = fieldType.MapKeyType
			valueType = fieldType.MapValueType
		}
		res, err = ReadMap(ctx, protocol, keyType, valueType)
	case thrift.SET:
		var setType *FieldType
		if fieldType != nil {
			setType = fieldType.SetType
		}
		res, err = ReadSet(ctx, protocol, setType)
	case thrift.LIST:
		var listType *FieldType
		if fieldType != nil {
			listType = fieldType.ListType
		}
		res, err = ReadList(ctx, protocol, listType)
	default:
		return nil, thrift.PrependError(fmt.Sprintf("%T type error: ", fieldType), errors.New("type unknown"))
	}
	return
}

func getBool(v any) (res bool) {
	if v == nil {
		return
	}
	res, ok := v.(bool)
	if ok {
		return
	}
	return util.IsTrue(v)
}

func toInt64(v any) (res int64, ok bool) {
	if f, ok := v.(float64); ok {
		return int64(f), true
	}
	if f, ok := v.(float32); ok {
		return int64(f), true
	}
	if f, ok := v.(json.Number); ok {
		res, _ = f.Int64()
		return res, true
	}
	return
}

func getByte(v any) (res int8) {
	if v == nil {
		return
	}
	res, ok := v.(int8)
	if ok {
		return
	}
	if f, ok := toInt64(v); ok {
		return int8(f)
	}
	str := util.GetStringValue(v)
	i64, _ := strconv.ParseInt(str, 10, 64)
	res = int8(i64)
	return
}

func getDouble(v any) (res float64) {
	if v == nil {
		return
	}
	res, ok := v.(float64)
	if ok {
		return
	}
	if f, ok := v.(float32); ok {
		return float64(f)
	}
	str := util.GetStringValue(v)
	i64, _ := strconv.ParseFloat(str, 64)
	res = i64
	return
}

func getInt16(v any) (res int16) {
	if v == nil {
		return
	}
	res, ok := v.(int16)
	if ok {
		return
	}
	if f, ok := toInt64(v); ok {
		return int16(f)
	}
	str := util.GetStringValue(v)
	i64, _ := strconv.ParseInt(str, 10, 64)
	res = int16(i64)
	return
}

func getInt32(v any) (res int32) {
	if v == nil {
		return
	}
	res, ok := v.(int32)
	if ok {
		return
	}
	if f, ok := toInt64(v); ok {
		return int32(f)
	}
	str := util.GetStringValue(v)
	i64, _ := strconv.ParseInt(str, 10, 64)
	res = int32(i64)
	return
}

func getInt64(v any) (res int64) {
	if v == nil {
		return
	}
	res, ok := v.(int64)
	if ok {
		return
	}
	if f, ok := toInt64(v); ok {
		return f
	}
	str := util.GetStringValue(v)
	i64, _ := strconv.ParseInt(str, 10, 64)
	res = i64
	return
}

func getString(v any) (res string) {
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
