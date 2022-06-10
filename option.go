package compress

type CompressionProvider int

const (
	Gzip CompressionProvider = iota << 1
	Brotli
)

type option struct {
	compressionLevel    CompressionLevel
	compressionProvider CompressionProvider
}

type Option func(opt *option)

func WithCompressionProvider(provider CompressionProvider) Option {
	return func(opt *option) {
		opt.compressionProvider = provider
	}
}

func WithCompressionLevel(level CompressionLevel) Option {
	return func(opt *option) {
		opt.compressionLevel = level
	}
}
