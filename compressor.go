package compress

type CompressionLevel int

const (
	NoCompression CompressionLevel = iota
	BetterSpeed
	SystemRecommend
	BestCompression
)

type Compressor interface {
	//TODO:Define interface
}
