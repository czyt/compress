package compress

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
)

const (
	headerAcceptEncoding  = "Accept-Encoding"
	headerContentEncoding = "Content-Encoding"
	headerVary            = "Vary"
)

func Server(opts ...Option) middleware.Middleware {
	o := &option{}
	for _, opt := range opts {
		opt(o)
	}
	switch o.compressionProvider {
	case Gzip:
	case Brotli:
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {

			return nil, nil
		}
	}
}
