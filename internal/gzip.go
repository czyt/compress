package internal

import (
	"github.com/czyt/compress"
)

// ref https://github.com/flamego/gzip/blob/main/gzip.go
const (
	gzipHeader = "gzip"
)

type GzipCompress struct {
	compressLevel int
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
