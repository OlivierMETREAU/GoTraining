package iomanager

type IOManager interface {
	ReadLines() ([]byte, error)
	WriteResult(data any) error
}
