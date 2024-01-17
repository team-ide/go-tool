package redis

import (
	"fmt"
	"testing"
)

func TestRedis(t *testing.T) {
	config := &Config{
		Address: "192.168.0.85:11080",
		Auth:    "q7ZtCl^5S3",
	}
	service, err := New(config)
	if err != nil {
		panic("redis new error:" + err.Error())
	}

	var key = "online:info:1"
	var deviceType = "1"
	err = service.HashSet(key, deviceType, `{"key":"1","connectID":1,"xx":11}`)
	if err != nil {
		panic("redis HashSet error:" + err.Error())
	}
	var script = `
local key = KEYS[1]
local deviceType = ARGV[2]
local newConnectID = ARGV[1]
local newOnlineInfo = ARGV[2]

local value = redis.call('HGET', key, deviceType)
if value then
    local data = cjson.decode(value)
    if tostring(data.connectID) == newConnectID then
        redis.call('HSET', key, deviceType, newOnlineInfo)
        return newOnlineInfo
    else
        return ""
    end
else
    return ""
end
`
	r1, err := service.HashGet(key, deviceType)
	if err != nil {
		panic("redis HashGet error:" + err.Error())
	}
	fmt.Println("查询信息:", r1)
	r2, err := service.ScriptRun(script, []string{key, deviceType}, []string{"2", "这是更新后的值"})
	if err != nil {
		panic("redis ScriptRun error:" + err.Error())
	}
	fmt.Println("第一次使用错误的信息更新:", r2)
	r3, err := service.HashGet(key, deviceType)
	if err != nil {
		panic("redis HashGet error:" + err.Error())
	}
	fmt.Println("查询信息:", r3)
	r4, err := service.ScriptRun(script, []string{key, deviceType}, []string{"1", "这是更新后的值"})
	if err != nil {
		panic("redis ScriptRun error:" + err.Error())
	}
	fmt.Println("第二次使用正确的信息更新:", r4)
	r5, err := service.HashGet(key, deviceType)
	if err != nil {
		panic("redis HashGet error:" + err.Error())
	}
	fmt.Println("查询信息:", r5)
}
