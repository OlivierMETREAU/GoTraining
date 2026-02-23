package compressor

type Compressor interface {
	Compress(input []byte) ([]byte, error)
	Decompress(input []byte) ([]byte, error)
	Name() string
}
