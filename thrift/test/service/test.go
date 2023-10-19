// Code generated by Thrift Compiler (0.14.1). DO NOT EDIT.

package service

import(
	"bytes"
	"context"
	"fmt"
	"time"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

// Attributes:
//  - Field1
//  - Field2
type Request struct {
  Field1 int8 `thrift:"field1,1" db:"field1" json:"field1"`
  Field2 int16 `thrift:"field2,2" db:"field2" json:"field2"`
}

func NewRequest() *Request {
  return &Request{}
}


func (p *Request) GetField1() int8 {
  return p.Field1
}

func (p *Request) GetField2() int16 {
  return p.Field2
}
func (p *Request) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.BYTE {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.I16 {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Request)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadByte(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := int8(v)
  p.Field1 = temp
}
  return nil
}

func (p *Request)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI16(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Field2 = v
}
  return nil
}

func (p *Request) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "Request"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Request) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "field1", thrift.BYTE, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:field1: ", p), err) }
  if err := oprot.WriteByte(ctx, int8(p.Field1)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.field1 (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:field1: ", p), err) }
  return err
}

func (p *Request) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "field2", thrift.I16, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:field2: ", p), err) }
  if err := oprot.WriteI16(ctx, int16(p.Field2)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.field2 (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:field2: ", p), err) }
  return err
}

func (p *Request) Equals(other *Request) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.Field1 != other.Field1 { return false }
  if p.Field2 != other.Field2 { return false }
  return true
}

func (p *Request) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Request(%+v)", *p)
}

// Attributes:
//  - Field1
//  - Field2
type Response struct {
  Field1 int8 `thrift:"field1,1" db:"field1" json:"field1"`
  Field2 int16 `thrift:"field2,2" db:"field2" json:"field2"`
}

func NewResponse() *Response {
  return &Response{}
}


func (p *Response) GetField1() int8 {
  return p.Field1
}

func (p *Response) GetField2() int16 {
  return p.Field2
}
func (p *Response) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.BYTE {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.I16 {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Response)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadByte(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := int8(v)
  p.Field1 = temp
}
  return nil
}

func (p *Response)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI16(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Field2 = v
}
  return nil
}

func (p *Response) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "Response"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Response) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "field1", thrift.BYTE, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:field1: ", p), err) }
  if err := oprot.WriteByte(ctx, int8(p.Field1)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.field1 (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:field1: ", p), err) }
  return err
}

func (p *Response) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "field2", thrift.I16, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:field2: ", p), err) }
  if err := oprot.WriteI16(ctx, int16(p.Field2)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.field2 (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:field2: ", p), err) }
  return err
}

func (p *Response) Equals(other *Response) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.Field1 != other.Field1 { return false }
  if p.Field2 != other.Field2 { return false }
  return true
}

func (p *Response) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Response(%+v)", *p)
}

type TestService interface {
  // Parameters:
  //  - Res
  //  - B
  Send(ctx context.Context, res *Request, b int8) (_r *Response, _err error)
  // Parameters:
  //  - Userid
  //  - Key
  //  - Value
  //  - SessionId
  SetUserKey(ctx context.Context, userid int64, key string, value []byte, sessionId string) (_r []byte, _err error)
}

type TestServiceClient struct {
  c thrift.TClient
  meta thrift.ResponseMeta
}

func NewTestServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *TestServiceClient {
  return &TestServiceClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

func NewTestServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *TestServiceClient {
  return &TestServiceClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewTestServiceClient(c thrift.TClient) *TestServiceClient {
  return &TestServiceClient{
    c: c,
  }
}

func (p *TestServiceClient) Client_() thrift.TClient {
  return p.c
}

func (p *TestServiceClient) LastResponseMeta_() thrift.ResponseMeta {
  return p.meta
}

func (p *TestServiceClient) SetLastResponseMeta_(meta thrift.ResponseMeta) {
  p.meta = meta
}

// Parameters:
//  - Res
//  - B
func (p *TestServiceClient) Send(ctx context.Context, res *Request, b int8) (_r *Response, _err error) {
  var _args0 TestServiceSendArgs
  _args0.Res = res
  _args0.B = b
  var _result2 TestServiceSendResult
  var _meta1 thrift.ResponseMeta
  _meta1, _err = p.Client_().Call(ctx, "send", &_args0, &_result2)
  p.SetLastResponseMeta_(_meta1)
  if _err != nil {
    return
  }
  return _result2.GetSuccess(), nil
}

// Parameters:
//  - Userid
//  - Key
//  - Value
//  - SessionId
func (p *TestServiceClient) SetUserKey(ctx context.Context, userid int64, key string, value []byte, sessionId string) (_r []byte, _err error) {
  var _args3 TestServiceSetUserKeyArgs
  _args3.Userid = userid
  _args3.Key = key
  _args3.Value = value
  _args3.SessionId = sessionId
  var _result5 TestServiceSetUserKeyResult
  var _meta4 thrift.ResponseMeta
  _meta4, _err = p.Client_().Call(ctx, "setUserKey", &_args3, &_result5)
  p.SetLastResponseMeta_(_meta4)
  if _err != nil {
    return
  }
  return _result5.GetSuccess(), nil
}

type TestServiceProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler TestService
}

func (p *TestServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *TestServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *TestServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewTestServiceProcessor(handler TestService) *TestServiceProcessor {

  self6 := &TestServiceProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self6.processorMap["send"] = &testServiceProcessorSend{handler:handler}
  self6.processorMap["setUserKey"] = &testServiceProcessorSetUserKey{handler:handler}
return self6
}

func (p *TestServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err2 := iprot.ReadMessageBegin(ctx)
  if err2 != nil { return false, thrift.WrapTException(err2) }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(ctx, thrift.STRUCT)
  iprot.ReadMessageEnd(ctx)
  x7 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(ctx, name, thrift.EXCEPTION, seqId)
  x7.Write(ctx, oprot)
  oprot.WriteMessageEnd(ctx)
  oprot.Flush(ctx)
  return false, x7

}

type testServiceProcessorSend struct {
  handler TestService
}

func (p *testServiceProcessorSend) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := TestServiceSendArgs{}
  var err2 error
  if err2 = args.Read(ctx, iprot); err2 != nil {
    iprot.ReadMessageEnd(ctx)
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
    oprot.WriteMessageBegin(ctx, "send", thrift.EXCEPTION, seqId)
    x.Write(ctx, oprot)
    oprot.WriteMessageEnd(ctx)
    oprot.Flush(ctx)
    return false, thrift.WrapTException(err2)
  }
  iprot.ReadMessageEnd(ctx)

  tickerCancel := func() {}
  // Start a goroutine to do server side connectivity check.
  if thrift.ServerConnectivityCheckInterval > 0 {
    var cancel context.CancelFunc
    ctx, cancel = context.WithCancel(ctx)
    defer cancel()
    var tickerCtx context.Context
    tickerCtx, tickerCancel = context.WithCancel(context.Background())
    defer tickerCancel()
    go func(ctx context.Context, cancel context.CancelFunc) {
      ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
      defer ticker.Stop()
      for {
        select {
        case <-ctx.Done():
          return
        case <-ticker.C:
          if !iprot.Transport().IsOpen() {
            cancel()
            return
          }
        }
      }
    }(tickerCtx, cancel)
  }

  result := TestServiceSendResult{}
  var retval *Response
  if retval, err2 = p.handler.Send(ctx, args.Res, args.B); err2 != nil {
    tickerCancel()
    if err2 == thrift.ErrAbandonRequest {
      return false, thrift.WrapTException(err2)
    }
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing send: " + err2.Error())
    oprot.WriteMessageBegin(ctx, "send", thrift.EXCEPTION, seqId)
    x.Write(ctx, oprot)
    oprot.WriteMessageEnd(ctx)
    oprot.Flush(ctx)
    return true, thrift.WrapTException(err2)
  } else {
    result.Success = retval
  }
  tickerCancel()
  if err2 = oprot.WriteMessageBegin(ctx, "send", thrift.REPLY, seqId); err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err2 = result.Write(ctx, oprot); err == nil && err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err2 = oprot.WriteMessageEnd(ctx); err == nil && err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err != nil {
    return
  }
  return true, err
}

type testServiceProcessorSetUserKey struct {
  handler TestService
}

func (p *testServiceProcessorSetUserKey) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := TestServiceSetUserKeyArgs{}
  var err2 error
  if err2 = args.Read(ctx, iprot); err2 != nil {
    iprot.ReadMessageEnd(ctx)
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
    oprot.WriteMessageBegin(ctx, "setUserKey", thrift.EXCEPTION, seqId)
    x.Write(ctx, oprot)
    oprot.WriteMessageEnd(ctx)
    oprot.Flush(ctx)
    return false, thrift.WrapTException(err2)
  }
  iprot.ReadMessageEnd(ctx)

  tickerCancel := func() {}
  // Start a goroutine to do server side connectivity check.
  if thrift.ServerConnectivityCheckInterval > 0 {
    var cancel context.CancelFunc
    ctx, cancel = context.WithCancel(ctx)
    defer cancel()
    var tickerCtx context.Context
    tickerCtx, tickerCancel = context.WithCancel(context.Background())
    defer tickerCancel()
    go func(ctx context.Context, cancel context.CancelFunc) {
      ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
      defer ticker.Stop()
      for {
        select {
        case <-ctx.Done():
          return
        case <-ticker.C:
          if !iprot.Transport().IsOpen() {
            cancel()
            return
          }
        }
      }
    }(tickerCtx, cancel)
  }

  result := TestServiceSetUserKeyResult{}
  var retval []byte
  if retval, err2 = p.handler.SetUserKey(ctx, args.Userid, args.Key, args.Value, args.SessionId); err2 != nil {
    tickerCancel()
    if err2 == thrift.ErrAbandonRequest {
      return false, thrift.WrapTException(err2)
    }
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing setUserKey: " + err2.Error())
    oprot.WriteMessageBegin(ctx, "setUserKey", thrift.EXCEPTION, seqId)
    x.Write(ctx, oprot)
    oprot.WriteMessageEnd(ctx)
    oprot.Flush(ctx)
    return true, thrift.WrapTException(err2)
  } else {
    result.Success = retval
  }
  tickerCancel()
  if err2 = oprot.WriteMessageBegin(ctx, "setUserKey", thrift.REPLY, seqId); err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err2 = result.Write(ctx, oprot); err == nil && err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err2 = oprot.WriteMessageEnd(ctx); err == nil && err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Res
//  - B
type TestServiceSendArgs struct {
  Res *Request `thrift:"res,1" db:"res" json:"res"`
  B int8 `thrift:"b,2" db:"b" json:"b"`
}

func NewTestServiceSendArgs() *TestServiceSendArgs {
  return &TestServiceSendArgs{}
}

var TestServiceSendArgs_Res_DEFAULT *Request
func (p *TestServiceSendArgs) GetRes() *Request {
  if !p.IsSetRes() {
    return TestServiceSendArgs_Res_DEFAULT
  }
return p.Res
}

func (p *TestServiceSendArgs) GetB() int8 {
  return p.B
}
func (p *TestServiceSendArgs) IsSetRes() bool {
  return p.Res != nil
}

func (p *TestServiceSendArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.BYTE {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TestServiceSendArgs)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  p.Res = &Request{}
  if err := p.Res.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Res), err)
  }
  return nil
}

func (p *TestServiceSendArgs)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadByte(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  temp := int8(v)
  p.B = temp
}
  return nil
}

func (p *TestServiceSendArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "send_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TestServiceSendArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "res", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:res: ", p), err) }
  if err := p.Res.Write(ctx, oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Res), err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:res: ", p), err) }
  return err
}

func (p *TestServiceSendArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "b", thrift.BYTE, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:b: ", p), err) }
  if err := oprot.WriteByte(ctx, int8(p.B)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.b (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:b: ", p), err) }
  return err
}

func (p *TestServiceSendArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TestServiceSendArgs(%+v)", *p)
}

// Attributes:
//  - Success
type TestServiceSendResult struct {
  Success *Response `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewTestServiceSendResult() *TestServiceSendResult {
  return &TestServiceSendResult{}
}

var TestServiceSendResult_Success_DEFAULT *Response
func (p *TestServiceSendResult) GetSuccess() *Response {
  if !p.IsSetSuccess() {
    return TestServiceSendResult_Success_DEFAULT
  }
return p.Success
}
func (p *TestServiceSendResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *TestServiceSendResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField0(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TestServiceSendResult)  ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
  p.Success = &Response{}
  if err := p.Success.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *TestServiceSendResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "send_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TestServiceSendResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(ctx, oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(ctx); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *TestServiceSendResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TestServiceSendResult(%+v)", *p)
}

// Attributes:
//  - Userid
//  - Key
//  - Value
//  - SessionId
type TestServiceSetUserKeyArgs struct {
  Userid int64 `thrift:"userid,1" db:"userid" json:"userid"`
  Key string `thrift:"key,2" db:"key" json:"key"`
  Value []byte `thrift:"value,3" db:"value" json:"value"`
  SessionId string `thrift:"sessionId,4" db:"sessionId" json:"sessionId"`
}

func NewTestServiceSetUserKeyArgs() *TestServiceSetUserKeyArgs {
  return &TestServiceSetUserKeyArgs{}
}


func (p *TestServiceSetUserKeyArgs) GetUserid() int64 {
  return p.Userid
}

func (p *TestServiceSetUserKeyArgs) GetKey() string {
  return p.Key
}

func (p *TestServiceSetUserKeyArgs) GetValue() []byte {
  return p.Value
}

func (p *TestServiceSetUserKeyArgs) GetSessionId() string {
  return p.SessionId
}
func (p *TestServiceSetUserKeyArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField3(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 4:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField4(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TestServiceSetUserKeyArgs)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Userid = v
}
  return nil
}

func (p *TestServiceSetUserKeyArgs)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Key = v
}
  return nil
}

func (p *TestServiceSetUserKeyArgs)  ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBinary(ctx); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.Value = v
}
  return nil
}

func (p *TestServiceSetUserKeyArgs)  ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.SessionId = v
}
  return nil
}

func (p *TestServiceSetUserKeyArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "setUserKey_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
    if err := p.writeField3(ctx, oprot); err != nil { return err }
    if err := p.writeField4(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TestServiceSetUserKeyArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "userid", thrift.I64, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:userid: ", p), err) }
  if err := oprot.WriteI64(ctx, int64(p.Userid)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.userid (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:userid: ", p), err) }
  return err
}

func (p *TestServiceSetUserKeyArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "key", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:key: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.Key)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.key (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:key: ", p), err) }
  return err
}

func (p *TestServiceSetUserKeyArgs) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "value", thrift.STRING, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:value: ", p), err) }
  if err := oprot.WriteBinary(ctx, p.Value); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.value (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:value: ", p), err) }
  return err
}

func (p *TestServiceSetUserKeyArgs) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "sessionId", thrift.STRING, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:sessionId: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.SessionId)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.sessionId (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:sessionId: ", p), err) }
  return err
}

func (p *TestServiceSetUserKeyArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TestServiceSetUserKeyArgs(%+v)", *p)
}

// Attributes:
//  - Success
type TestServiceSetUserKeyResult struct {
  Success []byte `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewTestServiceSetUserKeyResult() *TestServiceSetUserKeyResult {
  return &TestServiceSetUserKeyResult{}
}

var TestServiceSetUserKeyResult_Success_DEFAULT []byte

func (p *TestServiceSetUserKeyResult) GetSuccess() []byte {
  return p.Success
}
func (p *TestServiceSetUserKeyResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *TestServiceSetUserKeyResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField0(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TestServiceSetUserKeyResult)  ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBinary(ctx); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = v
}
  return nil
}

func (p *TestServiceSetUserKeyResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "setUserKey_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TestServiceSetUserKeyResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRING, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteBinary(ctx, p.Success); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(ctx); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *TestServiceSetUserKeyResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TestServiceSetUserKeyResult(%+v)", *p)
}


