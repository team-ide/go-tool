package compiler

import (
	"github.com/team-ide/go-tool/util"
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
		Args: []*FuncArg{
			{Name: "origData", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
		callFunc:  util.AesEncryptCBCByKey,
	})
	utilSpace.AddVar("AesEncryptCBCByKey", &VarFunc{
		ScriptName:  "AesEncryptCBCByKey",
		Args: []*FuncArg{
			{Name: "origData", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
		callFunc:  util.AesEncryptCBCByKey,
	})

	this_.AddVar("AesDecryptCBCByKey", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "AesDecryptCBCByKey",
		Args: []*FuncArg{
			{Name: "encrypt", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
		callFunc:  util.AesDecryptCBCByKey,
	})
	utilSpace.AddVar("AesDecryptCBCByKey", &VarFunc{
		ScriptName:  "AesDecryptCBCByKey",
		Args: []*FuncArg{
			{Name: "encrypt", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
		callFunc:  util.AesDecryptCBCByKey,
	})

	this_.AddVar("AesEncryptECBByKey", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "AesEncryptECBByKey",
		Args: []*FuncArg{
			{Name: "origData", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
		callFunc:  util.AesEncryptECBByKey,
	})
	utilSpace.AddVar("AesEncryptECBByKey", &VarFunc{
		ScriptName:  "AesEncryptECBByKey",
		Args: []*FuncArg{
			{Name: "origData", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
		callFunc:  util.AesEncryptECBByKey,
	})

	this_.AddVar("AesDecryptECBByKey", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "AesDecryptECBByKey",
		Args: []*FuncArg{
			{Name: "encrypt", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
		callFunc:  util.AesDecryptECBByKey,
	})
	utilSpace.AddVar("AesDecryptECBByKey", &VarFunc{
		ScriptName:  "AesDecryptECBByKey",
		Args: []*FuncArg{
			{Name: "encrypt", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
		callFunc:  util.AesDecryptECBByKey,
	})

	this_.AddVar("IsEmpty", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "IsEmpty",
		Args: []*FuncArg{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
		callFunc:  util.IsEmpty,
	})
	utilSpace.AddVar("IsEmpty", &VarFunc{
		ScriptName:  "IsEmpty",
		Args: []*FuncArg{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
		callFunc:  util.IsEmpty,
	})

	this_.AddVar("IsNotEmpty", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "IsNotEmpty",
		Args: []*FuncArg{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
		callFunc:  util.IsNotEmpty,
	})
	utilSpace.AddVar("IsNotEmpty", &VarFunc{
		ScriptName:  "IsNotEmpty",
		Args: []*FuncArg{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
		callFunc:  util.IsNotEmpty,
	})

	this_.AddVar("IsNull", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "IsNull",
		Args: []*FuncArg{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
		callFunc:  util.IsNull,
	})
	utilSpace.AddVar("IsNull", &VarFunc{
		ScriptName:  "IsNull",
		Args: []*FuncArg{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
		callFunc:  util.IsNull,
	})

	this_.AddVar("IsNotNull", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "IsNotNull",
		Args: []*FuncArg{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
		callFunc:  util.IsNotNull,
	})
	utilSpace.AddVar("IsNotNull", &VarFunc{
		ScriptName:  "IsNotNull",
		Args: []*FuncArg{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
		callFunc:  util.IsNotNull,
	})

	this_.AddVar("IsTrue", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "IsTrue",
		Args: []*FuncArg{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
		callFunc:  util.IsTrue,
	})
	utilSpace.AddVar("IsTrue", &VarFunc{
		ScriptName:  "IsTrue",
		Args: []*FuncArg{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
		callFunc:  util.IsTrue,
	})

	this_.AddVar("IsFalse", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "IsFalse",
		Args: []*FuncArg{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
		callFunc:  util.IsFalse,
	})
	utilSpace.AddVar("IsFalse", &VarFunc{
		ScriptName:  "IsFalse",
		Args: []*FuncArg{
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
		callFunc:  util.IsFalse,
	})

	this_.AddVar("IntIndexOf", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "IntIndexOf",
		Args: []*FuncArg{
			{Name: "array", Type: parser_tm.NewBindingTypeName("[]int")},
			{Name: "v", Type: parser_tm.NewBindingTypeName("int")},
		},
		Result: &FuncResult{Name: "index", Type: parser_tm.NewBindingTypeName("int")},
		callFunc:  util.IntIndexOf,
	})
	utilSpace.AddVar("IntIndexOf", &VarFunc{
		ScriptName:  "IntIndexOf",
		Args: []*FuncArg{
			{Name: "array", Type: parser_tm.NewBindingTypeName("[]int")},
			{Name: "v", Type: parser_tm.NewBindingTypeName("int")},
		},
		Result: &FuncResult{Name: "index", Type: parser_tm.NewBindingTypeName("int")},
		callFunc:  util.IntIndexOf,
	})

	this_.AddVar("Int64IndexOf", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "Int64IndexOf",
		Args: []*FuncArg{
			{Name: "array", Type: parser_tm.NewBindingTypeName("[]int64")},
			{Name: "v", Type: parser_tm.NewBindingTypeName("int64")},
		},
		Result: &FuncResult{Name: "index", Type: parser_tm.NewBindingTypeName("int")},
		callFunc:  util.Int64IndexOf,
	})
	utilSpace.AddVar("Int64IndexOf", &VarFunc{
		ScriptName:  "Int64IndexOf",
		Args: []*FuncArg{
			{Name: "array", Type: parser_tm.NewBindingTypeName("[]int64")},
			{Name: "v", Type: parser_tm.NewBindingTypeName("int64")},
		},
		Result: &FuncResult{Name: "index", Type: parser_tm.NewBindingTypeName("int")},
		callFunc:  util.Int64IndexOf,
	})

	this_.AddVar("StringIndexOf", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "StringIndexOf",
		Args: []*FuncArg{
			{Name: "array", Type: parser_tm.NewBindingTypeName("[]string")},
			{Name: "v", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "index", Type: parser_tm.NewBindingTypeName("int")},
		callFunc:  util.StringIndexOf,
	})
	utilSpace.AddVar("StringIndexOf", &VarFunc{
		ScriptName:  "StringIndexOf",
		Args: []*FuncArg{
			{Name: "array", Type: parser_tm.NewBindingTypeName("[]string")},
			{Name: "v", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "index", Type: parser_tm.NewBindingTypeName("int")},
		callFunc:  util.StringIndexOf,
	})

	this_.AddVar("ArrayIndexOf", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "ArrayIndexOf",
		Args: []*FuncArg{
			{Name: "array", Type: parser_tm.NewBindingTypeName("any")},
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "index", Type: parser_tm.NewBindingTypeName("int")},
		callFunc:  util.ArrayIndexOf,
	})
	utilSpace.AddVar("ArrayIndexOf", &VarFunc{
		ScriptName:  "ArrayIndexOf",
		Args: []*FuncArg{
			{Name: "array", Type: parser_tm.NewBindingTypeName("any")},
			{Name: "v", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "index", Type: parser_tm.NewBindingTypeName("int")},
		callFunc:  util.ArrayIndexOf,
	})

	this_.AddVar("GetTempDir", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetTempDir",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "dir", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
		callFunc:  util.GetTempDir,
	})
	utilSpace.AddVar("GetTempDir", &VarFunc{
		ScriptName:  "GetTempDir",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "dir", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
		callFunc:  util.GetTempDir,
	})

	this_.AddVar("NewWaitGroup", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "NewWaitGroup",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("*sync.WaitGroup")},
		callFunc:  util.NewWaitGroup,
	})
	utilSpace.AddVar("NewWaitGroup", &VarFunc{
		ScriptName:  "NewWaitGroup",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("*sync.WaitGroup")},
		callFunc:  util.NewWaitGroup,
	})

	this_.AddVar("NewLocker", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "NewLocker",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("sync.Locker")},
		callFunc:  util.NewLocker,
	})
	utilSpace.AddVar("NewLocker", &VarFunc{
		ScriptName:  "NewLocker",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("sync.Locker")},
		callFunc:  util.NewLocker,
	})

	this_.AddVar("GetRootDir", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetRootDir",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.GetRootDir,
	})
	utilSpace.AddVar("GetRootDir", &VarFunc{
		ScriptName:  "GetRootDir",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.GetRootDir,
	})

	this_.AddVar("FormatPath", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "FormatPath",
		Args: []*FuncArg{
			{Name: "path", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.FormatPath,
	})
	utilSpace.AddVar("FormatPath", &VarFunc{
		ScriptName:  "FormatPath",
		Args: []*FuncArg{
			{Name: "path", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.FormatPath,
	})

	this_.AddVar("GetAbsolutePath", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetAbsolutePath",
		Args: []*FuncArg{
			{Name: "path", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "absolutePath", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.GetAbsolutePath,
	})
	utilSpace.AddVar("GetAbsolutePath", &VarFunc{
		ScriptName:  "GetAbsolutePath",
		Args: []*FuncArg{
			{Name: "path", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "absolutePath", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.GetAbsolutePath,
	})

	this_.AddVar("PathExists", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "PathExists",
		Args: []*FuncArg{
			{Name: "path", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
		HasError: true,
		callFunc:  util.PathExists,
	})
	utilSpace.AddVar("PathExists", &VarFunc{
		ScriptName:  "PathExists",
		Args: []*FuncArg{
			{Name: "path", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
		HasError: true,
		callFunc:  util.PathExists,
	})

	this_.AddVar("LoadDirFiles", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "LoadDirFiles",
		Args: []*FuncArg{
			{Name: "dir", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "fileMap", Type: parser_tm.NewBindingTypeName("map[string][]byte")},
		HasError: true,
		callFunc:  util.LoadDirFiles,
	})
	utilSpace.AddVar("LoadDirFiles", &VarFunc{
		ScriptName:  "LoadDirFiles",
		Args: []*FuncArg{
			{Name: "dir", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "fileMap", Type: parser_tm.NewBindingTypeName("map[string][]byte")},
		HasError: true,
		callFunc:  util.LoadDirFiles,
	})

	this_.AddVar("LoadDirFilenames", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "LoadDirFilenames",
		Args: []*FuncArg{
			{Name: "dir", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "filenames", Type: parser_tm.NewBindingTypeList("string")},
		HasError: true,
		callFunc:  util.LoadDirFilenames,
	})
	utilSpace.AddVar("LoadDirFilenames", &VarFunc{
		ScriptName:  "LoadDirFilenames",
		Args: []*FuncArg{
			{Name: "dir", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "filenames", Type: parser_tm.NewBindingTypeList("string")},
		HasError: true,
		callFunc:  util.LoadDirFilenames,
	})

	this_.AddVar("ReadFile", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "ReadFile",
		Args: []*FuncArg{
			{Name: "filename", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "bs", Type: parser_tm.NewBindingTypeList("byte")},
		HasError: true,
		callFunc:  util.ReadFile,
	})
	utilSpace.AddVar("ReadFile", &VarFunc{
		ScriptName:  "ReadFile",
		Args: []*FuncArg{
			{Name: "filename", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "bs", Type: parser_tm.NewBindingTypeList("byte")},
		HasError: true,
		callFunc:  util.ReadFile,
	})

	this_.AddVar("ReadFileString", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "ReadFileString",
		Args: []*FuncArg{
			{Name: "filename", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
		callFunc:  util.ReadFileString,
	})
	utilSpace.AddVar("ReadFileString", &VarFunc{
		ScriptName:  "ReadFileString",
		Args: []*FuncArg{
			{Name: "filename", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
		callFunc:  util.ReadFileString,
	})

	this_.AddVar("StringToBytes", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "StringToBytes",
		Args: []*FuncArg{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeList("byte")},
		callFunc:  util.StringToBytes,
	})
	utilSpace.AddVar("StringToBytes", &VarFunc{
		ScriptName:  "StringToBytes",
		Args: []*FuncArg{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeList("byte")},
		callFunc:  util.StringToBytes,
	})

	this_.AddVar("WriteFile", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "WriteFile",
		Args: []*FuncArg{
			{Name: "filename", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "bs", Type: parser_tm.NewBindingTypeName("[]byte")},
		},
		HasError: true,
		callFunc:  util.WriteFile,
	})
	utilSpace.AddVar("WriteFile", &VarFunc{
		ScriptName:  "WriteFile",
		Args: []*FuncArg{
			{Name: "filename", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "bs", Type: parser_tm.NewBindingTypeName("[]byte")},
		},
		HasError: true,
		callFunc:  util.WriteFile,
	})

	this_.AddVar("WriteFileString", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "WriteFileString",
		Args: []*FuncArg{
			{Name: "filename", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		HasError: true,
		callFunc:  util.WriteFileString,
	})
	utilSpace.AddVar("WriteFileString", &VarFunc{
		ScriptName:  "WriteFileString",
		Args: []*FuncArg{
			{Name: "filename", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		HasError: true,
		callFunc:  util.WriteFileString,
	})

	this_.AddVar("ReadLine", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "ReadLine",
		Args: []*FuncArg{
			{Name: "filename", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "lines", Type: parser_tm.NewBindingTypeList("string")},
		HasError: true,
		callFunc:  util.ReadLine,
	})
	utilSpace.AddVar("ReadLine", &VarFunc{
		ScriptName:  "ReadLine",
		Args: []*FuncArg{
			{Name: "filename", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "lines", Type: parser_tm.NewBindingTypeList("string")},
		HasError: true,
		callFunc:  util.ReadLine,
	})

	this_.AddVar("IsSubPath", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "IsSubPath",
		Args: []*FuncArg{
			{Name: "parent", Type: parser_tm.NewBindingTypeName("any")},
			{Name: "child", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "isSub", Type: parser_tm.NewBindingTypeName("bool")},
		HasError: true,
		callFunc:  util.IsSubPath,
	})
	utilSpace.AddVar("IsSubPath", &VarFunc{
		ScriptName:  "IsSubPath",
		Args: []*FuncArg{
			{Name: "parent", Type: parser_tm.NewBindingTypeName("any")},
			{Name: "child", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "isSub", Type: parser_tm.NewBindingTypeName("bool")},
		HasError: true,
		callFunc:  util.IsSubPath,
	})

	this_.AddVar("LoadDirInfo", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "LoadDirInfo",
		Args: []*FuncArg{
			{Name: "dir", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "loadSubDir", Type: parser_tm.NewBindingTypeName("bool")},
		},
		Result: &FuncResult{Name: "dirInfo", Type: parser_tm.NewBindingTypeName("*DirInfo")},
		HasError: true,
		callFunc:  util.LoadDirInfo,
	})
	utilSpace.AddVar("LoadDirInfo", &VarFunc{
		ScriptName:  "LoadDirInfo",
		Args: []*FuncArg{
			{Name: "dir", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "loadSubDir", Type: parser_tm.NewBindingTypeName("bool")},
		},
		Result: &FuncResult{Name: "dirInfo", Type: parser_tm.NewBindingTypeName("*DirInfo")},
		HasError: true,
		callFunc:  util.LoadDirInfo,
	})

	this_.AddVar("GetFileType", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetFileType",
		Args: []*FuncArg{
			{Name: "fSrc", Type: parser_tm.NewBindingTypeName("[]byte")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.GetFileType,
	})
	utilSpace.AddVar("GetFileType", &VarFunc{
		ScriptName:  "GetFileType",
		Args: []*FuncArg{
			{Name: "fSrc", Type: parser_tm.NewBindingTypeName("[]byte")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.GetFileType,
	})

	this_.AddVar("NextId", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "NextId",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
		callFunc:  util.NextId,
	})
	utilSpace.AddVar("NextId", &VarFunc{
		ScriptName:  "NextId",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
		callFunc:  util.NextId,
	})

	this_.AddVar("NewIdWorker", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "NewIdWorker",
		Args: []*FuncArg{
			{Name: "workerId", Type: parser_tm.NewBindingTypeName("int64")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("*IdWorker")},
		HasError: true,
		callFunc:  util.NewIdWorker,
	})
	utilSpace.AddVar("NewIdWorker", &VarFunc{
		ScriptName:  "NewIdWorker",
		Args: []*FuncArg{
			{Name: "workerId", Type: parser_tm.NewBindingTypeName("int64")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("*IdWorker")},
		HasError: true,
		callFunc:  util.NewIdWorker,
	})

	this_.AddVar("GetIpFromAddr", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetIpFromAddr",
		Args: []*FuncArg{
			{Name: "addr", Type: parser_tm.NewBindingTypeName("net.Addr")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("net.IP")},
		callFunc:  util.GetIpFromAddr,
	})
	utilSpace.AddVar("GetIpFromAddr", &VarFunc{
		ScriptName:  "GetIpFromAddr",
		Args: []*FuncArg{
			{Name: "addr", Type: parser_tm.NewBindingTypeName("net.Addr")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("net.IP")},
		callFunc:  util.GetIpFromAddr,
	})

	this_.AddVar("GetLocalIPList", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetLocalIPList",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "ipList", Type: parser_tm.NewBindingTypeList("net.IP")},
		callFunc:  util.GetLocalIPList,
	})
	utilSpace.AddVar("GetLocalIPList", &VarFunc{
		ScriptName:  "GetLocalIPList",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "ipList", Type: parser_tm.NewBindingTypeList("net.IP")},
		callFunc:  util.GetLocalIPList,
	})

	this_.AddVar("ObjToJson", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "ObjToJson",
		Args: []*FuncArg{
			{Name: "obj", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
		callFunc:  util.ObjToJson,
	})
	utilSpace.AddVar("ObjToJson", &VarFunc{
		ScriptName:  "ObjToJson",
		Args: []*FuncArg{
			{Name: "obj", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
		callFunc:  util.ObjToJson,
	})

	this_.AddVar("JsonToMap", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "JsonToMap",
		Args: []*FuncArg{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("map[string]interface{}")},
		HasError: true,
		callFunc:  util.JsonToMap,
	})
	utilSpace.AddVar("JsonToMap", &VarFunc{
		ScriptName:  "JsonToMap",
		Args: []*FuncArg{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("map[string]interface{}")},
		HasError: true,
		callFunc:  util.JsonToMap,
	})

	this_.AddVar("JsonToObj", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "JsonToObj",
		Args: []*FuncArg{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "obj", Type: parser_tm.NewBindingTypeName("any")},
		},
		HasError: true,
		callFunc:  util.JsonToObj,
	})
	utilSpace.AddVar("JsonToObj", &VarFunc{
		ScriptName:  "JsonToObj",
		Args: []*FuncArg{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "obj", Type: parser_tm.NewBindingTypeName("any")},
		},
		HasError: true,
		callFunc:  util.JsonToObj,
	})

	this_.AddVar("GetLock", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetLock",
		Args: []*FuncArg{
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "lock", Type: parser_tm.NewBindingTypeName("sync.Locker")},
		callFunc:  util.GetLock,
	})
	utilSpace.AddVar("GetLock", &VarFunc{
		ScriptName:  "GetLock",
		Args: []*FuncArg{
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "lock", Type: parser_tm.NewBindingTypeName("sync.Locker")},
		callFunc:  util.GetLock,
	})

	this_.AddVar("LockByKey", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "LockByKey",
		Args: []*FuncArg{
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("")},
		callFunc:  util.LockByKey,
	})
	utilSpace.AddVar("LockByKey", &VarFunc{
		ScriptName:  "LockByKey",
		Args: []*FuncArg{
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("")},
		callFunc:  util.LockByKey,
	})

	this_.AddVar("UnlockByKey", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "UnlockByKey",
		Args: []*FuncArg{
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("")},
		callFunc:  util.UnlockByKey,
	})
	utilSpace.AddVar("UnlockByKey", &VarFunc{
		ScriptName:  "UnlockByKey",
		Args: []*FuncArg{
			{Name: "key", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("")},
		callFunc:  util.UnlockByKey,
	})

	this_.AddVar("GetLogger", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetLogger",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("*zap.Logger")},
		callFunc:  util.GetLogger,
	})
	utilSpace.AddVar("GetLogger", &VarFunc{
		ScriptName:  "GetLogger",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("*zap.Logger")},
		callFunc:  util.GetLogger,
	})

	this_.AddVar("NewLoggerByCallerSkip", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "NewLoggerByCallerSkip",
		Args: []*FuncArg{
			{Name: "skip", Type: parser_tm.NewBindingTypeName("int")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("*zap.Logger")},
		callFunc:  util.NewLoggerByCallerSkip,
	})
	utilSpace.AddVar("NewLoggerByCallerSkip", &VarFunc{
		ScriptName:  "NewLoggerByCallerSkip",
		Args: []*FuncArg{
			{Name: "skip", Type: parser_tm.NewBindingTypeName("int")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("*zap.Logger")},
		callFunc:  util.NewLoggerByCallerSkip,
	})

	this_.AddVar("GetMD5", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetMD5",
		Args: []*FuncArg{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.GetMD5,
	})
	utilSpace.AddVar("GetMD5", &VarFunc{
		ScriptName:  "GetMD5",
		Args: []*FuncArg{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.GetMD5,
	})

	this_.AddVar("RandomInt", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "RandomInt",
		Args: []*FuncArg{
			{Name: "min", Type: parser_tm.NewBindingTypeName("int")},
			{Name: "max", Type: parser_tm.NewBindingTypeName("int")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("int")},
		callFunc:  util.RandomInt,
	})
	utilSpace.AddVar("RandomInt", &VarFunc{
		ScriptName:  "RandomInt",
		Args: []*FuncArg{
			{Name: "min", Type: parser_tm.NewBindingTypeName("int")},
			{Name: "max", Type: parser_tm.NewBindingTypeName("int")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("int")},
		callFunc:  util.RandomInt,
	})

	this_.AddVar("RandomInt64", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "RandomInt64",
		Args: []*FuncArg{
			{Name: "min", Type: parser_tm.NewBindingTypeName("int64")},
			{Name: "max", Type: parser_tm.NewBindingTypeName("int64")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
		callFunc:  util.RandomInt64,
	})
	utilSpace.AddVar("RandomInt64", &VarFunc{
		ScriptName:  "RandomInt64",
		Args: []*FuncArg{
			{Name: "min", Type: parser_tm.NewBindingTypeName("int64")},
			{Name: "max", Type: parser_tm.NewBindingTypeName("int64")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
		callFunc:  util.RandomInt64,
	})

	this_.AddVar("StringToInt", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "StringToInt",
		Args: []*FuncArg{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("int")},
		callFunc:  util.StringToInt,
	})
	utilSpace.AddVar("StringToInt", &VarFunc{
		ScriptName:  "StringToInt",
		Args: []*FuncArg{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("int")},
		callFunc:  util.StringToInt,
	})

	this_.AddVar("StringToInt64", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "StringToInt64",
		Args: []*FuncArg{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
		callFunc:  util.StringToInt64,
	})
	utilSpace.AddVar("StringToInt64", &VarFunc{
		ScriptName:  "StringToInt64",
		Args: []*FuncArg{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
		callFunc:  util.StringToInt64,
	})

	this_.AddVar("StringToUint64", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "StringToUint64",
		Args: []*FuncArg{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("uint64")},
		callFunc:  util.StringToUint64,
	})
	utilSpace.AddVar("StringToUint64", &VarFunc{
		ScriptName:  "StringToUint64",
		Args: []*FuncArg{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("uint64")},
		callFunc:  util.StringToUint64,
	})

	this_.AddVar("StringToFloat64", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "StringToFloat64",
		Args: []*FuncArg{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("float64")},
		callFunc:  util.StringToFloat64,
	})
	utilSpace.AddVar("StringToFloat64", &VarFunc{
		ScriptName:  "StringToFloat64",
		Args: []*FuncArg{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("float64")},
		callFunc:  util.StringToFloat64,
	})

	this_.AddVar("SumToString", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "SumToString",
		Args: []*FuncArg{
			{Name: "nums", Type: parser_tm.NewBindingTypeName("...interface{}")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.SumToString,
	})
	utilSpace.AddVar("SumToString", &VarFunc{
		ScriptName:  "SumToString",
		Args: []*FuncArg{
			{Name: "nums", Type: parser_tm.NewBindingTypeName("...interface{}")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.SumToString,
	})

	this_.AddVar("ValueToInt64", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "ValueToInt64",
		Args: []*FuncArg{
			{Name: "value", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
		HasError: true,
		callFunc:  util.ValueToInt64,
	})
	utilSpace.AddVar("ValueToInt64", &VarFunc{
		ScriptName:  "ValueToInt64",
		Args: []*FuncArg{
			{Name: "value", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
		HasError: true,
		callFunc:  util.ValueToInt64,
	})

	this_.AddVar("ValueToUint64", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "ValueToUint64",
		Args: []*FuncArg{
			{Name: "value", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("uint64")},
		HasError: true,
		callFunc:  util.ValueToUint64,
	})
	utilSpace.AddVar("ValueToUint64", &VarFunc{
		ScriptName:  "ValueToUint64",
		Args: []*FuncArg{
			{Name: "value", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("uint64")},
		HasError: true,
		callFunc:  util.ValueToUint64,
	})

	this_.AddVar("ValueToFloat64", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "ValueToFloat64",
		Args: []*FuncArg{
			{Name: "value", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("float64")},
		HasError: true,
		callFunc:  util.ValueToFloat64,
	})
	utilSpace.AddVar("ValueToFloat64", &VarFunc{
		ScriptName:  "ValueToFloat64",
		Args: []*FuncArg{
			{Name: "value", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("float64")},
		HasError: true,
		callFunc:  util.ValueToFloat64,
	})

	this_.AddVar("RsaEncryptByKey", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "RsaEncryptByKey",
		Args: []*FuncArg{
			{Name: "origData", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "publicKey", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
		callFunc:  util.RsaEncryptByKey,
	})
	utilSpace.AddVar("RsaEncryptByKey", &VarFunc{
		ScriptName:  "RsaEncryptByKey",
		Args: []*FuncArg{
			{Name: "origData", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "publicKey", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
		callFunc:  util.RsaEncryptByKey,
	})

	this_.AddVar("RsaDecryptByKey", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "RsaDecryptByKey",
		Args: []*FuncArg{
			{Name: "decrypt", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "privateKey", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
		callFunc:  util.RsaDecryptByKey,
	})
	utilSpace.AddVar("RsaDecryptByKey", &VarFunc{
		ScriptName:  "RsaDecryptByKey",
		Args: []*FuncArg{
			{Name: "decrypt", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "privateKey", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		HasError: true,
		callFunc:  util.RsaDecryptByKey,
	})

	this_.AddVar("FirstToUpper", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "FirstToUpper",
		Args: []*FuncArg{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.FirstToUpper,
	})
	utilSpace.AddVar("FirstToUpper", &VarFunc{
		ScriptName:  "FirstToUpper",
		Args: []*FuncArg{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.FirstToUpper,
	})

	this_.AddVar("FirstToLower", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "FirstToLower",
		Args: []*FuncArg{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.FirstToLower,
	})
	utilSpace.AddVar("FirstToLower", &VarFunc{
		ScriptName:  "FirstToLower",
		Args: []*FuncArg{
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.FirstToLower,
	})

	this_.AddVar("Marshal", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "Marshal",
		Args: []*FuncArg{
			{Name: "name", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.Marshal,
	})
	utilSpace.AddVar("Marshal", &VarFunc{
		ScriptName:  "Marshal",
		Args: []*FuncArg{
			{Name: "name", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.Marshal,
	})

	this_.AddVar("Hump", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "Hump",
		Args: []*FuncArg{
			{Name: "name", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.Hump,
	})
	utilSpace.AddVar("Hump", &VarFunc{
		ScriptName:  "Hump",
		Args: []*FuncArg{
			{Name: "name", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.Hump,
	})

	this_.AddVar("GetStringValue", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetStringValue",
		Args: []*FuncArg{
			{Name: "value", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "valueString", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.GetStringValue,
	})
	utilSpace.AddVar("GetStringValue", &VarFunc{
		ScriptName:  "GetStringValue",
		Args: []*FuncArg{
			{Name: "value", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "valueString", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.GetStringValue,
	})

	this_.AddVar("RandomString", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "RandomString",
		Args: []*FuncArg{
			{Name: "minLen", Type: parser_tm.NewBindingTypeName("int")},
			{Name: "maxLen", Type: parser_tm.NewBindingTypeName("int")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.RandomString,
	})
	utilSpace.AddVar("RandomString", &VarFunc{
		ScriptName:  "RandomString",
		Args: []*FuncArg{
			{Name: "minLen", Type: parser_tm.NewBindingTypeName("int")},
			{Name: "maxLen", Type: parser_tm.NewBindingTypeName("int")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.RandomString,
	})

	this_.AddVar("RandomUserName", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "RandomUserName",
		Args: []*FuncArg{
			{Name: "size", Type: parser_tm.NewBindingTypeName("int")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.RandomUserName,
	})
	utilSpace.AddVar("RandomUserName", &VarFunc{
		ScriptName:  "RandomUserName",
		Args: []*FuncArg{
			{Name: "size", Type: parser_tm.NewBindingTypeName("int")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.RandomUserName,
	})

	this_.AddVar("StrPadLeft", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "StrPadLeft",
		Args: []*FuncArg{
			{Name: "input", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "padLength", Type: parser_tm.NewBindingTypeName("int")},
			{Name: "padString", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.StrPadLeft,
	})
	utilSpace.AddVar("StrPadLeft", &VarFunc{
		ScriptName:  "StrPadLeft",
		Args: []*FuncArg{
			{Name: "input", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "padLength", Type: parser_tm.NewBindingTypeName("int")},
			{Name: "padString", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.StrPadLeft,
	})

	this_.AddVar("StrPadRight", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "StrPadRight",
		Args: []*FuncArg{
			{Name: "input", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "padLength", Type: parser_tm.NewBindingTypeName("int")},
			{Name: "padString", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.StrPadRight,
	})
	utilSpace.AddVar("StrPadRight", &VarFunc{
		ScriptName:  "StrPadRight",
		Args: []*FuncArg{
			{Name: "input", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "padLength", Type: parser_tm.NewBindingTypeName("int")},
			{Name: "padString", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.StrPadRight,
	})

	this_.AddVar("TrimSpace", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "TrimSpace",
		Args: []*FuncArg{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.TrimSpace,
	})
	utilSpace.AddVar("TrimSpace", &VarFunc{
		ScriptName:  "TrimSpace",
		Args: []*FuncArg{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.TrimSpace,
	})

	this_.AddVar("TrimPrefix", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "TrimPrefix",
		Args: []*FuncArg{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.TrimPrefix,
	})
	utilSpace.AddVar("TrimPrefix", &VarFunc{
		ScriptName:  "TrimPrefix",
		Args: []*FuncArg{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.TrimPrefix,
	})

	this_.AddVar("HasPrefix", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "HasPrefix",
		Args: []*FuncArg{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
		callFunc:  util.HasPrefix,
	})
	utilSpace.AddVar("HasPrefix", &VarFunc{
		ScriptName:  "HasPrefix",
		Args: []*FuncArg{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
		callFunc:  util.HasPrefix,
	})

	this_.AddVar("TrimSuffix", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "TrimSuffix",
		Args: []*FuncArg{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.TrimSuffix,
	})
	utilSpace.AddVar("TrimSuffix", &VarFunc{
		ScriptName:  "TrimSuffix",
		Args: []*FuncArg{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.TrimSuffix,
	})

	this_.AddVar("HasSuffix", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "HasSuffix",
		Args: []*FuncArg{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
		callFunc:  util.HasSuffix,
	})
	utilSpace.AddVar("HasSuffix", &VarFunc{
		ScriptName:  "HasSuffix",
		Args: []*FuncArg{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("bool")},
		callFunc:  util.HasSuffix,
	})

	this_.AddVar("TrimLeft", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "TrimLeft",
		Args: []*FuncArg{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.TrimLeft,
	})
	utilSpace.AddVar("TrimLeft", &VarFunc{
		ScriptName:  "TrimLeft",
		Args: []*FuncArg{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.TrimLeft,
	})

	this_.AddVar("TrimRight", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "TrimRight",
		Args: []*FuncArg{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.TrimRight,
	})
	utilSpace.AddVar("TrimRight", &VarFunc{
		ScriptName:  "TrimRight",
		Args: []*FuncArg{
			{Name: "arg", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "trim", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.TrimRight,
	})

	this_.AddVar("StringJoin", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "StringJoin",
		Args: []*FuncArg{
			{Name: "es", Type: parser_tm.NewBindingTypeName("[]string")},
			{Name: "sep", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.StringJoin,
	})
	utilSpace.AddVar("StringJoin", &VarFunc{
		ScriptName:  "StringJoin",
		Args: []*FuncArg{
			{Name: "es", Type: parser_tm.NewBindingTypeName("[]string")},
			{Name: "sep", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.StringJoin,
	})

	this_.AddVar("AnyJoin", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "AnyJoin",
		Args: []*FuncArg{
			{Name: "sep", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "es", Type: parser_tm.NewBindingTypeName("...any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.AnyJoin,
	})
	utilSpace.AddVar("AnyJoin", &VarFunc{
		ScriptName:  "AnyJoin",
		Args: []*FuncArg{
			{Name: "sep", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "es", Type: parser_tm.NewBindingTypeName("...any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.AnyJoin,
	})

	this_.AddVar("IntJoin", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "IntJoin",
		Args: []*FuncArg{
			{Name: "es", Type: parser_tm.NewBindingTypeName("[]int")},
			{Name: "sep", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.IntJoin,
	})
	utilSpace.AddVar("IntJoin", &VarFunc{
		ScriptName:  "IntJoin",
		Args: []*FuncArg{
			{Name: "es", Type: parser_tm.NewBindingTypeName("[]int")},
			{Name: "sep", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.IntJoin,
	})

	this_.AddVar("Int64Join", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "Int64Join",
		Args: []*FuncArg{
			{Name: "es", Type: parser_tm.NewBindingTypeName("[]int64")},
			{Name: "sep", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.Int64Join,
	})
	utilSpace.AddVar("Int64Join", &VarFunc{
		ScriptName:  "Int64Join",
		Args: []*FuncArg{
			{Name: "es", Type: parser_tm.NewBindingTypeName("[]int64")},
			{Name: "sep", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.Int64Join,
	})

	this_.AddVar("GenStringJoin", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GenStringJoin",
		Args: []*FuncArg{
			{Name: "len", Type: parser_tm.NewBindingTypeName("int")},
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "sep", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.GenStringJoin,
	})
	utilSpace.AddVar("GenStringJoin", &VarFunc{
		ScriptName:  "GenStringJoin",
		Args: []*FuncArg{
			{Name: "len", Type: parser_tm.NewBindingTypeName("int")},
			{Name: "str", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "sep", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.GenStringJoin,
	})

	this_.AddVar("GetNow", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetNow",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("time.Time")},
		callFunc:  util.GetNow,
	})
	utilSpace.AddVar("GetNow", &VarFunc{
		ScriptName:  "GetNow",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("time.Time")},
		callFunc:  util.GetNow,
	})

	this_.AddVar("GetNowNano", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetNowNano",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
		callFunc:  util.GetNowNano,
	})
	utilSpace.AddVar("GetNowNano", &VarFunc{
		ScriptName:  "GetNowNano",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
		callFunc:  util.GetNowNano,
	})

	this_.AddVar("GetNowMilli", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetNowMilli",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
		callFunc:  util.GetNowMilli,
	})
	utilSpace.AddVar("GetNowMilli", &VarFunc{
		ScriptName:  "GetNowMilli",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
		callFunc:  util.GetNowMilli,
	})

	this_.AddVar("GetNowSecond", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetNowSecond",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
		callFunc:  util.GetNowSecond,
	})
	utilSpace.AddVar("GetNowSecond", &VarFunc{
		ScriptName:  "GetNowSecond",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
		callFunc:  util.GetNowSecond,
	})

	this_.AddVar("GetNanoByTime", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetNanoByTime",
		Args: []*FuncArg{
			{Name: "v", Type: parser_tm.NewBindingTypeName("time.Time")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
		callFunc:  util.GetNanoByTime,
	})
	utilSpace.AddVar("GetNanoByTime", &VarFunc{
		ScriptName:  "GetNanoByTime",
		Args: []*FuncArg{
			{Name: "v", Type: parser_tm.NewBindingTypeName("time.Time")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
		callFunc:  util.GetNanoByTime,
	})

	this_.AddVar("GetMilliByTime", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetMilliByTime",
		Args: []*FuncArg{
			{Name: "v", Type: parser_tm.NewBindingTypeName("time.Time")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
		callFunc:  util.GetMilliByTime,
	})
	utilSpace.AddVar("GetMilliByTime", &VarFunc{
		ScriptName:  "GetMilliByTime",
		Args: []*FuncArg{
			{Name: "v", Type: parser_tm.NewBindingTypeName("time.Time")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
		callFunc:  util.GetMilliByTime,
	})

	this_.AddVar("GetSecondByTime", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetSecondByTime",
		Args: []*FuncArg{
			{Name: "v", Type: parser_tm.NewBindingTypeName("time.Time")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
		callFunc:  util.GetSecondByTime,
	})
	utilSpace.AddVar("GetSecondByTime", &VarFunc{
		ScriptName:  "GetSecondByTime",
		Args: []*FuncArg{
			{Name: "v", Type: parser_tm.NewBindingTypeName("time.Time")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("int64")},
		callFunc:  util.GetSecondByTime,
	})

	this_.AddVar("GetNowFormat", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetNowFormat",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.GetNowFormat,
	})
	utilSpace.AddVar("GetNowFormat", &VarFunc{
		ScriptName:  "GetNowFormat",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.GetNowFormat,
	})

	this_.AddVar("GetFormatByTime", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetFormatByTime",
		Args: []*FuncArg{
			{Name: "v", Type: parser_tm.NewBindingTypeName("time.Time")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.GetFormatByTime,
	})
	utilSpace.AddVar("GetFormatByTime", &VarFunc{
		ScriptName:  "GetFormatByTime",
		Args: []*FuncArg{
			{Name: "v", Type: parser_tm.NewBindingTypeName("time.Time")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.GetFormatByTime,
	})

	this_.AddVar("TimeFormat", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "TimeFormat",
		Args: []*FuncArg{
			{Name: "v", Type: parser_tm.NewBindingTypeName("time.Time")},
			{Name: "layout", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.TimeFormat,
	})
	utilSpace.AddVar("TimeFormat", &VarFunc{
		ScriptName:  "TimeFormat",
		Args: []*FuncArg{
			{Name: "v", Type: parser_tm.NewBindingTypeName("time.Time")},
			{Name: "layout", Type: parser_tm.NewBindingTypeName("string")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.TimeFormat,
	})

	this_.AddVar("MilliToTimeText", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "MilliToTimeText",
		Args: []*FuncArg{
			{Name: "milli", Type: parser_tm.NewBindingTypeName("int64")},
		},
		Result: &FuncResult{Name: "v", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.MilliToTimeText,
	})
	utilSpace.AddVar("MilliToTimeText", &VarFunc{
		ScriptName:  "MilliToTimeText",
		Args: []*FuncArg{
			{Name: "milli", Type: parser_tm.NewBindingTypeName("int64")},
		},
		Result: &FuncResult{Name: "v", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.MilliToTimeText,
	})

	this_.AddVar("GetUUID", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GetUUID",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.GetUUID,
	})
	utilSpace.AddVar("GetUUID", &VarFunc{
		ScriptName:  "GetUUID",
		Args: []*FuncArg{
			{Name: "", Type: parser_tm.NewBindingTypeName("any")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeName("string")},
		callFunc:  util.GetUUID,
	})

	this_.AddVar("GzipBytes", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "GzipBytes",
		Args: []*FuncArg{
			{Name: "data", Type: parser_tm.NewBindingTypeName("[]byte")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeList("byte")},
		HasError: true,
		callFunc:  util.GzipBytes,
	})
	utilSpace.AddVar("GzipBytes", &VarFunc{
		ScriptName:  "GzipBytes",
		Args: []*FuncArg{
			{Name: "data", Type: parser_tm.NewBindingTypeName("[]byte")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeList("byte")},
		HasError: true,
		callFunc:  util.GzipBytes,
	})

	this_.AddVar("UnGzipBytes", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "UnGzipBytes",
		Args: []*FuncArg{
			{Name: "data", Type: parser_tm.NewBindingTypeName("[]byte")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeList("byte")},
		HasError: true,
		callFunc:  util.UnGzipBytes,
	})
	utilSpace.AddVar("UnGzipBytes", &VarFunc{
		ScriptName:  "UnGzipBytes",
		Args: []*FuncArg{
			{Name: "data", Type: parser_tm.NewBindingTypeName("[]byte")},
		},
		Result: &FuncResult{Name: "res", Type: parser_tm.NewBindingTypeList("byte")},
		HasError: true,
		callFunc:  util.UnGzipBytes,
	})

	this_.AddVar("Zip", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "Zip",
		Args: []*FuncArg{
			{Name: "srcFile", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "destZip", Type: parser_tm.NewBindingTypeName("string")},
		},
		HasError: true,
		callFunc:  util.Zip,
	})
	utilSpace.AddVar("Zip", &VarFunc{
		ScriptName:  "Zip",
		Args: []*FuncArg{
			{Name: "srcFile", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "destZip", Type: parser_tm.NewBindingTypeName("string")},
		},
		HasError: true,
		callFunc:  util.Zip,
	})

	this_.AddVar("UnZip", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "UnZip",
		Args: []*FuncArg{
			{Name: "zipFile", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "destDir", Type: parser_tm.NewBindingTypeName("string")},
		},
		HasError: true,
		callFunc:  util.UnZip,
	})
	utilSpace.AddVar("UnZip", &VarFunc{
		ScriptName:  "UnZip",
		Args: []*FuncArg{
			{Name: "zipFile", Type: parser_tm.NewBindingTypeName("string")},
			{Name: "destDir", Type: parser_tm.NewBindingTypeName("string")},
		},
		HasError: true,
		callFunc:  util.UnZip,
	})

	return
}