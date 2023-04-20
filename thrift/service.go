package thrift

import "context"

type Service interface {
	Send(ctx context.Context, p *Struct) (_r *Struct, _err error)
}
