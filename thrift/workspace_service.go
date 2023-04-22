package thrift

import (
	"context"
	"errors"
	"github.com/team-ide/go-interpreter/thrift"
	"github.com/team-ide/go-tool/util"
)

func (this_ *Workspace) InvokeByServerAddress(serverAddress string, filename string, serviceName string, methodName string, args ...interface{}) (res interface{}, err error) {

	param, err := this_.GetMethodParam(filename, serviceName, methodName, args...)
	if err != nil {
		return
	}

	client, err := NewServiceClientByAddress(serverAddress)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = client.t.Close()
	}()

	res, err = client.Send(context.Background(), param)
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
	for _, one := range methodNode.Params {
		param.ArgFields = append(param.ArgFields, this_.GetFieldByNode(filename, one))
	}
	param.ResultType = this_.GetFieldTypeByNode(filename, methodNode.Return)

	return
}

func (this_ *Workspace) GetFieldByNode(filename string, fieldNode *thrift.FieldNode) (field *Field) {

	field = &Field{
		Name: fieldNode.Name,
		Num:  fieldNode.Num,
	}
	field.Type = this_.GetFieldTypeByNode(filename, fieldNode.Type)
	return
}

func (this_ *Workspace) GetFieldTypeByNode(filename string, fieldNode *thrift.FieldType) (fieldType *FieldType) {
	fieldType = &FieldType{
		TypeId: fieldNode.TypeId,
	}
	if fieldNode.StructName != "" {
		fieldType.Struct = this_.GetStructByName(filename, fieldNode.StructInclude, fieldNode.StructName)
	}
	if fieldNode.MapKeyType != nil {
		fieldType.MapKeyType = this_.GetFieldTypeByNode(filename, fieldNode.MapKeyType)
	}
	if fieldNode.MapValueType != nil {
		fieldType.MapValueType = this_.GetFieldTypeByNode(filename, fieldNode.MapValueType)
	}
	if fieldNode.ListType != nil {
		fieldType.ListType = this_.GetFieldTypeByNode(filename, fieldNode.ListType)
	}
	if fieldNode.SetType != nil {
		fieldType.SetType = this_.GetFieldTypeByNode(filename, fieldNode.SetType)
	}
	return
}

func (this_ *Workspace) GetStructByName(filename string, include string, name string) (res *Struct) {
	key := filename + "-" + include + "-" + name
	if res := this_.structCache_.Get(key); res != nil {
		return res.(*Struct)
	}
	res = &Struct{}
	res.Name = name
	structFilename := filename
	if include != "" {
		structFilename = this_.GetIncludePath(filename, include)
	}
	structNode := this_.GetStruct(structFilename, name)
	if structNode != nil {
		for _, fieldNode := range structNode.Fields {
			res.Fields = append(res.Fields, this_.GetFieldByNode(structFilename, fieldNode))
		}
	}

	this_.structCache_.Set(key, res)
	return
}
