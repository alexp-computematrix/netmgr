package file

import (
	"netmgr/schema"
	"netmgr/serialize"
	"os"
)

type NetConfigFileSerializer struct {
	serialize.NetConfigSerializer
	file string
}

func (s *NetConfigFileSerializer) SetFile(file string) {
	s.file = file
}
func (s *NetConfigFileSerializer) LoadFromFile() (*schema.NetSchema, error) {
	data, err := os.ReadFile(s.file)
	if err != nil {
		return nil, err
	}
	return s.Deserialize(data)
}
func (s *NetConfigFileSerializer) WriteToFile(schema *schema.NetSchema) error {
	data, err := s.Serialize(schema)
	if err != nil {
		return err
	}
	return os.WriteFile(s.file, data, 0644)
}

func NewNetConfigFileSerializer(file string) NetConfigFileSerializer {
	return NetConfigFileSerializer{file: file}
}
