package thrift

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/team-ide/go-tool/thrift/test/service"
	"github.com/team-ide/go-tool/util"
	"testing"
	"time"
)

var (
	testServiceAddress = `:10001`
)

func TestServiceClient(t *testing.T) {
	//client, err := NewServiceClientByAddress("192.168.6.152:11209")
	client, err := NewServiceClientByAddress("172.16.8.158:11209")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = client.TTransport.Close()
	}()
	fmt.Println("service client create success")
	param := &MethodParam{
		Name: "queryEmoticonPackListByLabel",
	}
	param.ArgFields = append(param.ArgFields, &Field{
		Name: "label",
		Num:  1,
		Type: &FieldType{
			TypeId: thrift.STRING,
		},
	})
	param.Args = append(param.Args, map[string]interface{}{
		"label": "111",
	})

	param.ResultType = &FieldType{
		TypeId: thrift.LIST,
		ListType: &FieldType{
			TypeId: thrift.STRUCT,
			structObj: &Struct{
				Fields: []*Field{
					{
						Name: "field1",
						Num:  1,
						Type: &FieldType{
							TypeId: thrift.I08,
						},
					},
				},
			},
		},
	}

	res, err := client.Send(context.Background(), param)
	if err != nil {
		panic(err)
	}
	fmt.Println("Send result:", res)
	s, _ := util.ObjToJson(res)
	fmt.Println("Send result JSON:", s)
	time.Sleep(time.Second * 10)

}

func TestServiceClientSetUserKey(t *testing.T) {
	client, err := NewServiceClientByAddress(testServiceAddress)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = client.TTransport.Close()
	}()
	fmt.Println("service client create success")
	param := &MethodParam{
		Name: "setUserKey",
	}
	param.ArgFields = append(param.ArgFields, &Field{
		Name: "userid",
		Num:  1,
		Type: &FieldType{
			TypeId: thrift.I64,
		},
	})
	param.ArgFields = append(param.ArgFields, &Field{
		Name: "key",
		Num:  2,
		Type: &FieldType{
			TypeId: thrift.STRING,
		},
	})
	param.ArgFields = append(param.ArgFields, &Field{
		Name: "value",
		Num:  3,
		Type: &FieldType{
			TypeId: BINARY,
		},
	})
	param.ArgFields = append(param.ArgFields, &Field{
		Name: "sessionId",
		Num:  4,
		Type: &FieldType{
			TypeId: thrift.STRING,
		},
	})
	param.Args = append(param.Args, 1, "2", "3", "4")

	param.ResultType = &FieldType{
		TypeId: BINARY,
	}

	res, err := client.Send(context.Background(), param)
	if err != nil {
		panic(err)
	}
	fmt.Println("Send result:", res)
	s, _ := util.ObjToJson(res)
	fmt.Println("Send result JSON:", s)
	time.Sleep(time.Second * 10)

}

type TestServiceImpl struct {
	thrift.TProcessor
}

func NewTestServiceImpl() *TestServiceImpl {

	return &TestServiceImpl{}
}

func (this_ *TestServiceImpl) Send(ctx context.Context, res *service.Request, b int8) (_r *service.Response, _err error) {
	if res == nil {
		return
	}
	_r = &service.Response{}
	//fmt.Println("Server On Send res:", toJSON(res), ",b:", b)
	_r.Field1 = res.Field1 + int8(util.RandomInt(100, 555))
	_r.Field2 = res.Field2 + int16(util.RandomInt(100, 555))
	time.Sleep(time.Millisecond * 1)
	//fmt.Println("Server On Send _r:", toJSON(_r))
	return
}

func (this_ *TestServiceImpl) SetUserKey(ctx context.Context, userid int64, key string, value []byte, sessionId string) (_r []byte, _err error) {
	fmt.Println("userid:", userid)
	fmt.Println("key:", key)
	fmt.Println("value:", value)
	fmt.Println("sessionId:", sessionId)
	_r = []byte("SetUserKey Success")
	return
}

func TestServiceServer(t *testing.T) {
	handler := NewTestServiceImpl()
	processor := service.NewTestServiceProcessor(handler)
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	serverTransport, err := thrift.NewTServerSocket(testServiceAddress)
	if err != nil {
		fmt.Println("NewTServerSocket error", err)
		return
	}

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server start...")
	err = server.Serve()
	if err != nil {
		fmt.Println("server Serve error", err)
		return
	}
}
