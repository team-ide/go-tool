namespace go test.service

/*
 * thrift -out ./thrift --gen "go:package_prefix=github.com/team-ide/go-tool/thrift/,thrift_import=github.com/apache/thrift/lib/go/thrift" thrift/test.thrift
 */
struct Request {
  1: i8 field1;
  2: i16 field2;
}
struct Response {
  1: i8 field1;
  2: i16 field2;
}
service TestService {
  Response send(1:Request res,2:i8 b)
  binary setUserKey(1: i64 userid, 2: string key, 3: binary value,4: string sessionId)

}