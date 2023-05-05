package thrift

import (
	"context"
	"errors"
	"github.com/apache/thrift/lib/go/thrift"
)

type Service interface {
	Send(ctx context.Context, param *MethodParam) (result interface{}, err error)
}

func NewServiceClientByAddress(address string) (client *ServiceClient, err error) {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport := thrift.NewTSocketConf(address, &thrift.TConfiguration{})
	useTransport, err := transportFactory.GetTransport(transport)
	if err != nil {
		err = errors.New("transportFactory.GetTransport error:" + err.Error())
		return
	}
	client = NewServiceClientFactory(useTransport, protocolFactory)
	if err = transport.Open(); err != nil {
		err = errors.New("opening socket to " + address + " error:" + err.Error())
		return
	}
	return
}

type ServiceClient struct {
	TTransport thrift.TTransport
	TClient    thrift.TClient
	meta       thrift.ResponseMeta
}

func NewServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *ServiceClient {
	return &ServiceClient{
		TTransport: t,
		TClient:    thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewServiceClientProtocol(t thrift.TTransport, inProtocol thrift.TProtocol, outProtocol thrift.TProtocol) *ServiceClient {
	return &ServiceClient{
		TTransport: t,
		TClient:    thrift.NewTStandardClient(inProtocol, outProtocol),
	}
}

func NewServiceClient(c thrift.TClient) *ServiceClient {
	return &ServiceClient{
		TClient: c,
	}
}

func (this_ *ServiceClient) Client_() thrift.TClient {
	return this_.TClient
}

func (this_ *ServiceClient) Stop() {
	if this_ != nil && this_.TTransport != nil {
		_ = this_.TTransport.Close()
	}
}

func (this_ *ServiceClient) LastResponseMeta_() thrift.ResponseMeta {
	return this_.meta
}

func (this_ *ServiceClient) SetLastResponseMeta_(meta thrift.ResponseMeta) {
	this_.meta = meta
}

func (this_ *ServiceClient) Send(ctx context.Context, param *MethodParam) (result interface{}, err error) {
	var _meta1 thrift.ResponseMeta
	_meta1, err = this_.Client_().Call(ctx, param.Name, param, param)
	this_.SetLastResponseMeta_(_meta1)
	if err != nil {
		param.Error = err.Error()
		return
	}
	return param.Result, nil
}
