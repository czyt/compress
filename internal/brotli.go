package internal

import (
	"github.com/czyt/compress"
)

// ref https://github.com/flamego/brotli/blob/main/brotli.go

const (
	brotliHeader = "br"
)

type BrotliCompress struct {
	compressLevel int
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
