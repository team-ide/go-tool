package thrift

import (
	"context"
	"encoding/json"
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
	client, err := NewServiceClientByAddress(testServiceAddress)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = client.TTransport.Close()
	}()
	fmt.Println("service client create success")
	param := &MethodParam{
		Name: "send",
	}
	param.ArgFields = append(param.ArgFields, &Field{
		Name: "res",
		Num:  1,
		Type: &FieldType{
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
					{
						Name: "field2",
						Num:  2,
						Type: &FieldType{
							TypeId: thrift.I16,
						},
					},
				},
			},
		},
	})
	param.Args = append(param.Args, map[string]interface{}{
		"field1": int8(1),
		"field2": int16(2),
	})

	param.ResultType = &FieldType{
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
				{
					Name: "field2",
					Num:  2,
					Type: &FieldType{
						TypeId: thrift.I16,
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
	bs, _ := json.Marshal(res)
	fmt.Println("Send result JSON:", string(bs))
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
