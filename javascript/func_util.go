package javascript

import (
	"github.com/team-ide/go-tool/javascript/context_map"
	"github.com/team-ide/go-tool/util"
)

func init() {
	context_map.AddModule(&context_map.ModuleInfo{
		Name:    "util",
		Comment: "工具模块",
		FuncList: []*context_map.FuncInfo{

		{
			Name:    "isEmpty",
			Comment: "是否为nil或空字符串",
			Func:    util.IsEmpty,
		},
		{
			Name:    "isNotEmpty",
			Comment: "是否不为nil或空字符串",
			Func:    util.IsNotEmpty,
		},
		{
			Name:    "isNull",
			Comment: "是否为nil",
			Func:    util.IsNull,
		},
		{
			Name:    "isNotNull",
			Comment: "是否不为nil或空字符串",
			Func:    util.IsNotNull,
		},
		{
			Name:    "isTrue",
			Comment: "是否为真 判断是true、\"true\"、1、\"1\"",
			Func:    util.IsTrue,
		},
		{
			Name:    "isFalse",
			Comment: "是否为否 判断不是true、\"true\"、1、\"1\"",
			Func:    util.IsFalse,
		},
		{
			Name:    "intIndexOf",
			Comment: "返回 某个值 在数组中的索引位置，未找到返回 -1",
			Func:    util.IntIndexOf,
		},
		{
			Name:    "int64IndexOf",
			Comment: "返回 某个值 在数组中的索引位置，未找到返回 -1",
			Func:    util.Int64IndexOf,
		},
		{
			Name:    "stringIndexOf",
			Comment: "返回 某个值 在数组中的索引位置，未找到返回 -1",
			Func:    util.StringIndexOf,
		},
		{
			Name:    "arrayIndexOf",
			Comment: "返回 某个值 在数组中的索引位置，未找到返回 -1",
			Func:    util.ArrayIndexOf,
		},
		{
			Name:    "setTempDir",
			Comment: "设置临时目录",
			Func:    util.SetTempDir,
		},
		{
			Name:    "getTempDir",
			Comment: "获取临时目录",
			Func:    util.GetTempDir,
		},
		{
			Name:    "newWaitGroup",
			Comment: "创建 WaitGroup",
			Func:    util.NewWaitGroup,
		},
		{
			Name:    "newLocker",
			Comment: "创建 Mutex Locker",
			Func:    util.NewLocker,
		},
		{
			Name:    "getRootDir",
			Comment: "获取当前程序根路径",
			Func:    util.GetRootDir,
		},
		{
			Name:    "formatPath",
			Comment: "格式化路径",
			Func:    util.FormatPath,
		},
		{
			Name:    "getAbsolutePath",
			Comment: "获取路径觉得路径",
			Func:    util.GetAbsolutePath,
		},
		{
			Name:    "pathExists",
			Comment: "路径文件是否存在",
			Func:    util.PathExists,
		},
		{
			Name:    "loadDirFiles",
			Comment: "加载目录下文件 读取文件内容（key为文件名为相对路径）",
			Func:    util.LoadDirFiles,
		},
		{
			Name:    "loadDirFilenames",
			Comment: "加载目录下文件（文件名为相对路径）",
			Func:    util.LoadDirFilenames,
		},
		{
			Name:    "readFile",
			Comment: "读取文件内容 返回 []byte",
			Func:    util.ReadFile,
		},
		{
			Name:    "readFileString",
			Comment: "读取文件内容 返回字符串",
			Func:    util.ReadFileString,
		},
		{
			Name:    "writeFile",
			Comment: "写入文件内容",
			Func:    util.WriteFile,
		},
		{
			Name:    "writeFileString",
			Comment: "写入文件内容",
			Func:    util.WriteFileString,
		},
		{
			Name:    "readLine",
			Comment: "逐行读取文件",
			Func:    util.ReadLine,
		},
		{
			Name:    "getFileType",
			Comment: "用文件前面几个字节来判断",
			Func:    util.GetFileType,
		},
		{
			Name:    "getIpFromAddr",
			Comment: "获取当前IP",
			Func:    util.GetIpFromAddr,
		},
		{
			Name:    "getLocalIPList",
			Comment: "获取当前IP列表",
			Func:    util.GetLocalIPList,
		},
		{
			Name:    "objToJson",
			Comment: "对象 转 json 字符串",
			Func:    util.ObjToJson,
		},
		{
			Name:    "jsonToMap",
			Comment: "json 字符串 转 map对象",
			Func:    util.JsonToMap,
		},
		{
			Name:    "jsonToObj",
			Comment: "json 字符串 转 对象",
			Func:    util.JsonToObj,
		},
		{
			Name:    "getLock",
			Comment: "获取一个Locker，如果不存在，则新建",
			Func:    util.GetLock,
		},
		{
			Name:    "lockByKey",
			Comment: "根据Key进行同步锁",
			Func:    util.LockByKey,
		},
		{
			Name:    "unlockByKey",
			Comment: "根据Key进行解锁同步锁",
			Func:    util.UnlockByKey,
		},
		{
			Name:    "getLogger",
			Comment: "获取logger输出对象",
			Func:    util.GetLogger,
		},
		{
			Name:    "getMD5",
			Comment: "获取MD5字符串",
			Func:    util.GetMD5,
		},
		{
			Name:    "randomInt",
			Comment: "获取随机数",
			Func:    util.RandomInt,
		},
		{
			Name:    "randomInt64",
			Comment: "获取随机数",
			Func:    util.RandomInt64,
		},
		{
			Name:    "stringToInt",
			Comment: "字符串转 int",
			Func:    util.StringToInt,
		},
		{
			Name:    "stringToInt64",
			Comment: "字符串转 int64",
			Func:    util.StringToInt64,
		},
		{
			Name:    "stringToFloat64",
			Comment: "字符串转 float64",
			Func:    util.StringToFloat64,
		},
		{
			Name:    "rsaEncryptByKey",
			Comment: "RSA加密",
			Func:    util.RsaEncryptByKey,
		},
		{
			Name:    "rsaDecryptByKey",
			Comment: "RSA解密",
			Func:    util.RsaDecryptByKey,
		},
		{
			Name:    "rsaEncrypt",
			Comment: "加密",
			Func:    util.RsaEncrypt,
		},
		{
			Name:    "rsaDecrypt",
			Comment: "解密",
			Func:    util.RsaDecrypt,
		},
		{
			Name:    "firstToUpper",
			Comment: "字符首字母大写",
			Func:    util.FirstToUpper,
		},
		{
			Name:    "firstToLower",
			Comment: "字符首字母小写",
			Func:    util.FirstToLower,
		},
		{
			Name:    "marshal",
			Comment: "转换为大驼峰命名法则 首字母大写，“_” 忽略后大写",
			Func:    util.Marshal,
		},
		{
			Name:    "getStringValue",
			Comment: "将传入的值转为字符串",
			Func:    util.GetStringValue,
		},
		{
			Name:    "toPinYin",
			Comment: "将姓名转为拼音",
			Func:    util.ToPinYin,
		},
		{
			Name:    "randomString",
			Comment: "获取随机字符串",
			Func:    util.RandomString,
		},
		{
			Name:    "randomUserName",
			Comment: "随机姓名",
			Func:    util.RandomUserName,
		},
		{
			Name:    "strPadLeft",
			Comment: "在字符串 左侧补全 字符串 到 指定长度",
			Func:    util.StrPadLeft,
		},
		{
			Name:    "strPadRight",
			Comment: "在字符串 右侧补全 字符串 到 指定长度",
			Func:    util.StrPadRight,
		},
		{
			Name:    "getNow",
			Comment: "获取当前时间",
			Func:    util.GetNow,
		},
		{
			Name:    "getNowNano",
			Comment: "获取当前时间戳  到纳秒",
			Func:    util.GetNowNano,
		},
		{
			Name:    "getNowMilli",
			Comment: "获取当前时间戳  到毫秒",
			Func:    util.GetNowMilli,
		},
		{
			Name:    "getNowSecond",
			Comment: "获取当前时间戳 到秒",
			Func:    util.GetNowSecond,
		},
		{
			Name:    "getNanoByTime",
			Comment: "获取时间戳  到纳秒",
			Func:    util.GetNanoByTime,
		},
		{
			Name:    "getMilliByTime",
			Comment: "获取时间戳  到毫秒",
			Func:    util.GetMilliByTime,
		},
		{
			Name:    "getSecondByTime",
			Comment: "获取时间戳 到秒",
			Func:    util.GetSecondByTime,
		},
		{
			Name:    "getNowFormat",
			Comment: "获取当前格式化时间 `2006-01-02 15:04:05`",
			Func:    util.GetNowFormat,
		},
		{
			Name:    "getFormatByTime",
			Comment: "获取格式化时间 `2006-01-02 15:04:05`",
			Func:    util.GetFormatByTime,
		},
		{
			Name:    "timeFormat",
			Comment: "时间格式化 默认 `2006-01-02 15:04:05`",
			Func:    util.TimeFormat,
		},
		{
			Name:    "getUUID",
			Comment: "生成UUID",
			Func:    util.GetUUID,
		},
		{
			Name:    "gzipBytes",
			Comment: "压缩",
			Func:    util.GzipBytes,
		},
		{
			Name:    "unGzipBytes",
			Comment: "解压",
			Func:    util.UnGzipBytes,
		},
		{
			Name:    "zip",
			Comment: "zip压缩 srcFile 文件路径，destZip压缩包保存路径",
			Func:    util.Zip,
		},
		{
			Name:    "unZip",
			Comment: "zip解压 zipFile 压缩包地址 destDir 解压保存文件夹",
			Func:    util.UnZip,
		},
		},
	})
}