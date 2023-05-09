package javascript

import (
	"github.com/team-ide/go-tool/javascript/context_map"
	"github.com/team-ide/go-tool/util"
)

func init() {

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "isEmpty",
		Comment: "是否为nil或空字符串",
		Func:    util.IsEmpty,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "isNotEmpty",
		Comment: "是否不为nil或空字符串",
		Func:    util.IsNotEmpty,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "isNull",
		Comment: "是否为nil",
		Func:    util.IsNull,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "isNotNull",
		Comment: "是否不为nil或空字符串",
		Func:    util.IsNotNull,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "isTrue",
		Comment: "是否为真 判断是true、\"true\"、1、\"1\"",
		Func:    util.IsTrue,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "isFalse",
		Comment: "是否为否 判断不是true、\"true\"、1、\"1\"",
		Func:    util.IsFalse,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "intIndexOf",
		Comment: "返回 某个值 在数组中的索引位置，未找到返回 -1",
		Func:    util.IntIndexOf,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "int64IndexOf",
		Comment: "返回 某个值 在数组中的索引位置，未找到返回 -1",
		Func:    util.Int64IndexOf,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "stringIndexOf",
		Comment: "返回 某个值 在数组中的索引位置，未找到返回 -1",
		Func:    util.StringIndexOf,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "arrayIndexOf",
		Comment: "返回 某个值 在数组中的索引位置，未找到返回 -1",
		Func:    util.ArrayIndexOf,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "setTempDir",
		Comment: "设置临时目录",
		Func:    util.SetTempDir,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "getTempDir",
		Comment: "获取临时目录",
		Func:    util.GetTempDir,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "getRootDir",
		Comment: "获取当前程序根路径",
		Func:    util.GetRootDir,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "formatPath",
		Comment: "格式化路径",
		Func:    util.FormatPath,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "getAbsolutePath",
		Comment: "获取路径觉得路径",
		Func:    util.GetAbsolutePath,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "pathExists",
		Comment: "路径文件是否存在",
		Func:    util.PathExists,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "loadDirFiles",
		Comment: "加载目录下文件 读取文件内容（key为文件名为相对路径）",
		Func:    util.LoadDirFiles,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "loadDirFilenames",
		Comment: "加载目录下文件（文件名为相对路径）",
		Func:    util.LoadDirFilenames,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "readFile",
		Comment: "读取文件内容 返回 []byte",
		Func:    util.ReadFile,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "readFileString",
		Comment: "读取文件内容 返回字符串",
		Func:    util.ReadFileString,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "writeFile",
		Comment: "写入文件内容",
		Func:    util.WriteFile,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "writeFileString",
		Comment: "写入文件内容",
		Func:    util.WriteFileString,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "readLine",
		Comment: "逐行读取文件",
		Func:    util.ReadLine,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "getFileType",
		Comment: "用文件前面几个字节来判断",
		Func:    util.GetFileType,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "getIpFromAddr",
		Comment: "获取当前IP",
		Func:    util.GetIpFromAddr,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "getLocalIPList",
		Comment: "获取当前IP列表",
		Func:    util.GetLocalIPList,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "getLock",
		Comment: "获取一个Locker，如果不存在，则新建",
		Func:    util.GetLock,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "lockByKey",
		Comment: "根据Key进行同步锁",
		Func:    util.LockByKey,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "unlockByKey",
		Comment: "根据Key进行解锁同步锁",
		Func:    util.UnlockByKey,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "getMD5",
		Comment: "获取MD5字符串",
		Func:    util.GetMD5,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "randomInt",
		Comment: "获取随机数",
		Func:    util.RandomInt,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "randomInt64",
		Comment: "获取随机数",
		Func:    util.RandomInt64,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "stringToInt",
		Comment: "字符串转 int",
		Func:    util.StringToInt,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "stringToInt64",
		Comment: "字符串转 int64",
		Func:    util.StringToInt64,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "stringToFloat64",
		Comment: "字符串转 float64",
		Func:    util.StringToFloat64,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "rsaEncryptByKey",
		Comment: "RSA加密",
		Func:    util.RsaEncryptByKey,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "rsaDecryptByKey",
		Comment: "RSA解密",
		Func:    util.RsaDecryptByKey,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "rsaEncrypt",
		Comment: "加密",
		Func:    util.RsaEncrypt,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "rsaDecrypt",
		Comment: "解密",
		Func:    util.RsaDecrypt,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "firstToUpper",
		Comment: "字符首字母大写",
		Func:    util.FirstToUpper,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "firstToLower",
		Comment: "字符首字母小写",
		Func:    util.FirstToLower,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "marshal",
		Comment: "转换为大驼峰命名法则 首字母大写，“_” 忽略后大写",
		Func:    util.Marshal,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "getStringValue",
		Comment: "将传入的值转为字符串",
		Func:    util.GetStringValue,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "toPinYin",
		Comment: "将姓名转为拼音",
		Func:    util.ToPinYin,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "randomString",
		Comment: "获取随机字符串",
		Func:    util.RandomString,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "randomUserName",
		Comment: "随机姓名",
		Func:    util.RandomUserName,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "getNow",
		Comment: "获取当前时间",
		Func:    util.GetNow,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "getNowTime",
		Comment: "获取当前时间戳  到毫秒",
		Func:    util.GetNowTime,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "getNowSecond",
		Comment: "获取当前时间戳 到秒",
		Func:    util.GetNowSecond,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "getTimeByTime",
		Comment: "获取时间戳  到毫秒",
		Func:    util.GetTimeByTime,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "getSecondByTime",
		Comment: "获取时间戳 到秒",
		Func:    util.GetSecondByTime,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "getNowFormat",
		Comment: "获取当前格式化时间 `2006-01-02 15:04:05`",
		Func:    util.GetNowFormat,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "getFormatByTime",
		Comment: "获取格式化时间 `2006-01-02 15:04:05`",
		Func:    util.GetFormatByTime,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "timeFormat",
		Comment: "时间格式化 默认 `2006-01-02 15:04:05`",
		Func:    util.TimeFormat,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "getUUID",
		Comment: "生成UUID",
		Func:    util.GetUUID,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "gzipBytes",
		Comment: "压缩",
		Func:    util.GzipBytes,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "unGzipBytes",
		Comment: "解压",
		Func:    util.UnGzipBytes,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "zip",
		Comment: "zip压缩 srcFile 文件路径，destZip压缩包保存路径",
		Func:    util.Zip,
	})

	context_map.AddFunc(&context_map.FuncInfo{
		Name:    "unZip",
		Comment: "zip解压 zipFile 压缩包地址 destDir 解压保存文件夹",
		Func:    util.UnZip,
	})

}
