package serialize

import (
	"netmgr/schema"
)

type NetConfigSerializer interface {
	Serialize(schema *schema.NetSchema) ([]byte, error)
	Deserialize(data []byte) (*schema.NetSchema, error)
}

type BaseNetConfigSerializer struct{}

func (s *BaseNetConfigSerializer) Serialize(schema *schema.NetSchema) ([]byte, error) {
	panic("not implemented")
}

func (s *BaseNetConfigSerializer) Deserialize(data []byte) (*schema.NetSchema, error) {
	panic("not implemented")
}
