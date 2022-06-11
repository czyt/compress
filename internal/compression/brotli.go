package compression

import (
	"context"
	"net/http"
	"strings"

	"github.com/andybalholm/brotli"
	"github.com/czyt/compress"
	"github.com/czyt/compress/internal/compression/internal"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

// ref https://github.com/flamego/brotli/blob/main/brotli.go

const (
	brotliHeader = "br"
)

type BrotliCompress struct {
	compressLevel int
}

type brotliResponseWriter struct {
	h http.ResponseWriter
	b *brotli.Writer
}

func (b *brotliResponseWriter) Header() http.Header {
	return b.h.Header()
}

func (b *brotliResponseWriter) Write(bytes []byte) (int, error) {
	return b.b.Write(bytes)
}

func (b *brotliResponseWriter) WriteHeader(statusCode int) {
	b.h.WriteHeader(statusCode)
}

func (b *BrotliCompress) Server() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				if !strings.Contains(tr.RequestHeader().Get(internal.HeaderAcceptEncoding), "br") {
					return nil, err
				}
				headers := tr.ReplyHeader()
				headers.Set(internal.HeaderContentEncoding, "br")
				headers.Set(internal.HeaderVary, internal.HeaderAcceptEncoding)

				if respWr, ok := tr.(http.ResponseWriter); ok {
					bw := brotli.NewWriterLevel(respWr, b.compressLevel)
					defer func() { _ = bw.Close() }()
					brw := &brotliResponseWriter{
						h: respWr,
						b: bw,
					}
					respWr = brw
					respWr.Header().Del(internal.HeaderContentLength)
					handler(ctx, req)
				}

			}
			return nil, internal.ErrNotAValidRequest
		}
	}

}

func (b *BrotliCompress) Client() middleware.Middleware {
	//TODO implement me
	panic("implement me")
}

func NewBrotliCompress(level compress.CompressionLevel) *BrotliCompress {
	compressionLevel := getBrotliCompressionLevel(level)
	return &BrotliCompress{compressLevel: compressionLevel}
}

func getBrotliCompressionLevel(level compress.CompressionLevel) int {
	switch level {
	case compress.NoCompression:
		return 0
	case compress.BetterSpeed:
		return 1
	case compress.SystemRecommend:
		return 6
	case compress.BestCompression:
		return 11
	default:
		return 5
	}
}
