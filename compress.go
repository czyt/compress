package compress

import (
	"github.com/czyt/compress/internal/compression"
	"github.com/go-kratos/kratos/v2/middleware"
)

func Server(opts ...Option) middleware.Middleware {
	o := &option{}
	for _, opt := range opts {
		opt(o)
	}
	var compressorServer CompressorServer
	switch o.compressionProvider {
	case Gzip:
		compressorServer = compression.NewGzipCompress(o.compressionLevel)
	case Brotli:
		compressorServer = compression.NewBrotliCompress(o.compressionLevel)
	}
	return compressorServer.Server()
}
