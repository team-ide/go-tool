package builder_golang

import (
	"maker/parser_tm"
)

func (this_ *Context) initUtilVar() {
	utilSpace := this_.NewVarSpace()
	utilSpace.PackImpl = "github.com/team-ide/go-tool/util"
	utilSpace.PackAsName = "util"
	this_.AddVar("util", utilSpace)


	this_.AddVar("AesEncryptCBCByKey", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "AesEncryptCBCByKey",
		Args: []*parser_tm.FuncArgNode{
			{Name: "origData", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})
	utilSpace.AddVar("AesEncryptCBCByKey", &VarFunc{
		ScriptName:  "AesEncryptCBCByKey",
		Args: []*parser_tm.FuncArgNode{
			{Name: "origData", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})

	this_.AddVar("AesDecryptCBCByKey", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "AesDecryptCBCByKey",
		Args: []*parser_tm.FuncArgNode{
			{Name: "encrypt", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})
	utilSpace.AddVar("AesDecryptCBCByKey", &VarFunc{
		ScriptName:  "AesDecryptCBCByKey",
		Args: []*parser_tm.FuncArgNode{
			{Name: "encrypt", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})

	this_.AddVar("AesEncryptECBByKey", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "AesEncryptECBByKey",
		Args: []*parser_tm.FuncArgNode{
			{Name: "origData", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})
	utilSpace.AddVar("AesEncryptECBByKey", &VarFunc{
		ScriptName:  "AesEncryptECBByKey",
		Args: []*parser_tm.FuncArgNode{
			{Name: "origData", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})

	this_.AddVar("AesDecryptECBByKey", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "AesDecryptECBByKey",
		Args: []*parser_tm.FuncArgNode{
			{Name: "encrypt", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})
	utilSpace.AddVar("AesDecryptECBByKey", &VarFunc{
		ScriptName:  "AesDecryptECBByKey",
		Args: []*parser_tm.FuncArgNode{
			{Name: "encrypt", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})

	this_.AddVar("IsEmpty", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "IsEmpty",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
	})
	utilSpace.AddVar("IsEmpty", &VarFunc{
		ScriptName:  "IsEmpty",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
	})

	this_.AddVar("IsNotEmpty", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "IsNotEmpty",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
	})
	utilSpace.AddVar("IsNotEmpty", &VarFunc{
		ScriptName:  "IsNotEmpty",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
	})

	this_.AddVar("IsNull", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "IsNull",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
	})
	utilSpace.AddVar("IsNull", &VarFunc{
		ScriptName:  "IsNull",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
	})

	this_.AddVar("IsNotNull", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "IsNotNull",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
	})
	utilSpace.AddVar("IsNotNull", &VarFunc{
		ScriptName:  "IsNotNull",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
	})

	this_.AddVar("IsTrue", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "IsTrue",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
	})
	utilSpace.AddVar("IsTrue", &VarFunc{
		ScriptName:  "IsTrue",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
	})

	this_.AddVar("IsFalse", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "IsFalse",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
	})
	utilSpace.AddVar("IsFalse", &VarFunc{
		ScriptName:  "IsFalse",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
	})

	this_.AddVar("IntIndexOf", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "IntIndexOf",
		Args: []*parser_tm.FuncArgNode{
			{Name: "array", Type: parser_tm.NewBindingTypeName("[]int")},
			{Name: "v", Type: parser_tm.NewBindingTypeName("int")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "index", Type: parser_tm.NewBindingTypeName("int")},
	})
	utilSpace.AddVar("IntIndexOf", &VarFunc{
		ScriptName:  "IntIndexOf",
		Args: []*parser_tm.FuncArgNode{
			{Name: "array", Type: parser_tm.NewBindingTypeName("[]int")},
			{Name: "v", Type: parser_tm.NewBindingTypeName("int")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "index", Type: parser_tm.NewBindingTypeName("int")},
	})

	this_.AddVar("Int64IndexOf", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "Int64IndexOf",
		Args: []*parser_tm.FuncArgNode{
			{Name: "array", Type: parser_tm.NewBindingTypeName("[]int64")},
			{Name: "v", Type: parser_tm.NewBindingTypeName("int64")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "index", Type: parser_tm.NewBindingTypeName("int")},
	})
	utilSpace.AddVar("Int64IndexOf", &VarFunc{
		ScriptName:  "Int64IndexOf",
		Args: []*parser_tm.FuncArgNode{
			{Name: "array", Type: parser_tm.NewBindingTypeName("[]int64")},
			{Name: "v", Type: parser_tm.NewBindingTypeName("int64")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "index", Type: parser_tm.NewBindingTypeName("int")},
	})

	this_.AddVar("StringIndexOf", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "StringIndexOf",
		Args: []*parser_tm.FuncArgNode{
			{Name: "array", Type: parser_tm.NewBindingTypeName("[]string")},
			{Name: "v", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "index", Type: parser_tm.NewBindingTypeName("int")},
	})
	utilSpace.AddVar("StringIndexOf", &VarFunc{
		ScriptName:  "StringIndexOf",
		Args: []*parser_tm.FuncArgNode{
			{Name: "array", Type: parser_tm.NewBindingTypeName("[]string")},
			{Name: "v", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "index", Type: parser_tm.NewBindingTypeName("int")},
	})

	this_.AddVar("ArrayIndexOf", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "ArrayIndexOf",
		Args: []*parser_tm.FuncArgNode{
			{Name: "array", Type: parser_tm.NewBindingTypeName("any")},
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "index", Type: parser_tm.NewBindingTypeName("int")},
	})
	utilSpace.AddVar("ArrayIndexOf", &VarFunc{
		ScriptName:  "ArrayIndexOf",
		Args: []*parser_tm.FuncArgNode{
			{Name: "array", Type: parser_tm.NewBindingTypeName("any")},
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "index", Type: parser_tm.NewBindingTypeName("int")},
	})

	this_.AddVar("GetTempDir", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetTempDir",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "dir", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})
	utilSpace.AddVar("GetTempDir", &VarFunc{
		ScriptName:  "GetTempDir",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "dir", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})

	this_.AddVar("NewWaitGroup", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "NewWaitGroup",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("*sync.WaitGroup")},
	})
	utilSpace.AddVar("NewWaitGroup", &VarFunc{
		ScriptName:  "NewWaitGroup",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("*sync.WaitGroup")},
	})

	this_.AddVar("NewLocker", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "NewLocker",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("sync.Locker")},
	})
	utilSpace.AddVar("NewLocker", &VarFunc{
		ScriptName:  "NewLocker",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("sync.Locker")},
	})

	this_.AddVar("GetRootDir", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetRootDir",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("GetRootDir", &VarFunc{
		ScriptName:  "GetRootDir",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("FormatPath", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "FormatPath",
		Args: []*parser_tm.FuncArgNode{
			{Name: "path", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("FormatPath", &VarFunc{
		ScriptName:  "FormatPath",
		Args: []*parser_tm.FuncArgNode{
			{Name: "path", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("GetAbsolutePath", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetAbsolutePath",
		Args: []*parser_tm.FuncArgNode{
			{Name: "path", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "absolutePath", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("GetAbsolutePath", &VarFunc{
		ScriptName:  "GetAbsolutePath",
		Args: []*parser_tm.FuncArgNode{
			{Name: "path", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "absolutePath", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("PathExists", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "PathExists",
		Args: []*parser_tm.FuncArgNode{
			{Name: "path", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
		HasError: true,
	})
	utilSpace.AddVar("PathExists", &VarFunc{
		ScriptName:  "PathExists",
		Args: []*parser_tm.FuncArgNode{
			{Name: "path", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
		HasError: true,
	})

	this_.AddVar("LoadDirFiles", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "LoadDirFiles",
		Args: []*parser_tm.FuncArgNode{
			{Name: "dir", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "fileMap", Type: parser_tm.NewBindingTypeName("map[string][]byte")},
		HasError: true,
	})
	utilSpace.AddVar("LoadDirFiles", &VarFunc{
		ScriptName:  "LoadDirFiles",
		Args: []*parser_tm.FuncArgNode{
			{Name: "dir", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "fileMap", Type: parser_tm.NewBindingTypeName("map[string][]byte")},
		HasError: true,
	})

	this_.AddVar("LoadDirFilenames", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "LoadDirFilenames",
		Args: []*parser_tm.FuncArgNode{
			{Name: "dir", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "filenames", Type: parser_tm.NewBindingTypeList("string")},
		HasError: true,
	})
	utilSpace.AddVar("LoadDirFilenames", &VarFunc{
		ScriptName:  "LoadDirFilenames",
		Args: []*parser_tm.FuncArgNode{
			{Name: "dir", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "filenames", Type: parser_tm.NewBindingTypeList("string")},
		HasError: true,
	})

	this_.AddVar("ReadFile", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "ReadFile",
		Args: []*parser_tm.FuncArgNode{
			{Name: "filename", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "bs", Type: parser_tm.NewBindingTypeList("byte")},
		HasError: true,
	})
	utilSpace.AddVar("ReadFile", &VarFunc{
		ScriptName:  "ReadFile",
		Args: []*parser_tm.FuncArgNode{
			{Name: "filename", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "bs", Type: parser_tm.NewBindingTypeList("byte")},
		HasError: true,
	})

	this_.AddVar("ReadFileString", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "ReadFileString",
		Args: []*parser_tm.FuncArgNode{
			{Name: "filename", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})
	utilSpace.AddVar("ReadFileString", &VarFunc{
		ScriptName:  "ReadFileString",
		Args: []*parser_tm.FuncArgNode{
			{Name: "filename", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})

	this_.AddVar("StringToBytes", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "StringToBytes",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeList("byte")},
	})
	utilSpace.AddVar("StringToBytes", &VarFunc{
		ScriptName:  "StringToBytes",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeList("byte")},
	})

	this_.AddVar("WriteFile", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "WriteFile",
		Args: []*parser_tm.FuncArgNode{
			{Name: "filename", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "bs", Type: parser_tm.NewBindingTypeName("[]byte")},
		},
		HasError: true,
	})
	utilSpace.AddVar("WriteFile", &VarFunc{
		ScriptName:  "WriteFile",
		Args: []*parser_tm.FuncArgNode{
			{Name: "filename", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "bs", Type: parser_tm.NewBindingTypeName("[]byte")},
		},
		HasError: true,
	})

	this_.AddVar("WriteFileString", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "WriteFileString",
		Args: []*parser_tm.FuncArgNode{
			{Name: "filename", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		HasError: true,
	})
	utilSpace.AddVar("WriteFileString", &VarFunc{
		ScriptName:  "WriteFileString",
		Args: []*parser_tm.FuncArgNode{
			{Name: "filename", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		HasError: true,
	})

	this_.AddVar("ReadLine", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "ReadLine",
		Args: []*parser_tm.FuncArgNode{
			{Name: "filename", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "lines", Type: parser_tm.NewBindingTypeList("string")},
		HasError: true,
	})
	utilSpace.AddVar("ReadLine", &VarFunc{
		ScriptName:  "ReadLine",
		Args: []*parser_tm.FuncArgNode{
			{Name: "filename", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "lines", Type: parser_tm.NewBindingTypeList("string")},
		HasError: true,
	})

	this_.AddVar("IsSubPath", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "IsSubPath",
		Args: []*parser_tm.FuncArgNode{
			{Name: "parent", Type: parser_tm.NewBindingTypeName("any")},
			{Name: "child", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "isSub", Type: parser_tm.NewBindingTypeName("bool")},
		HasError: true,
	})
	utilSpace.AddVar("IsSubPath", &VarFunc{
		ScriptName:  "IsSubPath",
		Args: []*parser_tm.FuncArgNode{
			{Name: "parent", Type: parser_tm.NewBindingTypeName("any")},
			{Name: "child", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "isSub", Type: parser_tm.NewBindingTypeName("bool")},
		HasError: true,
	})

	this_.AddVar("LoadDirInfo", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "LoadDirInfo",
		Args: []*parser_tm.FuncArgNode{
			{Name: "dir", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "loadSubDir", Type: parser_tm.NewBindingTypeName("bool")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "dirInfo", Type: parser_tm.NewBindingTypeName("*DirInfo")},
		HasError: true,
	})
	utilSpace.AddVar("LoadDirInfo", &VarFunc{
		ScriptName:  "LoadDirInfo",
		Args: []*parser_tm.FuncArgNode{
			{Name: "dir", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "loadSubDir", Type: parser_tm.NewBindingTypeName("bool")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "dirInfo", Type: parser_tm.NewBindingTypeName("*DirInfo")},
		HasError: true,
	})

	this_.AddVar("GetFileType", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetFileType",
		Args: []*parser_tm.FuncArgNode{
			{Name: "fSrc", Type: parser_tm.NewBindingTypeName("[]byte")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("GetFileType", &VarFunc{
		ScriptName:  "GetFileType",
		Args: []*parser_tm.FuncArgNode{
			{Name: "fSrc", Type: parser_tm.NewBindingTypeName("[]byte")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("NextId", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "NextId",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})
	utilSpace.AddVar("NextId", &VarFunc{
		ScriptName:  "NextId",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})

	this_.AddVar("NewIdWorker", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "NewIdWorker",
		Args: []*parser_tm.FuncArgNode{
			{Name: "workerId", Type: parser_tm.NewBindingTypeName("int64")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("*IdWorker")},
		HasError: true,
	})
	utilSpace.AddVar("NewIdWorker", &VarFunc{
		ScriptName:  "NewIdWorker",
		Args: []*parser_tm.FuncArgNode{
			{Name: "workerId", Type: parser_tm.NewBindingTypeName("int64")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("*IdWorker")},
		HasError: true,
	})

	this_.AddVar("GetIpFromAddr", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetIpFromAddr",
		Args: []*parser_tm.FuncArgNode{
			{Name: "addr", Type: parser_tm.NewBindingTypeName("net.Addr")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("net.IP")},
	})
	utilSpace.AddVar("GetIpFromAddr", &VarFunc{
		ScriptName:  "GetIpFromAddr",
		Args: []*parser_tm.FuncArgNode{
			{Name: "addr", Type: parser_tm.NewBindingTypeName("net.Addr")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("net.IP")},
	})

	this_.AddVar("GetLocalIPList", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetLocalIPList",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "ipList", Type: parser_tm.NewBindingTypeList("net.IP")},
	})
	utilSpace.AddVar("GetLocalIPList", &VarFunc{
		ScriptName:  "GetLocalIPList",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "ipList", Type: parser_tm.NewBindingTypeList("net.IP")},
	})

	this_.AddVar("ObjToJson", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "ObjToJson",
		Args: []*parser_tm.FuncArgNode{
			{Name: "obj", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})
	utilSpace.AddVar("ObjToJson", &VarFunc{
		ScriptName:  "ObjToJson",
		Args: []*parser_tm.FuncArgNode{
			{Name: "obj", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})

	this_.AddVar("JsonToMap", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "JsonToMap",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("map[string]interface{}")},
		HasError: true,
	})
	utilSpace.AddVar("JsonToMap", &VarFunc{
		ScriptName:  "JsonToMap",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("map[string]interface{}")},
		HasError: true,
	})

	this_.AddVar("JsonToObj", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "JsonToObj",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "obj", Type: parser_tm.NewBindingTypeName("any")},
		},
		HasError: true,
	})
	utilSpace.AddVar("JsonToObj", &VarFunc{
		ScriptName:  "JsonToObj",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "obj", Type: parser_tm.NewBindingTypeName("any")},
		},
		HasError: true,
	})

	this_.AddVar("GetLock", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetLock",
		Args: []*parser_tm.FuncArgNode{
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "lock", Type: parser_tm.NewBindingTypeName("sync.Locker")},
	})
	utilSpace.AddVar("GetLock", &VarFunc{
		ScriptName:  "GetLock",
		Args: []*parser_tm.FuncArgNode{
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "lock", Type: parser_tm.NewBindingTypeName("sync.Locker")},
	})

	this_.AddVar("LockByKey", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "LockByKey",
		Args: []*parser_tm.FuncArgNode{
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("")},
	})
	utilSpace.AddVar("LockByKey", &VarFunc{
		ScriptName:  "LockByKey",
		Args: []*parser_tm.FuncArgNode{
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("")},
	})

	this_.AddVar("UnlockByKey", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "UnlockByKey",
		Args: []*parser_tm.FuncArgNode{
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("")},
	})
	utilSpace.AddVar("UnlockByKey", &VarFunc{
		ScriptName:  "UnlockByKey",
		Args: []*parser_tm.FuncArgNode{
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("")},
	})

	this_.AddVar("GetLogger", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetLogger",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("*zap.Logger")},
	})
	utilSpace.AddVar("GetLogger", &VarFunc{
		ScriptName:  "GetLogger",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("*zap.Logger")},
	})

	this_.AddVar("NewLoggerByCallerSkip", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "NewLoggerByCallerSkip",
		Args: []*parser_tm.FuncArgNode{
			{Name: "skip", Type: parser_tm.NewBindingTypeName("int")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("*zap.Logger")},
	})
	utilSpace.AddVar("NewLoggerByCallerSkip", &VarFunc{
		ScriptName:  "NewLoggerByCallerSkip",
		Args: []*parser_tm.FuncArgNode{
			{Name: "skip", Type: parser_tm.NewBindingTypeName("int")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("*zap.Logger")},
	})

	this_.AddVar("GetMD5", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetMD5",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("GetMD5", &VarFunc{
		ScriptName:  "GetMD5",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("RandomInt", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "RandomInt",
		Args: []*parser_tm.FuncArgNode{
			{Name: "min", Type: parser_tm.NewBindingTypeName("int")},
			{Name: "max", Type: parser_tm.NewBindingTypeName("int")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int")},
	})
	utilSpace.AddVar("RandomInt", &VarFunc{
		ScriptName:  "RandomInt",
		Args: []*parser_tm.FuncArgNode{
			{Name: "min", Type: parser_tm.NewBindingTypeName("int")},
			{Name: "max", Type: parser_tm.NewBindingTypeName("int")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int")},
	})

	this_.AddVar("RandomInt64", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "RandomInt64",
		Args: []*parser_tm.FuncArgNode{
			{Name: "min", Type: parser_tm.NewBindingTypeName("int64")},
			{Name: "max", Type: parser_tm.NewBindingTypeName("int64")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})
	utilSpace.AddVar("RandomInt64", &VarFunc{
		ScriptName:  "RandomInt64",
		Args: []*parser_tm.FuncArgNode{
			{Name: "min", Type: parser_tm.NewBindingTypeName("int64")},
			{Name: "max", Type: parser_tm.NewBindingTypeName("int64")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})

	this_.AddVar("StringToInt", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "StringToInt",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int")},
	})
	utilSpace.AddVar("StringToInt", &VarFunc{
		ScriptName:  "StringToInt",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int")},
	})

	this_.AddVar("StringToInt64", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "StringToInt64",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})
	utilSpace.AddVar("StringToInt64", &VarFunc{
		ScriptName:  "StringToInt64",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})

	this_.AddVar("StringToUint64", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "StringToUint64",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("uint64")},
	})
	utilSpace.AddVar("StringToUint64", &VarFunc{
		ScriptName:  "StringToUint64",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("uint64")},
	})

	this_.AddVar("StringToFloat64", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "StringToFloat64",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("float64")},
	})
	utilSpace.AddVar("StringToFloat64", &VarFunc{
		ScriptName:  "StringToFloat64",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("float64")},
	})

	this_.AddVar("SumToString", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "SumToString",
		Args: []*parser_tm.FuncArgNode{
			{Name: "nums", Type: parser_tm.NewBindingTypeName("...interface{}")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("SumToString", &VarFunc{
		ScriptName:  "SumToString",
		Args: []*parser_tm.FuncArgNode{
			{Name: "nums", Type: parser_tm.NewBindingTypeName("...interface{}")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("ValueToInt64", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "ValueToInt64",
		Args: []*parser_tm.FuncArgNode{
			{Name: "value", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
		HasError: true,
	})
	utilSpace.AddVar("ValueToInt64", &VarFunc{
		ScriptName:  "ValueToInt64",
		Args: []*parser_tm.FuncArgNode{
			{Name: "value", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
		HasError: true,
	})

	this_.AddVar("ValueToUint64", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "ValueToUint64",
		Args: []*parser_tm.FuncArgNode{
			{Name: "value", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("uint64")},
		HasError: true,
	})
	utilSpace.AddVar("ValueToUint64", &VarFunc{
		ScriptName:  "ValueToUint64",
		Args: []*parser_tm.FuncArgNode{
			{Name: "value", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("uint64")},
		HasError: true,
	})

	this_.AddVar("ValueToFloat64", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "ValueToFloat64",
		Args: []*parser_tm.FuncArgNode{
			{Name: "value", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("float64")},
		HasError: true,
	})
	utilSpace.AddVar("ValueToFloat64", &VarFunc{
		ScriptName:  "ValueToFloat64",
		Args: []*parser_tm.FuncArgNode{
			{Name: "value", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("float64")},
		HasError: true,
	})

	this_.AddVar("RsaEncryptByKey", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "RsaEncryptByKey",
		Args: []*parser_tm.FuncArgNode{
			{Name: "origData", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "publicKey", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})
	utilSpace.AddVar("RsaEncryptByKey", &VarFunc{
		ScriptName:  "RsaEncryptByKey",
		Args: []*parser_tm.FuncArgNode{
			{Name: "origData", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "publicKey", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})

	this_.AddVar("RsaDecryptByKey", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "RsaDecryptByKey",
		Args: []*parser_tm.FuncArgNode{
			{Name: "decrypt", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "privateKey", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})
	utilSpace.AddVar("RsaDecryptByKey", &VarFunc{
		ScriptName:  "RsaDecryptByKey",
		Args: []*parser_tm.FuncArgNode{
			{Name: "decrypt", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "privateKey", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
	})

	this_.AddVar("FirstToUpper", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "FirstToUpper",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("FirstToUpper", &VarFunc{
		ScriptName:  "FirstToUpper",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("FirstToLower", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "FirstToLower",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("FirstToLower", &VarFunc{
		ScriptName:  "FirstToLower",
		Args: []*parser_tm.FuncArgNode{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("Marshal", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "Marshal",
		Args: []*parser_tm.FuncArgNode{
			{Name: "name", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("Marshal", &VarFunc{
		ScriptName:  "Marshal",
		Args: []*parser_tm.FuncArgNode{
			{Name: "name", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("Hump", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "Hump",
		Args: []*parser_tm.FuncArgNode{
			{Name: "name", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("Hump", &VarFunc{
		ScriptName:  "Hump",
		Args: []*parser_tm.FuncArgNode{
			{Name: "name", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("GetStringValue", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetStringValue",
		Args: []*parser_tm.FuncArgNode{
			{Name: "value", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "valueString", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("GetStringValue", &VarFunc{
		ScriptName:  "GetStringValue",
		Args: []*parser_tm.FuncArgNode{
			{Name: "value", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "valueString", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("RandomString", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "RandomString",
		Args: []*parser_tm.FuncArgNode{
			{Name: "minLen", Type: parser_tm.NewBindingTypeName("int")},
			{Name: "maxLen", Type: parser_tm.NewBindingTypeName("int")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("RandomString", &VarFunc{
		ScriptName:  "RandomString",
		Args: []*parser_tm.FuncArgNode{
			{Name: "minLen", Type: parser_tm.NewBindingTypeName("int")},
			{Name: "maxLen", Type: parser_tm.NewBindingTypeName("int")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("RandomUserName", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "RandomUserName",
		Args: []*parser_tm.FuncArgNode{
			{Name: "size", Type: parser_tm.NewBindingTypeName("int")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("RandomUserName", &VarFunc{
		ScriptName:  "RandomUserName",
		Args: []*parser_tm.FuncArgNode{
			{Name: "size", Type: parser_tm.NewBindingTypeName("int")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("StrPadLeft", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "StrPadLeft",
		Args: []*parser_tm.FuncArgNode{
			{Name: "input", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "padLength", Type: parser_tm.NewBindingTypeName("int")},
			{Name: "padString", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("StrPadLeft", &VarFunc{
		ScriptName:  "StrPadLeft",
		Args: []*parser_tm.FuncArgNode{
			{Name: "input", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "padLength", Type: parser_tm.NewBindingTypeName("int")},
			{Name: "padString", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("StrPadRight", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "StrPadRight",
		Args: []*parser_tm.FuncArgNode{
			{Name: "input", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "padLength", Type: parser_tm.NewBindingTypeName("int")},
			{Name: "padString", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("StrPadRight", &VarFunc{
		ScriptName:  "StrPadRight",
		Args: []*parser_tm.FuncArgNode{
			{Name: "input", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "padLength", Type: parser_tm.NewBindingTypeName("int")},
			{Name: "padString", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("TrimSpace", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "TrimSpace",
		Args: []*parser_tm.FuncArgNode{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("TrimSpace", &VarFunc{
		ScriptName:  "TrimSpace",
		Args: []*parser_tm.FuncArgNode{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("TrimPrefix", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "TrimPrefix",
		Args: []*parser_tm.FuncArgNode{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("TrimPrefix", &VarFunc{
		ScriptName:  "TrimPrefix",
		Args: []*parser_tm.FuncArgNode{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("HasPrefix", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "HasPrefix",
		Args: []*parser_tm.FuncArgNode{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
	})
	utilSpace.AddVar("HasPrefix", &VarFunc{
		ScriptName:  "HasPrefix",
		Args: []*parser_tm.FuncArgNode{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
	})

	this_.AddVar("TrimSuffix", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "TrimSuffix",
		Args: []*parser_tm.FuncArgNode{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("TrimSuffix", &VarFunc{
		ScriptName:  "TrimSuffix",
		Args: []*parser_tm.FuncArgNode{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("HasSuffix", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "HasSuffix",
		Args: []*parser_tm.FuncArgNode{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
	})
	utilSpace.AddVar("HasSuffix", &VarFunc{
		ScriptName:  "HasSuffix",
		Args: []*parser_tm.FuncArgNode{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
	})

	this_.AddVar("TrimLeft", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "TrimLeft",
		Args: []*parser_tm.FuncArgNode{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("TrimLeft", &VarFunc{
		ScriptName:  "TrimLeft",
		Args: []*parser_tm.FuncArgNode{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("TrimRight", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "TrimRight",
		Args: []*parser_tm.FuncArgNode{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("TrimRight", &VarFunc{
		ScriptName:  "TrimRight",
		Args: []*parser_tm.FuncArgNode{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("StringJoin", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "StringJoin",
		Args: []*parser_tm.FuncArgNode{
			{Name: "es", Type: parser_tm.NewBindingTypeName("[]string")},
			{Name: "sep", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("StringJoin", &VarFunc{
		ScriptName:  "StringJoin",
		Args: []*parser_tm.FuncArgNode{
			{Name: "es", Type: parser_tm.NewBindingTypeName("[]string")},
			{Name: "sep", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("AnyJoin", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "AnyJoin",
		Args: []*parser_tm.FuncArgNode{
			{Name: "sep", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "es", Type: parser_tm.NewBindingTypeName("...any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("AnyJoin", &VarFunc{
		ScriptName:  "AnyJoin",
		Args: []*parser_tm.FuncArgNode{
			{Name: "sep", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "es", Type: parser_tm.NewBindingTypeName("...any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("IntJoin", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "IntJoin",
		Args: []*parser_tm.FuncArgNode{
			{Name: "es", Type: parser_tm.NewBindingTypeName("[]int")},
			{Name: "sep", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("IntJoin", &VarFunc{
		ScriptName:  "IntJoin",
		Args: []*parser_tm.FuncArgNode{
			{Name: "es", Type: parser_tm.NewBindingTypeName("[]int")},
			{Name: "sep", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("Int64Join", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "Int64Join",
		Args: []*parser_tm.FuncArgNode{
			{Name: "es", Type: parser_tm.NewBindingTypeName("[]int64")},
			{Name: "sep", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("Int64Join", &VarFunc{
		ScriptName:  "Int64Join",
		Args: []*parser_tm.FuncArgNode{
			{Name: "es", Type: parser_tm.NewBindingTypeName("[]int64")},
			{Name: "sep", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("GenStringJoin", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GenStringJoin",
		Args: []*parser_tm.FuncArgNode{
			{Name: "len", Type: parser_tm.NewBindingTypeName("int")},
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "sep", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("GenStringJoin", &VarFunc{
		ScriptName:  "GenStringJoin",
		Args: []*parser_tm.FuncArgNode{
			{Name: "len", Type: parser_tm.NewBindingTypeName("int")},
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "sep", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("GetNow", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetNow",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("time.Time")},
	})
	utilSpace.AddVar("GetNow", &VarFunc{
		ScriptName:  "GetNow",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("time.Time")},
	})

	this_.AddVar("GetNowNano", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetNowNano",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})
	utilSpace.AddVar("GetNowNano", &VarFunc{
		ScriptName:  "GetNowNano",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})

	this_.AddVar("GetNowMilli", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetNowMilli",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})
	utilSpace.AddVar("GetNowMilli", &VarFunc{
		ScriptName:  "GetNowMilli",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})

	this_.AddVar("GetNowSecond", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetNowSecond",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})
	utilSpace.AddVar("GetNowSecond", &VarFunc{
		ScriptName:  "GetNowSecond",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})

	this_.AddVar("GetNanoByTime", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetNanoByTime",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("time.Time")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})
	utilSpace.AddVar("GetNanoByTime", &VarFunc{
		ScriptName:  "GetNanoByTime",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("time.Time")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})

	this_.AddVar("GetMilliByTime", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetMilliByTime",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("time.Time")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})
	utilSpace.AddVar("GetMilliByTime", &VarFunc{
		ScriptName:  "GetMilliByTime",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("time.Time")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})

	this_.AddVar("GetSecondByTime", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetSecondByTime",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("time.Time")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})
	utilSpace.AddVar("GetSecondByTime", &VarFunc{
		ScriptName:  "GetSecondByTime",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("time.Time")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
	})

	this_.AddVar("GetNowFormat", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetNowFormat",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("GetNowFormat", &VarFunc{
		ScriptName:  "GetNowFormat",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("GetFormatByTime", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetFormatByTime",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("time.Time")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("GetFormatByTime", &VarFunc{
		ScriptName:  "GetFormatByTime",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("time.Time")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("TimeFormat", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "TimeFormat",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("time.Time")},
			{Name: "layout", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("TimeFormat", &VarFunc{
		ScriptName:  "TimeFormat",
		Args: []*parser_tm.FuncArgNode{
			{Name: "v", Type: parser_tm.NewBindingTypeName("time.Time")},
			{Name: "layout", Type: parser_tm.NewBindingTypeName("string")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("MilliToTimeText", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "MilliToTimeText",
		Args: []*parser_tm.FuncArgNode{
			{Name: "milli", Type: parser_tm.NewBindingTypeName("int64")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "v", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("MilliToTimeText", &VarFunc{
		ScriptName:  "MilliToTimeText",
		Args: []*parser_tm.FuncArgNode{
			{Name: "milli", Type: parser_tm.NewBindingTypeName("int64")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "v", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("GetUUID", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetUUID",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})
	utilSpace.AddVar("GetUUID", &VarFunc{
		ScriptName:  "GetUUID",
		Args: []*parser_tm.FuncArgNode{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
	})

	this_.AddVar("GzipBytes", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GzipBytes",
		Args: []*parser_tm.FuncArgNode{
			{Name: "data", Type: parser_tm.NewBindingTypeName("[]byte")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeList("byte")},
		HasError: true,
	})
	utilSpace.AddVar("GzipBytes", &VarFunc{
		ScriptName:  "GzipBytes",
		Args: []*parser_tm.FuncArgNode{
			{Name: "data", Type: parser_tm.NewBindingTypeName("[]byte")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeList("byte")},
		HasError: true,
	})

	this_.AddVar("UnGzipBytes", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "UnGzipBytes",
		Args: []*parser_tm.FuncArgNode{
			{Name: "data", Type: parser_tm.NewBindingTypeName("[]byte")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeList("byte")},
		HasError: true,
	})
	utilSpace.AddVar("UnGzipBytes", &VarFunc{
		ScriptName:  "UnGzipBytes",
		Args: []*parser_tm.FuncArgNode{
			{Name: "data", Type: parser_tm.NewBindingTypeName("[]byte")},
		},
		Return: &parser_tm.FuncReturnNode{Name: "res", Type: parser_tm.NewBindingTypeList("byte")},
		HasError: true,
	})

	this_.AddVar("Zip", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "Zip",
		Args: []*parser_tm.FuncArgNode{
			{Name: "srcFile", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "destZip", Type: parser_tm.NewBindingTypeName("string")},
		},
		HasError: true,
	})
	utilSpace.AddVar("Zip", &VarFunc{
		ScriptName:  "Zip",
		Args: []*parser_tm.FuncArgNode{
			{Name: "srcFile", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "destZip", Type: parser_tm.NewBindingTypeName("string")},
		},
		HasError: true,
	})

	this_.AddVar("UnZip", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "UnZip",
		Args: []*parser_tm.FuncArgNode{
			{Name: "zipFile", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "destDir", Type: parser_tm.NewBindingTypeName("string")},
		},
		HasError: true,
	})
	utilSpace.AddVar("UnZip", &VarFunc{
		ScriptName:  "UnZip",
		Args: []*parser_tm.FuncArgNode{
			{Name: "zipFile", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "destDir", Type: parser_tm.NewBindingTypeName("string")},
		},
		HasError: true,
	})

	return
}