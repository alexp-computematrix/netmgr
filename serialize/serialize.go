package serialize

type NetConfigSerializer interface {
	Serialize() ([]byte, error)
	Deserialize([]byte) error
	Bytes() []byte
	String() string
}

type BaseNetConfigSerializer struct {
	NetConfigSerializer
}

func (s *BaseNetConfigSerializer) Serialize() ([]byte, error) {
	panic("not implemented")
}

func (s *BaseNetConfigSerializer) Deserialize(data []byte) error {
	panic("not implemented")
}

func (s *BaseNetConfigSerializer) Bytes() []byte {
	panic("not implemented")
}

func (s *BaseNetConfigSerializer) String() string {
	panic("not implemented")
}
