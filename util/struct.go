package util

import (
	"reflect"
	"strings"
	"sync"
	"time"
)

func NewStructCache() *StructCache {
	cache := new(StructCache)
	cache.cache = make(map[reflect.Type]*StructInfo)
	return cache
}

type StructCache struct {
	cache     map[reflect.Type]*StructInfo
	cacheLock sync.Mutex
}

type StructInfo struct {
	IsMap    bool         `json:"isMap,omitempty"`
	IsList   bool         `json:"isList,omitempty"`
	IsStruct bool         `json:"isStruct,omitempty"`
	Type     reflect.Type `json:"type,omitempty"`
	Fields   []*FieldInfo `json:"fields,omitempty"`

	NameMap   map[string]*FieldInfo `json:"nameMap,omitempty"`
	ColumnMap map[string]*FieldInfo `json:"columnMap,omitempty"`
	JsonMap   map[string]*FieldInfo `json:"jsonMap,omitempty"`
	BsonMap   map[string]*FieldInfo `json:"bsonMap,omitempty"`
	YamlMap   map[string]*FieldInfo `json:"yamlMap,omitempty"`
	XmlMap    map[string]*FieldInfo `json:"xmlMap,omitempty"`
	TomlMap   map[string]*FieldInfo `json:"tomlMap,omitempty"`
	KeyMap    map[string]*FieldInfo `json:"keyMap,omitempty"`

	Extend any `json:"extend,omitempty"`
}

func (this_ *StructInfo) Get(key string) *FieldInfo {

	findKey := "," + strings.ToLower(key) + ","
	for k, v := range this_.KeyMap {
		if strings.Contains(k, findKey) {
			return v
		}
	}
	return nil
}

type FieldInfo struct {
	Name   string `json:"name,omitempty"`
	Index  int    `json:"index,omitempty"`
	Column string `json:"column,omitempty"`
	Json   string `json:"json,omitempty"`
	Bson   string `json:"bson,omitempty"`
	Xml    string `json:"xml,omitempty"`
	Yaml   string `json:"yaml,omitempty"`
	Toml   string `json:"toml,omitempty"`

	IsString bool `json:"isString,omitempty"`
	IsNumber bool `json:"isNumber,omitempty"`
	IsBool   bool `json:"isBool,omitempty"`
	IsStruct bool `json:"isStruct,omitempty"`
	IsMap    bool `json:"isMap,omitempty"`
	IsList   bool `json:"isList,omitempty"`

	StructType reflect.Type `json:"structType,omitempty"`

	Extend any `json:"extend,omitempty"`

	structField reflect.StructField
	value       reflect.Value
}

func (this_ *FieldInfo) GetStructField() reflect.StructField {
	return this_.structField
}

func (this_ *FieldInfo) GetValue() reflect.Value {
	return this_.value
}

func (this_ *StructInfo) GetFieldValue(data any, key string) (field *FieldInfo, fV reflect.Value, find bool) {
	field = this_.Get(key)
	if field == nil {
		return
	}
	dV := reflect.ValueOf(data) // 取得struct变量的指针
	for dV.Kind() == reflect.Ptr {
		dV = dV.Elem()
	}

	fV = dV.FieldByName(field.Name)
	find = true
	return
}

func (this_ *StructInfo) SetFieldValue(data any, name string, value any) (ok bool, err error) {
	_, fieldValue, find := this_.GetFieldValue(data, name)
	if !find {
		return
	}
	if !fieldValue.CanSet() {
		return
	}
	fieldType := fieldValue.Type()
	valueType := reflect.TypeOf(value)
	if valueType == fieldType {
		ok = true
		fieldValue.Set(reflect.ValueOf(value))
		return
	}
	var isPtr = fieldType.Kind() == reflect.Ptr
	if isPtr {
		fieldType = fieldType.Elem()
	}
	fieldTypeKind := fieldType.Kind()
	if fieldTypeKind == reflect.String {
		var str = GetStringValue(value)
		if isPtr {
			fieldValue.Set(reflect.ValueOf(&str))
		} else {
			fieldValue.SetString(str)
		}
	} else if fieldTypeKind >= reflect.Int && fieldTypeKind <= reflect.Int64 {
		var num int64
		if num, err = ValueToInt64(value); err != nil {
			return
		}
		if isPtr {
			switch fieldTypeKind {
			case reflect.Int:
				v := int(num)
				fieldValue.Set(reflect.ValueOf(&v))
			case reflect.Int8:
				v := int8(num)
				fieldValue.Set(reflect.ValueOf(&v))
			case reflect.Int16:
				v := int16(num)
				fieldValue.Set(reflect.ValueOf(&v))
			case reflect.Int32:
				v := int32(num)
				fieldValue.Set(reflect.ValueOf(&v))
			default:
				fieldValue.Set(reflect.ValueOf(&num))
			}
		} else {
			fieldValue.SetInt(num)
		}
	} else if fieldTypeKind >= reflect.Uint && fieldTypeKind <= reflect.Uint64 {
		var num uint64
		if num, err = ValueToUint64(value); err != nil {
			return
		}
		if isPtr {
			switch fieldTypeKind {
			case reflect.Uint:
				v := uint(num)
				fieldValue.Set(reflect.ValueOf(&v))
			case reflect.Uint8:
				v := uint8(num)
				fieldValue.Set(reflect.ValueOf(&v))
			case reflect.Uint16:
				v := uint16(num)
				fieldValue.Set(reflect.ValueOf(&v))
			case reflect.Uint32:
				v := uint32(num)
				fieldValue.Set(reflect.ValueOf(&v))
			default:
				fieldValue.Set(reflect.ValueOf(&num))
			}
		} else {
			fieldValue.SetUint(num)
		}
	} else if fieldTypeKind == reflect.Float32 || fieldTypeKind == reflect.Float64 {
		var num float64
		if num, err = ValueToFloat64(value); err != nil {
			return
		}
		if isPtr {
			switch fieldTypeKind {
			case reflect.Float32:
				v := float32(num)
				fieldValue.Set(reflect.ValueOf(&v))
			default:
				fieldValue.Set(reflect.ValueOf(&num))
			}
		} else {
			fieldValue.SetFloat(num)
		}
	} else if fieldTypeKind == reflect.Bool {
		var v bool
		v, isB := value.(bool)
		if !isB {
			str := GetStringValue(value)
			if str == "1" || str == "true" {
				v = true
			}
		}
		if isPtr {
			fieldValue.Set(reflect.ValueOf(&v))
		} else {
			fieldValue.SetBool(v)
		}
	} else if fieldType.String() == "time.Time" {
		var v time.Time
		v, isT := value.(time.Time)
		if !isT {
			var num int64
			if num, err = ValueToInt64(value); err != nil {
				return
			}
			v = time.UnixMilli(num)
		}
		if isPtr {
			fieldValue.Set(reflect.ValueOf(&v))
		} else {
			fieldValue.Set(reflect.ValueOf(v))
		}
	} else {
		return
	}
	ok = true
	return
}

func (this_ *StructCache) GetStructInfo(structType reflect.Type) (info *StructInfo) {
	for structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
	}

	this_.cacheLock.Lock()
	defer this_.cacheLock.Unlock()

	info = this_.cache[structType]
	if info != nil {
		return
	}

	info = CreateStructInfo(structType)
	this_.cache[structType] = info
	return
}

func CreateStructInfo(structType reflect.Type) (info *StructInfo) {

	for structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
	}
	info = &StructInfo{}
	info.Type = structType
	if structType.Kind() == reflect.Map {
		info.IsMap = true
		return
	}
	if structType.Kind() == reflect.Slice {
		info.IsList = true
		return
	}
	if structType.Kind() != reflect.Struct {
		return
	}

	info.IsStruct = true
	info.NameMap = make(map[string]*FieldInfo)
	info.ColumnMap = make(map[string]*FieldInfo)
	info.JsonMap = make(map[string]*FieldInfo)
	info.BsonMap = make(map[string]*FieldInfo)
	info.XmlMap = make(map[string]*FieldInfo)
	info.YamlMap = make(map[string]*FieldInfo)
	info.TomlMap = make(map[string]*FieldInfo)
	info.KeyMap = make(map[string]*FieldInfo)
	for i := 0; i < structType.NumField(); i++ {
		field := &FieldInfo{
			structField: structType.Field(i),
			Index:       i,
		}
		field.Name = field.structField.Name
		field.Json = field.structField.Tag.Get("json")
		field.Column = field.structField.Tag.Get("column")
		field.Xml = field.structField.Tag.Get("xml")
		field.Yaml = field.structField.Tag.Get("yaml")
		field.Toml = field.structField.Tag.Get("toml")
		field.Bson = field.structField.Tag.Get("bson")
		var keys []string
		info.NameMap[field.Name] = field
		keys = append(keys, field.Name)
		if field.Json != "" && field.Json != "-" {
			field.Json = strings.Split(field.Json, ",")[0]
			info.JsonMap[field.Json] = field
			keys = append(keys, field.Json)
		}
		if field.Column != "" && field.Column != "-" {
			field.Column = strings.Split(field.Column, ",")[0]
			info.ColumnMap[field.Column] = field
			keys = append(keys, field.Column)
		}
		if field.Xml != "" && field.Xml != "-" {
			field.Xml = strings.Split(field.Xml, ",")[0]
			info.XmlMap[field.Xml] = field
			keys = append(keys, field.Xml)
		}
		if field.Yaml != "" && field.Yaml != "-" {
			field.Yaml = strings.Split(field.Yaml, ",")[0]
			info.YamlMap[field.Yaml] = field
			keys = append(keys, field.Yaml)
		}
		if field.Toml != "" && field.Toml != "-" {
			field.Toml = strings.Split(field.Toml, ",")[0]
			info.TomlMap[field.Toml] = field
			keys = append(keys, field.Toml)
		}
		if field.Bson != "" && field.Bson != "-" {
			field.Bson = strings.Split(field.Bson, ",")[0]
			info.TomlMap[field.Bson] = field
			keys = append(keys, field.Bson)
		}
		info.KeyMap[","+strings.ToLower(strings.Join(keys, ","))+","] = field
		fullFieldValue(field)
		info.Fields = append(info.Fields, field)
	}
	return
}

func GetMapFields(v any) (fields []*FieldInfo) {
	objV := reflect.ValueOf(v)
	for objV.Kind() == reflect.Ptr {
		objV = objV.Elem()
	}

	for _, kV := range objV.MapKeys() {
		if kV.Type().Kind() != reflect.String {
			continue
		}
		k := kV.String()
		vV := objV.MapIndex(kV)

		field := &FieldInfo{
			Name: k,
			structField: reflect.StructField{
				Name: k,
				Type: vV.Type(),
			},
			value: vV,
		}
		fullFieldValue(field)
		fields = append(fields, field)
	}
	return
}

func fullFieldValue(field *FieldInfo) {
	fT := field.structField.Type
	fTKind := fT.Kind()
	for fTKind == reflect.Ptr {
		fT = fT.Elem()
		fTKind = fT.Kind()
	}
	field.IsString = fTKind == reflect.String
	field.IsNumber = fTKind >= reflect.Int && fTKind <= reflect.Uint64
	field.IsBool = fTKind == reflect.Bool
	field.IsMap = fTKind == reflect.Map
	field.IsList = fTKind == reflect.Slice
	field.IsStruct = fTKind == reflect.Struct
	if field.IsMap || field.IsList {
		sT := fT.Elem()
		for sT.Kind() == reflect.Ptr {
			sT = sT.Elem()
		}
		field.StructType = sT
	} else if field.IsStruct {
		field.StructType = fT
	}

}
