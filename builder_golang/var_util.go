package builder_golang

import (
	"maker/parser_tm"
)

func (this_ *Context) initUtilVar() {
	utilSpace := this_.NewVarFuncSpace()
	utilSpace.PackImpl = "github.com/team-ide/go-tool/util"
	utilSpace.PackAsName = "util"
	this_.AddVar("util", utilSpace)


	utilSpace.AddVar("AesEncryptCBCByKey", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "AesEncryptCBCByKey",
		Args: []*parser_tm.FuncArgNode{
			{Name: "origData", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})
	this_.AddVar("AesEncryptCBCByKey", utilSpace.getVar("AesEncryptCBCByKey"))

	utilSpace.AddVar("AesDecryptCBCByKey", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "AesDecryptCBCByKey",
		Args: []*parser_tm.FuncArgNode{
			{Name: "encrypt", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})
	this_.AddVar("AesDecryptCBCByKey", utilSpace.getVar("AesDecryptCBCByKey"))

	utilSpace.AddVar("AesEncryptECBByKey", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "AesEncryptECBByKey",
		Args: []*parser_tm.FuncArgNode{
			{Name: "origData", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})
	this_.AddVar("AesEncryptECBByKey", utilSpace.getVar("AesEncryptECBByKey"))

	utilSpace.AddVar("AesDecryptECBByKey", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "AesDecryptECBByKey",
		Args: []*parser_tm.FuncArgNode{
			{Name: "encrypt", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})
	this_.AddVar("AesDecryptECBByKey", utilSpace.getVar("AesDecryptECBByKey"))

	utilSpace.AddVar("IsEmpty", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "IsEmpty",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
	})
	this_.AddVar("IsEmpty", utilSpace.getVar("IsEmpty"))

	utilSpace.AddVar("IsNotEmpty", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "IsNotEmpty",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
	})
	this_.AddVar("IsNotEmpty", utilSpace.getVar("IsNotEmpty"))

	utilSpace.AddVar("IsNull", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "IsNull",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
	})
	this_.AddVar("IsNull", utilSpace.getVar("IsNull"))

	utilSpace.AddVar("IsNotNull", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "IsNotNull",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
	})
	this_.AddVar("IsNotNull", utilSpace.getVar("IsNotNull"))

	utilSpace.AddVar("IsTrue", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "IsTrue",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
	})
	this_.AddVar("IsTrue", utilSpace.getVar("IsTrue"))

	utilSpace.AddVar("IsFalse", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "IsFalse",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
	})
	this_.AddVar("IsFalse", utilSpace.getVar("IsFalse"))

	utilSpace.AddVar("IntIndexOf", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "IntIndexOf",
		Args: []*parser_tm.FuncArgNode{
			{Name: "array", Type: parser_tm.NewBindingTypeName("[]int")},
			{Name: "v", Type: parser_tm.NewBindingTypeName("int")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "index", Type: parser_tm.NewBindingTypeName("int")},
	})
	this_.AddVar("IntIndexOf", utilSpace.getVar("IntIndexOf"))

	utilSpace.AddVar("Int64IndexOf", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "Int64IndexOf",
		Args: []*parser_tm.FuncArgNode{
			{Name: "array", Type: parser_tm.NewBindingTypeName("[]int64")},
			{Name: "v", Type: parser_tm.NewBindingTypeName("int64")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "index", Type: parser_tm.NewBindingTypeName("int")},
	})
	this_.AddVar("Int64IndexOf", utilSpace.getVar("Int64IndexOf"))

	utilSpace.AddVar("StringIndexOf", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "StringIndexOf",
		Args: []*parser_tm.FuncArgNode{
			{Name: "array", Type: parser_tm.NewBindingTypeName("[]string")},
			{Name: "v", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "index", Type: parser_tm.NewBindingTypeName("int")},
	})
	this_.AddVar("StringIndexOf", utilSpace.getVar("StringIndexOf"))

	utilSpace.AddVar("ArrayIndexOf", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "ArrayIndexOf",
		Args: []*parser_tm.FuncArgNode{
			{Name: "array", Type: parser_tm.NewBindingTypeName("any")},
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "index", Type: parser_tm.NewBindingTypeName("int")},
	})
	this_.AddVar("ArrayIndexOf", utilSpace.getVar("ArrayIndexOf"))

	utilSpace.AddVar("GetTempDir", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "GetTempDir",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "dir", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})
	this_.AddVar("GetTempDir", utilSpace.getVar("GetTempDir"))

	utilSpace.AddVar("NewWaitGroup", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "NewWaitGroup",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("*sync.WaitGroup")},
	})
	this_.AddVar("NewWaitGroup", utilSpace.getVar("NewWaitGroup"))

	utilSpace.AddVar("NewLocker", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "NewLocker",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("sync.Locker")},
	})
	this_.AddVar("NewLocker", utilSpace.getVar("NewLocker"))

	utilSpace.AddVar("GetRootDir", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "GetRootDir",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("GetRootDir", utilSpace.getVar("GetRootDir"))

	utilSpace.AddVar("FormatPath", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "FormatPath",
		Args: []*parser_tm.FuncArgNode{
			{Name: "path", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("FormatPath", utilSpace.getVar("FormatPath"))

	utilSpace.AddVar("GetAbsolutePath", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "GetAbsolutePath",
		Args: []*parser_tm.FuncArgNode{
			{Name: "path", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "absolutePath", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("GetAbsolutePath", utilSpace.getVar("GetAbsolutePath"))

	utilSpace.AddVar("PathExists", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "PathExists",
		Args: []*parser_tm.FuncArgNode{
			{Name: "path", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
		HasError: true,
	})
	this_.AddVar("PathExists", utilSpace.getVar("PathExists"))

	utilSpace.AddVar("LoadDirFiles", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "LoadDirFiles",
		Args: []*parser_tm.FuncArgNode{
			{Name: "dir", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "fileMap", Type: parser_tm.NewBindingTypeName("map[string][]byte")},
		HasError: true,
	})
	this_.AddVar("LoadDirFiles", utilSpace.getVar("LoadDirFiles"))

	utilSpace.AddVar("LoadDirFilenames", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "LoadDirFilenames",
		Args: []*parser_tm.FuncArgNode{
			{Name: "dir", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "filenames", Type: parser_tm.NewBindingTypeList("string")},
		HasError: true,
	})
	this_.AddVar("LoadDirFilenames", utilSpace.getVar("LoadDirFilenames"))

	utilSpace.AddVar("ReadFile", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "ReadFile",
		Args: []*parser_tm.FuncArgNode{
			{Name: "filename", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "bs", Type: parser_tm.NewBindingTypeList("byte")},
		HasError: true,
	})
	this_.AddVar("ReadFile", utilSpace.getVar("ReadFile"))

	utilSpace.AddVar("ReadFileString", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "ReadFileString",
		Args: []*parser_tm.FuncArgNode{
			{Name: "filename", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})
	this_.AddVar("ReadFileString", utilSpace.getVar("ReadFileString"))

	utilSpace.AddVar("StringToBytes", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "StringToBytes",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeList("byte")},
	})
	this_.AddVar("StringToBytes", utilSpace.getVar("StringToBytes"))

	utilSpace.AddVar("WriteFile", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "WriteFile",
		Args: []*parser_tm.FuncArgNode{
			{Name: "filename", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "bs", Type: parser_tm.NewBindingTypeName("[]byte")},
		},
		HasError: true,
	})
	this_.AddVar("WriteFile", utilSpace.getVar("WriteFile"))

	utilSpace.AddVar("WriteFileString", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "WriteFileString",
		Args: []*parser_tm.FuncArgNode{
			{Name: "filename", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		HasError: true,
	})
	this_.AddVar("WriteFileString", utilSpace.getVar("WriteFileString"))

	utilSpace.AddVar("ReadLine", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "ReadLine",
		Args: []*parser_tm.FuncArgNode{
			{Name: "filename", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "lines", Type: parser_tm.NewBindingTypeList("string")},
		HasError: true,
	})
	this_.AddVar("ReadLine", utilSpace.getVar("ReadLine"))

	utilSpace.AddVar("IsSubPath", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "IsSubPath",
		Args: []*parser_tm.FuncArgNode{
			{Name: "parent", Type: parser_tm.NewBindingTypeName("any")},
			{Name: "child", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "isSub", Type: parser_tm.NewBindingTypeName("bool")},
		HasError: true,
	})
	this_.AddVar("IsSubPath", utilSpace.getVar("IsSubPath"))

	utilSpace.AddVar("LoadDirInfo", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "LoadDirInfo",
		Args: []*parser_tm.FuncArgNode{
			{Name: "dir", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "loadSubDir", Type: parser_tm.NewBindingTypeName("bool")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "dirInfo", Type: parser_tm.NewBindingTypeName("*DirInfo")},
		HasError: true,
	})
	this_.AddVar("LoadDirInfo", utilSpace.getVar("LoadDirInfo"))

	utilSpace.AddVar("GetFileType", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "GetFileType",
		Args: []*parser_tm.FuncArgNode{
			{Name: "fSrc", Type: parser_tm.NewBindingTypeName("[]byte")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("GetFileType", utilSpace.getVar("GetFileType"))

	utilSpace.AddVar("NextId", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "NextId",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})
	this_.AddVar("NextId", utilSpace.getVar("NextId"))

	utilSpace.AddVar("NewIdWorker", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "NewIdWorker",
		Args: []*parser_tm.FuncArgNode{
			{Name: "workerId", Type: parser_tm.NewBindingTypeName("int64")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("*IdWorker")},
		HasError: true,
	})
	this_.AddVar("NewIdWorker", utilSpace.getVar("NewIdWorker"))

	utilSpace.AddVar("GetIpFromAddr", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "GetIpFromAddr",
		Args: []*parser_tm.FuncArgNode{
			{Name: "addr", Type: parser_tm.NewBindingTypeName("net.Addr")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("net.IP")},
	})
	this_.AddVar("GetIpFromAddr", utilSpace.getVar("GetIpFromAddr"))

	utilSpace.AddVar("GetLocalIPList", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "GetLocalIPList",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "ipList", Type: parser_tm.NewBindingTypeList("net.IP")},
	})
	this_.AddVar("GetLocalIPList", utilSpace.getVar("GetLocalIPList"))

	utilSpace.AddVar("ObjToJson", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "ObjToJson",
		Args: []*parser_tm.FuncArgNode{
			{Name: "obj", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})
	this_.AddVar("ObjToJson", utilSpace.getVar("ObjToJson"))

	utilSpace.AddVar("JsonToMap", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "JsonToMap",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("map[string]interface{}")},
		HasError: true,
	})
	this_.AddVar("JsonToMap", utilSpace.getVar("JsonToMap"))

	utilSpace.AddVar("JsonToObj", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "JsonToObj",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "obj", Type: parser_tm.NewBindingTypeName("any")},
		},
		HasError: true,
	})
	this_.AddVar("JsonToObj", utilSpace.getVar("JsonToObj"))

	utilSpace.AddVar("GetLock", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "GetLock",
		Args: []*parser_tm.FuncArgNode{
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "lock", Type: parser_tm.NewBindingTypeName("sync.Locker")},
	})
	this_.AddVar("GetLock", utilSpace.getVar("GetLock"))

	utilSpace.AddVar("LockByKey", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "LockByKey",
		Args: []*parser_tm.FuncArgNode{
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("")},
	})
	this_.AddVar("LockByKey", utilSpace.getVar("LockByKey"))

	utilSpace.AddVar("UnlockByKey", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "UnlockByKey",
		Args: []*parser_tm.FuncArgNode{
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("")},
	})
	this_.AddVar("UnlockByKey", utilSpace.getVar("UnlockByKey"))

	utilSpace.AddVar("GetLogger", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "GetLogger",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("*zap.Logger")},
	})
	this_.AddVar("GetLogger", utilSpace.getVar("GetLogger"))

	utilSpace.AddVar("NewLoggerByCallerSkip", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "NewLoggerByCallerSkip",
		Args: []*parser_tm.FuncArgNode{
			{Name: "skip", Type: parser_tm.NewBindingTypeName("int")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("*zap.Logger")},
	})
	this_.AddVar("NewLoggerByCallerSkip", utilSpace.getVar("NewLoggerByCallerSkip"))

	utilSpace.AddVar("GetMD5", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "GetMD5",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("GetMD5", utilSpace.getVar("GetMD5"))

	utilSpace.AddVar("RandomInt", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "RandomInt",
		Args: []*parser_tm.FuncArgNode{
			{Name: "min", Type: parser_tm.NewBindingTypeName("int")},
			{Name: "max", Type: parser_tm.NewBindingTypeName("int")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int")},
	})
	this_.AddVar("RandomInt", utilSpace.getVar("RandomInt"))

	utilSpace.AddVar("RandomInt64", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "RandomInt64",
		Args: []*parser_tm.FuncArgNode{
			{Name: "min", Type: parser_tm.NewBindingTypeName("int64")},
			{Name: "max", Type: parser_tm.NewBindingTypeName("int64")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})
	this_.AddVar("RandomInt64", utilSpace.getVar("RandomInt64"))

	utilSpace.AddVar("StringToInt", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "StringToInt",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int")},
	})
	this_.AddVar("StringToInt", utilSpace.getVar("StringToInt"))

	utilSpace.AddVar("StringToInt64", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "StringToInt64",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})
	this_.AddVar("StringToInt64", utilSpace.getVar("StringToInt64"))

	utilSpace.AddVar("StringToUint64", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "StringToUint64",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("uint64")},
	})
	this_.AddVar("StringToUint64", utilSpace.getVar("StringToUint64"))

	utilSpace.AddVar("StringToFloat64", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "StringToFloat64",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("float64")},
	})
	this_.AddVar("StringToFloat64", utilSpace.getVar("StringToFloat64"))

	utilSpace.AddVar("SumToString", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "SumToString",
		Args: []*parser_tm.FuncArgNode{
			{Name: "nums", Type: parser_tm.NewBindingTypeName("...interface{}")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("SumToString", utilSpace.getVar("SumToString"))

	utilSpace.AddVar("ValueToInt64", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "ValueToInt64",
		Args: []*parser_tm.FuncArgNode{
			{Name: "value", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
		HasError: true,
	})
	this_.AddVar("ValueToInt64", utilSpace.getVar("ValueToInt64"))

	utilSpace.AddVar("ValueToUint64", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "ValueToUint64",
		Args: []*parser_tm.FuncArgNode{
			{Name: "value", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("uint64")},
		HasError: true,
	})
	this_.AddVar("ValueToUint64", utilSpace.getVar("ValueToUint64"))

	utilSpace.AddVar("ValueToFloat64", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "ValueToFloat64",
		Args: []*parser_tm.FuncArgNode{
			{Name: "value", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("float64")},
		HasError: true,
	})
	this_.AddVar("ValueToFloat64", utilSpace.getVar("ValueToFloat64"))

	utilSpace.AddVar("RsaEncryptByKey", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "RsaEncryptByKey",
		Args: []*parser_tm.FuncArgNode{
			{Name: "origData", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "publicKey", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})
	this_.AddVar("RsaEncryptByKey", utilSpace.getVar("RsaEncryptByKey"))

	utilSpace.AddVar("RsaDecryptByKey", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "RsaDecryptByKey",
		Args: []*parser_tm.FuncArgNode{
			{Name: "decrypt", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "privateKey", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})
	this_.AddVar("RsaDecryptByKey", utilSpace.getVar("RsaDecryptByKey"))

	utilSpace.AddVar("FirstToUpper", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "FirstToUpper",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("FirstToUpper", utilSpace.getVar("FirstToUpper"))

	utilSpace.AddVar("FirstToLower", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "FirstToLower",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("FirstToLower", utilSpace.getVar("FirstToLower"))

	utilSpace.AddVar("Marshal", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "Marshal",
		Args: []*parser_tm.FuncArgNode{
			{Name: "name", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("Marshal", utilSpace.getVar("Marshal"))

	utilSpace.AddVar("Hump", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "Hump",
		Args: []*parser_tm.FuncArgNode{
			{Name: "name", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("Hump", utilSpace.getVar("Hump"))

	utilSpace.AddVar("GetStringValue", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "GetStringValue",
		Args: []*parser_tm.FuncArgNode{
			{Name: "value", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "valueString", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("GetStringValue", utilSpace.getVar("GetStringValue"))

	utilSpace.AddVar("RandomString", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "RandomString",
		Args: []*parser_tm.FuncArgNode{
			{Name: "minLen", Type: parser_tm.NewBindingTypeName("int")},
			{Name: "maxLen", Type: parser_tm.NewBindingTypeName("int")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("RandomString", utilSpace.getVar("RandomString"))

	utilSpace.AddVar("RandomUserName", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "RandomUserName",
		Args: []*parser_tm.FuncArgNode{
			{Name: "size", Type: parser_tm.NewBindingTypeName("int")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("RandomUserName", utilSpace.getVar("RandomUserName"))

	utilSpace.AddVar("StrPadLeft", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "StrPadLeft",
		Args: []*parser_tm.FuncArgNode{
			{Name: "input", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "padLength", Type: parser_tm.NewBindingTypeName("int")},
			{Name: "padString", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("StrPadLeft", utilSpace.getVar("StrPadLeft"))

	utilSpace.AddVar("StrPadRight", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "StrPadRight",
		Args: []*parser_tm.FuncArgNode{
			{Name: "input", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "padLength", Type: parser_tm.NewBindingTypeName("int")},
			{Name: "padString", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("StrPadRight", utilSpace.getVar("StrPadRight"))

	utilSpace.AddVar("TrimSpace", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "TrimSpace",
		Args: []*parser_tm.FuncArgNode{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("TrimSpace", utilSpace.getVar("TrimSpace"))

	utilSpace.AddVar("TrimPrefix", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "TrimPrefix",
		Args: []*parser_tm.FuncArgNode{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("TrimPrefix", utilSpace.getVar("TrimPrefix"))

	utilSpace.AddVar("HasPrefix", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "HasPrefix",
		Args: []*parser_tm.FuncArgNode{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
	})
	this_.AddVar("HasPrefix", utilSpace.getVar("HasPrefix"))

	utilSpace.AddVar("TrimSuffix", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "TrimSuffix",
		Args: []*parser_tm.FuncArgNode{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("TrimSuffix", utilSpace.getVar("TrimSuffix"))

	utilSpace.AddVar("HasSuffix", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "HasSuffix",
		Args: []*parser_tm.FuncArgNode{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
	})
	this_.AddVar("HasSuffix", utilSpace.getVar("HasSuffix"))

	utilSpace.AddVar("TrimLeft", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "TrimLeft",
		Args: []*parser_tm.FuncArgNode{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("TrimLeft", utilSpace.getVar("TrimLeft"))

	utilSpace.AddVar("TrimRight", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "TrimRight",
		Args: []*parser_tm.FuncArgNode{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("TrimRight", utilSpace.getVar("TrimRight"))

	utilSpace.AddVar("StringJoin", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "StringJoin",
		Args: []*parser_tm.FuncArgNode{
			{Name: "es", Type: parser_tm.NewBindingTypeName("[]string")},
			{Name: "sep", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("StringJoin", utilSpace.getVar("StringJoin"))

	utilSpace.AddVar("AnyJoin", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "AnyJoin",
		Args: []*parser_tm.FuncArgNode{
			{Name: "sep", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "es", Type: parser_tm.NewBindingTypeName("...any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("AnyJoin", utilSpace.getVar("AnyJoin"))

	utilSpace.AddVar("IntJoin", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "IntJoin",
		Args: []*parser_tm.FuncArgNode{
			{Name: "es", Type: parser_tm.NewBindingTypeName("[]int")},
			{Name: "sep", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("IntJoin", utilSpace.getVar("IntJoin"))

	utilSpace.AddVar("Int64Join", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "Int64Join",
		Args: []*parser_tm.FuncArgNode{
			{Name: "es", Type: parser_tm.NewBindingTypeName("[]int64")},
			{Name: "sep", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("Int64Join", utilSpace.getVar("Int64Join"))

	utilSpace.AddVar("GenStringJoin", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "GenStringJoin",
		Args: []*parser_tm.FuncArgNode{
			{Name: "len", Type: parser_tm.NewBindingTypeName("int")},
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "sep", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("GenStringJoin", utilSpace.getVar("GenStringJoin"))

	utilSpace.AddVar("GetNow", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "GetNow",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("time.Time")},
	})
	this_.AddVar("GetNow", utilSpace.getVar("GetNow"))

	utilSpace.AddVar("GetNowNano", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "GetNowNano",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})
	this_.AddVar("GetNowNano", utilSpace.getVar("GetNowNano"))

	utilSpace.AddVar("GetNowMilli", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "GetNowMilli",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})
	this_.AddVar("GetNowMilli", utilSpace.getVar("GetNowMilli"))

	utilSpace.AddVar("GetNowSecond", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "GetNowSecond",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})
	this_.AddVar("GetNowSecond", utilSpace.getVar("GetNowSecond"))

	utilSpace.AddVar("GetNanoByTime", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "GetNanoByTime",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("time.Time")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})
	this_.AddVar("GetNanoByTime", utilSpace.getVar("GetNanoByTime"))

	utilSpace.AddVar("GetMilliByTime", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "GetMilliByTime",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("time.Time")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})
	this_.AddVar("GetMilliByTime", utilSpace.getVar("GetMilliByTime"))

	utilSpace.AddVar("GetSecondByTime", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "GetSecondByTime",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("time.Time")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})
	this_.AddVar("GetSecondByTime", utilSpace.getVar("GetSecondByTime"))

	utilSpace.AddVar("GetNowFormat", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "GetNowFormat",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("GetNowFormat", utilSpace.getVar("GetNowFormat"))

	utilSpace.AddVar("GetFormatByTime", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "GetFormatByTime",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("time.Time")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("GetFormatByTime", utilSpace.getVar("GetFormatByTime"))

	utilSpace.AddVar("TimeFormat", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "TimeFormat",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("time.Time")},
			{Name: "layout", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("TimeFormat", utilSpace.getVar("TimeFormat"))

	utilSpace.AddVar("MilliToTimeText", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "MilliToTimeText",
		Args: []*parser_tm.FuncArgNode{
			{Name: "milli", Type: parser_tm.NewBindingTypeName("int64")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "v", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("MilliToTimeText", utilSpace.getVar("MilliToTimeText"))

	utilSpace.AddVar("GetUUID", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "GetUUID",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	this_.AddVar("GetUUID", utilSpace.getVar("GetUUID"))

	utilSpace.AddVar("GzipBytes", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "GzipBytes",
		Args: []*parser_tm.FuncArgNode{
			{Name: "data", Type: parser_tm.NewBindingTypeName("[]byte")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeList("byte")},
		HasError: true,
	})
	this_.AddVar("GzipBytes", utilSpace.getVar("GzipBytes"))

	utilSpace.AddVar("UnGzipBytes", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "UnGzipBytes",
		Args: []*parser_tm.FuncArgNode{
			{Name: "data", Type: parser_tm.NewBindingTypeName("[]byte")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeList("byte")},
		HasError: true,
	})
	this_.AddVar("UnGzipBytes", utilSpace.getVar("UnGzipBytes"))

	utilSpace.AddVar("Zip", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "Zip",
		Args: []*parser_tm.FuncArgNode{
			{Name: "srcFile", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "destZip", Type: parser_tm.NewBindingTypeName("string")},
		},
		HasError: true,
	})
	this_.AddVar("Zip", utilSpace.getVar("Zip"))

	utilSpace.AddVar("UnZip", &VarFunc{
		VarFuncSpace: utilSpace,
		ScriptName:  "UnZip",
		Args: []*parser_tm.FuncArgNode{
			{Name: "zipFile", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "destDir", Type: parser_tm.NewBindingTypeName("string")},
		},
		HasError: true,
	})
	this_.AddVar("UnZip", utilSpace.getVar("UnZip"))

	return
}