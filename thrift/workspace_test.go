package thrift

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestWorkspace(t *testing.T) {
	dir := `C:\Workspaces\Code\thrift`

	workspace := NewWorkspace(dir)

	workspace.Load()
	if len(workspace.errorCache) > 0 {
		for path, err := range workspace.errorCache {
			fmt.Println("path:", path)
			fmt.Println("err:", err)
		}
	}

	filename := workspace.GetFormatDir() + "/test.thrift"

	var args []interface{}

	args = append(args, map[string]interface{}{
		"field1": int8(1),
		"field2": int16(2),
	})
	args = append(args, 2)

	res, err := workspace.InvokeByServerAddress(testServiceAddress, filename, "TestService", "send", args...)
	if err != nil {
		panic(err)
	}

	fmt.Println("Send result:", res)
	bs, _ := json.Marshal(res)
	fmt.Println("Send result JSON:", string(bs))
}

func TestSendMessageByServer(t *testing.T) {
	dir := `D:\Code\linkdood\thrift`

	workspace := NewWorkspace(dir)

	workspace.Load()
	if len(workspace.errorCache) > 0 {
		for path, err := range workspace.errorCache {
			fmt.Println("path:", path)
			fmt.Println("err:", err)
		}
	}

	filename := workspace.GetFormatDir() + "/chat.thrift"

	var args []interface{}

	args = append(args, map[string]interface{}{
		"userID":        "4611686027042922242",
		"targetID":      "4611686029164805890",
		"message":       "这是测试消息",
		"messageType":   2,
		"msgProperties": map[string]interface{}{},
	})
	args = append(args, 2)

	res, err := workspace.InvokeByServerAddress(`192.168.0.85:11203`, filename, "ChatService", "sendMessageByServer", args...)
	if err != nil {
		panic(err)
	}

	fmt.Println("Send result:", res)
	bs, _ := json.Marshal(res)
	fmt.Println("Send result JSON:", string(bs))
}
