package schema

import (
	"fmt"
	"net"
)

// NetSchema ...
// FIXME(alexp): Refactor schema for efficiency
type NetSchema struct {

	// Interfaces ...
	Interfaces map[string]*NetSchemaInterface
}

// AssociateInterface appends a new net.Interface to the schema
// Interfaces property as NetSchemaInterface.
func (s *NetSchema) AssociateInterface(i *net.Interface) (*NetSchemaInterface, error) {
	nsi := NetSchemaInterfaceFromNetInterface(i)
	s.Interfaces[i.Name] = nsi
	return nsi, nil
}

// AssociateInterfaceByName appends a new interface name string to the schema
// Interfaces property as NetSchemaInterface.
func (s *NetSchema) AssociateInterfaceByName(name string) (*NetSchemaInterface, error) {
	nsi, err := NewNetSchemaInterfaceByName(name)
	if err != nil {
		return nil, err
	}

	s.Interfaces[nsi.i.Name] = nsi

	return nsi, nil
}

// GetInterface ...
func (s *NetSchema) GetInterface(name string) (*NetSchemaInterface, error) {
	i, ok := s.Interfaces[name]
	if !ok {
		return nil, fmt.Errorf("interface %s not found", name)
	}
	return i, nil
}

// NetSchemaRoute ...
// TODO(alexp)
type NetSchemaRoute struct{}

// NewNetSchema ...
func NewNetSchema() *NetSchema {
	// TODO(alexp)
	return &NetSchema{Interfaces: make(map[string]*NetSchemaInterface)}
}
