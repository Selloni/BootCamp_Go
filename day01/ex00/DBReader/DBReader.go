package DBReader

type Interface interface {
	ReadFile(data []byte)
	Write() []byte
	Create(data []byte)
}
