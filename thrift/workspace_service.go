package thrift

import (
	"context"
	"errors"
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
		_ = client.t.Close()
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
	var structCache map[string]*Struct
	param.ArgFields, structCache, err = this_.GetMethodArgFields(filename, serviceName, methodName)
	if err != nil {
		return
	}
	param.ResultType = this_.GetFieldTypeByNode(filename, methodNode.Return, structCache)

	return
}

func (this_ *Workspace) GetMethodArgFields(filename string, serviceName string, methodName string) (argFields []*Field, structCache map[string]*Struct, err error) {
	structCache = map[string]*Struct{}

	filename = util.FormatPath(filename)

	methodNode := this_.GetServiceMethod(filename, serviceName, methodName)
	if methodNode == nil {
		err = errors.New("service method node [" + filename + "][" + serviceName + "][" + methodName + "] not found")
		return
	}
	for _, one := range methodNode.Params {
		argFields = append(argFields, this_.GetFieldByNode(filename, one, structCache))
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
