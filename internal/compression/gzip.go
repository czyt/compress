package compression

import (
	"compress/gzip"
	"context"
	"net/http"
	"strings"

	"github.com/czyt/compress"
	"github.com/czyt/compress/internal/compression/internal"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

// ref https://github.com/flamego/gzip/blob/main/gzip.go
const (
	gzipHeader = "gzip"
)

type GzipCompress struct {
	compressLevel int
}

type gzipResponseWriter struct {
	h http.ResponseWriter
	b *gzip.Writer
}

func (g *gzipResponseWriter) Header() http.Header {
	return g.h.Header()
}

func (g *gzipResponseWriter) Write(bytes []byte) (int, error) {
	return g.b.Write(bytes)
}

func (g *gzipResponseWriter) WriteHeader(statusCode int) {
	g.h.WriteHeader(statusCode)
}

func (g GzipCompress) Server() middleware.Middleware {
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
					gw, _ := gzip.NewWriterLevel(respWr, g.compressLevel)
					defer func() { _ = gw.Close() }()
					brw := &gzipResponseWriter{
						h: respWr,
						b: gw,
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

func (g GzipCompress) Client() middleware.Middleware {
	//TODO implement me
	panic("implement me")
}

func NewGzipCompress(level compress.CompressionLevel) *GzipCompress {
	compressionLevel := getGzipCompressionLevel(level)
	return &GzipCompress{compressLevel: compressionLevel}
}

func getGzipCompressionLevel(level compress.CompressionLevel) int {
	switch level {
	case compress.NoCompression:
		return 0
	case compress.BetterSpeed:
		return 1
	case compress.SystemRecommend:
		return 4
	case compress.BestCompression:
		return 9
	default:
		return 4
	}
}
