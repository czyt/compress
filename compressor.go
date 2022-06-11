package compress

import (
	"github.com/go-kratos/kratos/v2/middleware"
)

type CompressionLevel int

const (
	NoCompression CompressionLevel = iota
	BetterSpeed
	SystemRecommend
	BestCompression
)

type CompressorServer interface {
	// Server is the Server MiddlewareHandler
	Server() middleware.Middleware
}
type CompressorClient interface {
	// Client is the Client MiddlewareHandler
	Client() middleware.Middleware
}
