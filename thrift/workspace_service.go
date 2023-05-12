package thrift

import (
	"context"
	"errors"
	thrift2 "github.com/apache/thrift/lib/go/thrift"
	"github.com/team-ide/go-interpreter/thrift"
	"github.com/team-ide/go-tool/util"
)

func (this_ *Workspace) InvokeByServerAddress(serverAddress string, filename string, serviceName string, methodName string, args ...interface{}) (param *MethodParam, err error) {

	param, err = this_.GetMethodParam(filename, serviceName, methodName, args...)
	if err != nil {
		return
	}

	client, err := NewServiceClientByAddress(serverAddress)
	if err != nil {
		return
	}
	defer func() {
		_ = client.TTransport.Close()
	}()

	_, err = client.Send(context.Background(), param)
	return
}

func (this_ *Workspace) GetMethodParam(filename string, serviceName string, methodName string, args ...interface{}) (param *MethodParam, err error) {

	filename = util.FormatPath(filename)

	methodNode := this_.GetServiceMethod(filename, serviceName, methodName)
	if methodNode == nil {
		err = errors.New("service method node [" + filename + "][" + serviceName + "][" + methodName + "] not found")
		return
	}
	param = &MethodParam{
		Args: args,
		Name: methodName,
	}
	var structCache = map[string]*Struct{}
	param.ArgFields = this_.GetFields(filename, methodNode.Params, structCache)
	param.ResultType = this_.GetFieldTypeByNode(filename, methodNode.Return, structCache)
	param.ExceptionFields = this_.GetFields(filename, methodNode.Exceptions, structCache)

	return
}

func (this_ *Workspace) GetFields(filename string, params []*thrift.FieldNode, structCache map[string]*Struct) (fields []*Field) {

	filename = util.FormatPath(filename)

	for _, one := range params {
		fields = append(fields, this_.GetFieldByNode(filename, one, structCache))
	}

	return
}

func (this_ *Workspace) GetFieldByNode(filename string, fieldNode *thrift.FieldNode, structCache map[string]*Struct) (field *Field) {

	field = &Field{
		Name: fieldNode.Name,
		Num:  fieldNode.Num,
	}
	field.Type = this_.GetFieldTypeByNode(filename, fieldNode.Type, structCache)
	return
}

func (this_ *Workspace) GetFieldDemoData(filename string, fieldNode *Field) (res interface{}) {
	res = this_.GetFieldDemoDataByType(filename, fieldNode.Type)
	return
}

func (this_ *Workspace) GetFieldTypeByNode(filename string, fieldNode *thrift.FieldType, structCache map[string]*Struct) (fieldType *FieldType) {
	fieldType = &FieldType{
		TypeId: fieldNode.TypeId,
	}
	if fieldNode.StructName != "" {
		fieldType.StructName = fieldNode.StructName
		fieldType.StructInclude = fieldNode.StructInclude
		fieldType.structObj = this_.GetStructByName(filename, fieldNode.StructInclude, fieldNode.StructName, structCache)
	}
	if fieldNode.MapKeyType != nil {
		fieldType.MapKeyType = this_.GetFieldTypeByNode(filename, fieldNode.MapKeyType, structCache)
	}
	if fieldNode.MapValueType != nil {
		fieldType.MapValueType = this_.GetFieldTypeByNode(filename, fieldNode.MapValueType, structCache)
	}
	if fieldNode.ListType != nil {
		fieldType.ListType = this_.GetFieldTypeByNode(filename, fieldNode.ListType, structCache)
	}
	if fieldNode.SetType != nil {
		fieldType.SetType = this_.GetFieldTypeByNode(filename, fieldNode.SetType, structCache)
	}
	//fmt.Println("GetFieldTypeByNode filename:", filename, ",fieldNode:", toJSON(fieldNode), ",fieldType:", toJSON(fieldType))
	return
}

func (this_ *Workspace) GetFieldDemoDataByType(filename string, fieldNode *FieldType) (res interface{}) {

	if fieldNode.StructName != "" {
		return this_.GetStructDemoData(filename, fieldNode.StructInclude, fieldNode.StructName)
	}
	if fieldNode.MapKeyType != nil || fieldNode.MapValueType != nil {
		return map[string]interface{}{}
	}
	if fieldNode.ListType != nil {
		return []interface{}{}
	}
	if fieldNode.SetType != nil {
		return []interface{}{}
	}
	if fieldNode.TypeId == thrift2.STRING { // || fieldNode.TypeId == thrift2.UUID
		res = ""
	} else if fieldNode.TypeId == thrift2.BOOL {
		res = false
	} else {
		res = 0
	}
	//fmt.Println("GetFieldTypeByNode filename:", filename, ",fieldNode:", toJSON(fieldNode), ",fieldType:", toJSON(fieldType))
	return
}

func (this_ *Workspace) GetStructByName(filename string, include string, name string, structCache map[string]*Struct) (res *Struct) {
	defer func() {
		//fmt.Println("GetStructByName filename:", filename, ",include:", include, ",name:", name, ",res:", toJSON(res))
		structCache[include+"-"+name] = res
	}()
	key := filename + "-" + include + "-" + name
	if d := this_.structCache_.Get(key); d != nil {
		res = d.(*Struct)
		return
	}

	res = &Struct{}
	res.Name = name
	structFilename := filename
	if include != "" {
		structFilename = this_.GetIncludePath(filename, include)
	}
	//fmt.Println("GetStruct structFilename:", structFilename, ",name:", name)
	structNode := this_.GetStruct(structFilename, name)
	//fmt.Println("GetStructByName filename:", filename, ",include:", include, ",name:", name, ",structNode:", toJSON(structNode))
	if structNode != nil {
		for _, fieldNode := range structNode.Fields {
			res.Fields = append(res.Fields, this_.GetFieldByNode(structFilename, fieldNode, structCache))
		}
	}

	this_.structCache_.Set(key, res)
	return
}

func (this_ *Workspace) GetStructDemoData(filename string, include string, name string) (res interface{}) {

	key := filename + "-" + include + "-" + name
	r := this_.structCache_.Get(key)
	s := r.(*Struct)

	structFilename := filename
	if include != "" {
		structFilename = this_.GetIncludePath(filename, include)
	}

	data := map[string]interface{}{}
	res = data
	for _, fieldNode := range s.Fields {
		data[fieldNode.Name] = this_.GetFieldDemoData(structFilename, fieldNode)
	}
	return
}
