package file

import (
	"netmgr/serialize"
	"os"
)

type NetConfigFileSerializer struct {
	serialize.BaseNetConfigSerializer
	file string
}

func (s *NetConfigFileSerializer) SetFile(file string) {
	s.file = file
}
func (s *NetConfigFileSerializer) LoadFromFile() error {
	data, err := os.ReadFile(s.file)
	if err != nil {
		return err
	}
	return s.Deserialize(data)
}
func (s *NetConfigFileSerializer) WriteToFile() error {
	data, err := s.Serialize()
	if err != nil {
		return err
	}
	return os.WriteFile(s.file, data, 0644)
}

func NewNetConfigFileSerializer(file string) NetConfigFileSerializer {
	return NetConfigFileSerializer{file: file}
}
