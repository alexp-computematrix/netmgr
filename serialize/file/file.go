package file

import (
	"fmt"
	"netmgr/schema"
	"netmgr/serialize"
	"os"
)

// Handler handles file-specific operations
type Handler struct {
	filePath string
}

// NewFileHandler constructor for file handler
func NewFileHandler(file string) *Handler {
	return &Handler{filePath: file}
}

// Load loads data from file
func (f *Handler) Load() ([]byte, error) {
	data, err := os.ReadFile(f.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load file: %w", err)
	}
	return data, nil
}

// Save saves data to file
func (f *Handler) Save(data []byte) error {
	return os.WriteFile(f.filePath, data, 0644)
}

// SerializedHandler responsible for working with file-based configuration serializers
type SerializedHandler struct {
	fileHandler *Handler
	serializer  serialize.NetConfigSerializer
}

// NewSerializedHandler constructor to initialize SerializedHandler
func NewSerializedHandler(file string, serializer serialize.NetConfigSerializer) *SerializedHandler {
	return &SerializedHandler{
		fileHandler: NewFileHandler(file),
		serializer:  serializer,
	}
}

// LoadFromFile loads and deserializes data from a file using the serializer
func (cfh *SerializedHandler) LoadFromFile() (*schema.NetSchema, error) {
	data, err := cfh.fileHandler.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading file: %w", err)
	}
	return cfh.serializer.Deserialize(data)
}

// WriteToFile serializes data and writes to file using the serializer
func (cfh *SerializedHandler) WriteToFile(schema *schema.NetSchema) error {
	data, err := cfh.serializer.Serialize(schema)
	if err != nil {
		return fmt.Errorf("error serializing schema: %w", err)
	}

	return cfh.fileHandler.Save(data)
}

func NetSchemaFromFile(filename string, serializer serialize.NetConfigSerializer) (*schema.NetSchema, error) {
	fileHandler := NewSerializedHandler(filename, serializer)
	return fileHandler.LoadFromFile()
}

func NetSchemaToFile(filename string, serializer serialize.NetConfigSerializer, ns *schema.NetSchema) error {
	fileHandler := NewSerializedHandler(filename, serializer)
	return fileHandler.WriteToFile(ns)
}
