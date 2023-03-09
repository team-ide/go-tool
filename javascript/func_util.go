package javascript

import "github.com/team-ide/go-tool/util"

func init() {

	AddFunc(&FuncInfo{
		Name:    "isEmpty",
		Comment: "是否为nil或空字符串",
		Func:    util.IsEmpty,
	})

	AddFunc(&FuncInfo{
		Name:    "isNotEmpty",
		Comment: "是否不为nil或空字符串",
		Func:    util.IsNotEmpty,
	})

	AddFunc(&FuncInfo{
		Name:    "isTrue",
		Comment: "是否为真 判断是true、\"true\"、1、\"1\"",
		Func:    util.IsTrue,
	})

	AddFunc(&FuncInfo{
		Name:    "isFalse",
		Comment: "是否为否 判断不是true、\"true\"、1、\"1\"",
		Func:    util.IsFalse,
	})

	AddFunc(&FuncInfo{
		Name:    "intIndexOf",
		Comment: "返回 某个值 在数组中的索引位置，未找到返回 -1",
		Func:    util.IntIndexOf,
	})

	AddFunc(&FuncInfo{
		Name:    "int64IndexOf",
		Comment: "返回 某个值 在数组中的索引位置，未找到返回 -1",
		Func:    util.Int64IndexOf,
	})

	AddFunc(&FuncInfo{
		Name:    "stringIndexOf",
		Comment: "返回 某个值 在数组中的索引位置，未找到返回 -1",
		Func:    util.StringIndexOf,
	})

	AddFunc(&FuncInfo{
		Name:    "arrayIndexOf",
		Comment: "返回 某个值 在数组中的索引位置，未找到返回 -1",
		Func:    util.ArrayIndexOf,
	})

	AddFunc(&FuncInfo{
		Name:    "getRootDir",
		Comment: "获取当前程序根路径",
		Func:    util.GetRootDir,
	})

	AddFunc(&FuncInfo{
		Name:    "formatPath",
		Comment: "格式化路径",
		Func:    util.FormatPath,
	})

	AddFunc(&FuncInfo{
		Name:    "getAbsolutePath",
		Comment: "获取路径觉得路径",
		Func:    util.GetAbsolutePath,
	})

	AddFunc(&FuncInfo{
		Name:    "pathExists",
		Comment: "路径文件是否存在",
		Func:    util.PathExists,
	})

	AddFunc(&FuncInfo{
		Name:    "loadDirFiles",
		Comment: "加载目录下文件 读取文件内容",
		Func:    util.LoadDirFiles,
	})

	AddFunc(&FuncInfo{
		Name:    "loadDirFilenames",
		Comment: "加载目录下文件",
		Func:    util.LoadDirFilenames,
	})

	AddFunc(&FuncInfo{
		Name:    "readFile",
		Comment: "读取文件内容",
		Func:    util.ReadFile,
	})

	AddFunc(&FuncInfo{
		Name:    "writeFile",
		Comment: "写入文件内容",
		Func:    util.WriteFile,
	})

	AddFunc(&FuncInfo{
		Name:    "readLine",
		Comment: "逐行读取文件",
		Func:    util.ReadLine,
	})

	AddFunc(&FuncInfo{
		Name:    "getFileType",
		Comment: "用文件前面几个字节来判断",
		Func:    util.GetFileType,
	})

	AddFunc(&FuncInfo{
		Name:    "getIpFromAddr",
		Comment: "获取当前IP",
		Func:    util.GetIpFromAddr,
	})

	AddFunc(&FuncInfo{
		Name:    "getLocalIPList",
		Comment: "获取当前IP列表",
		Func:    util.GetLocalIPList,
	})

	AddFunc(&FuncInfo{
		Name:    "getLock",
		Comment: "获取一个Locker，如果不存在，则新建",
		Func:    util.GetLock,
	})

	AddFunc(&FuncInfo{
		Name:    "lockByKey",
		Comment: "根据Key进行同步锁",
		Func:    util.LockByKey,
	})

	AddFunc(&FuncInfo{
		Name:    "unlockByKey",
		Comment: "根据Key进行解锁同步锁",
		Func:    util.UnlockByKey,
	})

	AddFunc(&FuncInfo{
		Name:    "getMD5",
		Comment: "获取MD5字符串",
		Func:    util.GetMD5,
	})

	AddFunc(&FuncInfo{
		Name:    "randomInt",
		Comment: "获取随机数",
		Func:    util.RandomInt,
	})

	AddFunc(&FuncInfo{
		Name:    "randomInt64",
		Comment: "获取随机数",
		Func:    util.RandomInt64,
	})

	AddFunc(&FuncInfo{
		Name:    "rsaEncryptByKey",
		Comment: "RSA加密",
		Func:    util.RsaEncryptByKey,
	})

	AddFunc(&FuncInfo{
		Name:    "rsaDecryptByKey",
		Comment: "RSA解密",
		Func:    util.RsaDecryptByKey,
	})

	AddFunc(&FuncInfo{
		Name:    "rsaEncrypt",
		Comment: "加密",
		Func:    util.RsaEncrypt,
	})

	AddFunc(&FuncInfo{
		Name:    "rsaDecrypt",
		Comment: "解密",
		Func:    util.RsaDecrypt,
	})

	AddFunc(&FuncInfo{
		Name:    "firstToUpper",
		Comment: "字符首字母大写",
		Func:    util.FirstToUpper,
	})

	AddFunc(&FuncInfo{
		Name:    "firstToLower",
		Comment: "字符首字母小写",
		Func:    util.FirstToLower,
	})

	AddFunc(&FuncInfo{
		Name:    "getStringValue",
		Comment: "将传入的值转为字符串",
		Func:    util.GetStringValue,
	})

	AddFunc(&FuncInfo{
		Name:    "toPinYin",
		Comment: "将姓名转为拼音",
		Func:    util.ToPinYin,
	})

	AddFunc(&FuncInfo{
		Name:    "randomString",
		Comment: "获取随机字符串",
		Func:    util.RandomString,
	})

	AddFunc(&FuncInfo{
		Name:    "randomUserName",
		Comment: "随机姓名",
		Func:    util.RandomUserName,
	})

	AddFunc(&FuncInfo{
		Name:    "getNow",
		Comment: "获取当前时间",
		Func:    util.GetNow,
	})

	AddFunc(&FuncInfo{
		Name:    "getNowTime",
		Comment: "获取当前时间戳  到毫秒",
		Func:    util.GetNowTime,
	})

	AddFunc(&FuncInfo{
		Name:    "getNowSecond",
		Comment: "获取当前时间戳 到秒",
		Func:    util.GetNowSecond,
	})

	AddFunc(&FuncInfo{
		Name:    "getTimeByTime",
		Comment: "获取时间戳  到毫秒",
		Func:    util.GetTimeByTime,
	})

	AddFunc(&FuncInfo{
		Name:    "getSecondByTime",
		Comment: "获取时间戳 到秒",
		Func:    util.GetSecondByTime,
	})

	AddFunc(&FuncInfo{
		Name:    "getNowFormat",
		Comment: "获取当前格式化时间 `2006-01-02 15:04:05`",
		Func:    util.GetNowFormat,
	})

	AddFunc(&FuncInfo{
		Name:    "getFormatByTime",
		Comment: "获取格式化时间 `2006-01-02 15:04:05`",
		Func:    util.GetFormatByTime,
	})

	AddFunc(&FuncInfo{
		Name:    "timeFormat",
		Comment: "时间格式化 默认 `2006-01-02 15:04:05`",
		Func:    util.TimeFormat,
	})

	AddFunc(&FuncInfo{
		Name:    "getUUID",
		Comment: "生成UUID",
		Func:    util.GetUUID,
	})

	AddFunc(&FuncInfo{
		Name:    "gzipBytes",
		Comment: "压缩",
		Func:    util.GzipBytes,
	})

	AddFunc(&FuncInfo{
		Name:    "unGzipBytes",
		Comment: "解压",
		Func:    util.UnGzipBytes,
	})

	AddFunc(&FuncInfo{
		Name:    "zip",
		Comment: "zip压缩 srcFile 文件路径，destZip压缩包保存路径",
		Func:    util.Zip,
	})

	AddFunc(&FuncInfo{
		Name:    "unZip",
		Comment: "zip解压 zipFile 压缩包地址 destDir 解压保存文件夹",
		Func:    util.UnZip,
	})

}