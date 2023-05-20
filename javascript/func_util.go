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
			Name:    "aesEncryptCBCByKey",
			Comment: `AES加密,CBC模式，返回 base64 字符
AesEncryptCBCByKey(\"这是需要加密的文本\", \"这是密钥\")`,
			Func:    util.AesEncryptCBCByKey,
		},
		{
			Name:    "aesDecryptCBCByKey",
			Comment: `AES解密,CBC模式
AesDecryptCBCByKey(\"这是加密后的文本\", \"这是密钥\")`,
			Func:    util.AesDecryptCBCByKey,
		},
		{
			Name:    "aesEncryptECBByKey",
			Comment: `AES加密,ECB模式，返回 base64 字符
AesEncryptECBByKey(\"这是需要加密的文本\", \"这是密钥\")`,
			Func:    util.AesEncryptECBByKey,
		},
		{
			Name:    "aesDecryptECBByKey",
			Comment: `AES解密,ECB模式
AesDecryptECBByKey(\"这是加密后的文本\", \"这是密钥\")`,
			Func:    util.AesDecryptECBByKey,
		},
		{
			Name:    "isEmpty",
			Comment: `是否为nil或空字符串
@param v interface{} \"传入任意值\"
@return bool
IsEmpty(arg)`,
			Func:    util.IsEmpty,
		},
		{
			Name:    "isNotEmpty",
			Comment: `是否不为nil或空字符串
@param v interface{} \"传入任意值\"
@return bool
IsNotEmpty(arg)`,
			Func:    util.IsNotEmpty,
		},
		{
			Name:    "isNull",
			Comment: `是否为nil
@param v interface{} \"传入任意值\"
@return bool
IsNull(arg)`,
			Func:    util.IsNull,
		},
		{
			Name:    "isNotNull",
			Comment: `是否不为nil或空字符串
@param v interface{} \"传入任意值\"
@return bool
IsNotNull(arg)`,
			Func:    util.IsNotNull,
		},
		{
			Name:    "isTrue",
			Comment: `是否为真 判断是true、\"true\"、1、\"1\"
@param v interface{} \"传入任意值\"
@return bool
IsTrue(arg)`,
			Func:    util.IsTrue,
		},
		{
			Name:    "isFalse",
			Comment: `是否为否 判断不是true、\"true\"、1、\"1\"
@param v interface{} \"传入任意值\"
@return bool
IsFalse(arg)`,
			Func:    util.IsFalse,
		},
		{
			Name:    "intIndexOf",
			Comment: `返回 某个值 在数组中的索引位置，未找到返回 -1
IntIndexOf([1,2,3], 2)`,
			Func:    util.IntIndexOf,
		},
		{
			Name:    "int64IndexOf",
			Comment: `返回 某个值 在数组中的索引位置，未找到返回 -1
Int64IndexOf([1,2,3], 2)`,
			Func:    util.Int64IndexOf,
		},
		{
			Name:    "stringIndexOf",
			Comment: `返回 某个值 在数组中的索引位置，未找到返回 -1
StringIndexOf([\"a\", \"b\", \"c\"], \"d\")`,
			Func:    util.StringIndexOf,
		},
		{
			Name:    "arrayIndexOf",
			Comment: `返回 某个值 在数组中的索引位置，未找到返回 -1
ArrayIndexOf([\"a\", \"b\", \"c\"], \"d\")`,
			Func:    util.ArrayIndexOf,
		},
		{
			Name:    "getTempDir",
			Comment: `获取临时目录`,
			Func:    util.GetTempDir,
		},
		{
			Name:    "newWaitGroup",
			Comment: `创建 WaitGroup ，
obj = NewWaitGroup()
obj.Add(1)
obj.Done()
obj.Wait()`,
			Func:    util.NewWaitGroup,
		},
		{
			Name:    "newLocker",
			Comment: `创建 Mutex Locker
obj = NewLocker()
obj.Lock()
obj.Unlock()`,
			Func:    util.NewLocker,
		},
		{
			Name:    "getRootDir",
			Comment: `获取当前程序根路径`,
			Func:    util.GetRootDir,
		},
		{
			Name:    "formatPath",
			Comment: `格式化路径
FormatPath(\"/x/x/xxx\xx\xx\")`,
			Func:    util.FormatPath,
		},
		{
			Name:    "getAbsolutePath",
			Comment: `获取路径觉得路径
GetAbsolutePath(\"/x/x/xxx\xx\xx\")`,
			Func:    util.GetAbsolutePath,
		},
		{
			Name:    "pathExists",
			Comment: `路径文件是否存在
PathExists(\"/x/x/xxx\xx\xx\")`,
			Func:    util.PathExists,
		},
		{
			Name:    "loadDirFiles",
			Comment: `加载目录下文件 读取文件内容（key为文件名为相对路径）
LoadDirFiles(\"/x/x/xxx\xx\xx\")`,
			Func:    util.LoadDirFiles,
		},
		{
			Name:    "loadDirFilenames",
			Comment: `加载目录下文件（文件名为相对路径）
LoadDirFilenames(\"/x/x/xxx\xx\xx\")`,
			Func:    util.LoadDirFilenames,
		},
		{
			Name:    "readFile",
			Comment: `读取文件内容 返回 []byte
ReadFile(\"/x/x/xxx\xx\xx\")`,
			Func:    util.ReadFile,
		},
		{
			Name:    "readFileString",
			Comment: `读取文件内容 返回字符串
ReadFileString(\"/x/x/xxx\xx\xx\")`,
			Func:    util.ReadFileString,
		},
		{
			Name:    "stringToBytes",
			Comment: `字符串转为 []byte
StringToBytes(\"这是文本\")`,
			Func:    util.StringToBytes,
		},
		{
			Name:    "writeFile",
			Comment: `写入文件内容,
WriteFile(\"/x/x/xxx\xx\xx\", StringToBytes(\"这是文本\"))`,
			Func:    util.WriteFile,
		},
		{
			Name:    "writeFileString",
			Comment: `写入文件内容
WriteFileString(\"/x/x/xxx\xx\xx\", \"这是文本\")`,
			Func:    util.WriteFileString,
		},
		{
			Name:    "readLine",
			Comment: `逐行读取文件
ReadLine(\"/x/x/xxx\xx\xx\")`,
			Func:    util.ReadLine,
		},
		{
			Name:    "getFileType",
			Comment: `用文件前面几个字节来判断
fSrc: 文件字节流（就用前面几个字节）`,
			Func:    util.GetFileType,
		},
		{
			Name:    "getIpFromAddr",
			Comment: `获取当前IP`,
			Func:    util.GetIpFromAddr,
		},
		{
			Name:    "getLocalIPList",
			Comment: `获取当前IP列表`,
			Func:    util.GetLocalIPList,
		},
		{
			Name:    "objToJson",
			Comment: `对象 转 json 字符串
ObjToJson(obj)`,
			Func:    util.ObjToJson,
		},
		{
			Name:    "jsonToMap",
			Comment: `json 字符串 转 map对象
JsonToMap(\"{\\"a\\":1}\")`,
			Func:    util.JsonToMap,
		},
		{
			Name:    "jsonToObj",
			Comment: `json 字符串 转 对象
JsonToObj(\"{\\"a\\":1}\", &obj)`,
			Func:    util.JsonToObj,
		},
		{
			Name:    "getLock",
			Comment: `获取一个Locker，如果不存在，则新建
obj = GetLock(\"user:1\")
obj.Lock()
obj.Unlock()`,
			Func:    util.GetLock,
		},
		{
			Name:    "lockByKey",
			Comment: `根据Key进行同步锁
LockByKey(\"user:1\")`,
			Func:    util.LockByKey,
		},
		{
			Name:    "unlockByKey",
			Comment: `根据Key进行解锁同步锁
UnlockByKey(\"user:1\")`,
			Func:    util.UnlockByKey,
		},
		{
			Name:    "getLogger",
			Comment: `获取logger输出对象`,
			Func:    util.GetLogger,
		},
		{
			Name:    "getMD5",
			Comment: `获取MD5字符串
@param str string \"需要MD5的字符串\"
GetMD5(\"这是需要MD5的文本\")`,
			Func:    util.GetMD5,
		},
		{
			Name:    "randomInt",
			Comment: `获取随机数
@param min int \"最小值\"
@param max int \"最大值\"
@return int \"随机数\"
RandomInt(1, 10)`,
			Func:    util.RandomInt,
		},
		{
			Name:    "randomInt64",
			Comment: `获取随机数
@param min int64 \"最小值\"
@param max int64 \"最大值\"
@return int64 \"随机数\"
RandomInt64(1, 10)`,
			Func:    util.RandomInt64,
		},
		{
			Name:    "stringToInt",
			Comment: `字符串转 int
StringToInt(\"11\")`,
			Func:    util.StringToInt,
		},
		{
			Name:    "stringToInt64",
			Comment: `字符串转 int64
StringToInt64(\"11\")`,
			Func:    util.StringToInt64,
		},
		{
			Name:    "stringToFloat64",
			Comment: `字符串转 float64
StringToFloat64(\"11.2\")`,
			Func:    util.StringToFloat64,
		},
		{
			Name:    "rsaEncryptByKey",
			Comment: `RSA加密，返回 base64 字符
RsaEncryptByKey(\"这是需要加密的文本\", \"这是密钥\")`,
			Func:    util.RsaEncryptByKey,
		},
		{
			Name:    "rsaDecryptByKey",
			Comment: `RSA解密
RsaDecryptByKey(\"这是加密后的文本\", \"这是密钥\")`,
			Func:    util.RsaDecryptByKey,
		},
		{
			Name:    "firstToUpper",
			Comment: `字符首字母大写
@param str string \"任意字符串\"
@return string
FirstToUpper(\"abc\")`,
			Func:    util.FirstToUpper,
		},
		{
			Name:    "firstToLower",
			Comment: `字符首字母小写
@param str string \"任意字符串\"
@return string
FirstToLower(\"Abc\")`,
			Func:    util.FirstToLower,
		},
		{
			Name:    "marshal",
			Comment: `转换为大驼峰命名法则 首字母大写，“_” 忽略后大写
Marshal(\"abc_def\")`,
			Func:    util.Marshal,
		},
		{
			Name:    "getStringValue",
			Comment: `将传入的值转为字符串
@param value interface{} \"任意值\"
@return string
GetStringValue(arg)`,
			Func:    util.GetStringValue,
		},
		{
			Name:    "toPinYin",
			Comment: `将姓名转为拼音
@param name string \"姓名\"
@return string
ToPinYin(\"张三\")`,
			Func:    util.ToPinYin,
		},
		{
			Name:    "randomString",
			Comment: `获取随机字符串
@param minLen int \"最小长度\"
@param maxLen int \"最大长度\"
@return string
RandomString(2, 20)`,
			Func:    util.RandomString,
		},
		{
			Name:    "randomUserName",
			Comment: `随机姓名
@param size int \"名长度\"
@return string
RandomUserName(2)`,
			Func:    util.RandomUserName,
		},
		{
			Name:    "strPadLeft",
			Comment: `在字符串 左侧补全 字符串 到 指定长度
input string 原字符串
padLength int 规定补齐后的字符串长度
padString string 自定义填充字符串
StrPadLeft(\"xx\", 5, \"0\") 左侧补”0“达到5位长度`,
			Func:    util.StrPadLeft,
		},
		{
			Name:    "strPadRight",
			Comment: `在字符串 右侧补全 字符串 到 指定长度
input string 原字符串
padLength int 规定补齐后的字符串长度
padString string 自定义填充字符串
StrPadRight(\"xx\", 5, \"0\") 右侧补”0“达到5位长度`,
			Func:    util.StrPadRight,
		},
		{
			Name:    "getNow",
			Comment: `获取当前时间
GetNow()`,
			Func:    util.GetNow,
		},
		{
			Name:    "getNowNano",
			Comment: `获取当前时间戳  到纳秒
GetNowNano()`,
			Func:    util.GetNowNano,
		},
		{
			Name:    "getNowMilli",
			Comment: `获取当前时间戳  到毫秒
GetNowMilli()`,
			Func:    util.GetNowMilli,
		},
		{
			Name:    "getNowSecond",
			Comment: `获取当前时间戳 到秒
GetNowSecond()`,
			Func:    util.GetNowSecond,
		},
		{
			Name:    "getNanoByTime",
			Comment: `获取时间戳  到纳秒
@param v time.Time \"时间\"
GetNanoByTime(time)`,
			Func:    util.GetNanoByTime,
		},
		{
			Name:    "getMilliByTime",
			Comment: `获取时间戳  到毫秒
@param v time.Time \"时间\"
GetMilliByTime(time)`,
			Func:    util.GetMilliByTime,
		},
		{
			Name:    "getSecondByTime",
			Comment: `获取时间戳 到秒
@param v time.Time \"时间\"
GetSecondByTime(time)`,
			Func:    util.GetSecondByTime,
		},
		{
			Name:    "getNowFormat",
			Comment: `获取当前格式化时间 \"2006-01-02 15:04:05\"
GetNowFormat()`,
			Func:    util.GetNowFormat,
		},
		{
			Name:    "getFormatByTime",
			Comment: `获取格式化时间 \"2006-01-02 15:04:05\"
@param v time.Time \"时间\"
GetFormatByTime(time)`,
			Func:    util.GetFormatByTime,
		},
		{
			Name:    "timeFormat",
			Comment: `时间格式化 默认 \"2006-01-02 15:04:05\"
@param v time.Time \"时间\"
@param layout string \"格式化字符串，默认使用\"2006-01-02 15:04:05\"\"
TimeFormat(time, \"2006-01-02 15:04:05\")`,
			Func:    util.TimeFormat,
		},
		{
			Name:    "getUUID",
			Comment: `生成UUID
GetUUID()`,
			Func:    util.GetUUID,
		},
		{
			Name:    "gzipBytes",
			Comment: `压缩`,
			Func:    util.GzipBytes,
		},
		{
			Name:    "unGzipBytes",
			Comment: `解压`,
			Func:    util.UnGzipBytes,
		},
		{
			Name:    "zip",
			Comment: `zip压缩 srcFile 文件路径，destZip压缩包保存路径`,
			Func:    util.Zip,
		},
		{
			Name:    "unZip",
			Comment: `zip解压 zipFile 压缩包地址 destDir 解压保存文件夹`,
			Func:    util.UnZip,
		},
		},
	})
}