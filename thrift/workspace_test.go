package thrift

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/team-ide/go-tool/util"
	"reflect"
	"testing"
	"time"
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
	s, _ := util.ObjToJson(res)
	fmt.Println("Send result JSON:", s)

	ws2 := NewWorkspace(dir)

	ws2.Load()
	ws3 := NewWorkspace(dir)

	ws3.Load()
	time.Sleep(time.Second * 10)
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
		"userID":        4611686027042922242,
		"targetID":      4611686029164805890,
		"message":       "这是测试消息",
		"messageType":   2,
		"msgProperties": map[string]interface{}{},
	})
	args = append(args, 2)
	w := bytes.NewBufferString("")
	d := json.NewEncoder(w)
	_ = d.Encode(args)
	fmt.Println(w.String())

	str := `{"message":"这是测试消息","messageType":2,"msgProperties":{},"targetID":4611686027042922242.1,"userID":4611686029164805890}`
	data := map[string]interface{}{}
	_ = util.JSONDecodeUseNumber([]byte(str), &data)

	fmt.Println(data)
	fmt.Println(data["targetID"])
	fmt.Println(reflect.TypeOf(data["targetID"]).String())
	fmt.Println(data["targetID"].(json.Number).String())

	bs, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(bs))

	res, err := workspace.InvokeByServerAddress(`192.168.0.85:11203`, filename, "ChatService", "sendMessageByServer", args...)
	if err != nil {
		panic(err)
	}

	fmt.Println("Send result:", res)

	bs, _ = json.MarshalIndent(res, "", "  ")
	fmt.Println("Send result JSON:", string(bs))
}

func TestListIds(t *testing.T) {
	dir := `D:\Code\linkdood\thrift`

	workspace := NewWorkspace(dir)

	workspace.Load()
	if len(workspace.errorCache) > 0 {
		for path, err := range workspace.errorCache {
			fmt.Println("path:", path)
			fmt.Println("err:", err)
		}
	}

	filename := workspace.GetFormatDir() + "/idgenerator.thrift"

	var args []interface{}

	args = append(args, 1)
	args = append(args, 10)
	res, err := workspace.InvokeByServerAddress(`192.168.0.85:11251`, filename, "IdGeneratorService", "listIds", args...)
	if err != nil {
		panic(err)
	}

	fmt.Println("Send result:", res)

	bs, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println("Send result JSON:", string(bs))
}
