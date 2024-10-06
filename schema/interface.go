package schema

import (
	"fmt"
	"net"
)

// NetSchemaInterface ...
// TODO(alexp)
type NetSchemaInterface struct {
	i         *net.Interface
	Addresses map[string]NetSchemaAddress
	Routes    []NetSchemaRoute
}

// Name returns the name of the interface associated with the NetSchemaInterface.
func (n *NetSchemaInterface) Name() string {
	return n.i.Name
}

// Interface returns the NetSchemaInterface associated net.Interface.
func (n *NetSchemaInterface) Interface() *net.Interface {
	return n.i
}

// AssociateIPAddress appends a new IP address string in CIDR format to an
// interfaces Addresses property as NetSchemaAddress.
func (n *NetSchemaInterface) AssociateIPAddress(address string) (NetSchemaAddress, error) {
	nsa, err := NewNetSchemaAddress(address)
	if err != nil {
		return NetSchemaAddress{}, err
	}

	// associate address
	n.Addresses[nsa.CIDRFormat()] = nsa
	return nsa, nil
}

// DisassociateIPAddress removes an IP address string in CIDR format from an
// interfaces Addresses property as NetSchemaAddress.
func (n *NetSchemaInterface) DisassociateIPAddress(address string) (NetSchemaAddress, error) {
	nsa, ok := n.Addresses[address]
	if !ok {
		return NetSchemaAddress{}, fmt.Errorf("address not associated: %s", address)
	}

	// disassociate address
	delete(n.Addresses, address)
	return nsa, nil
}

func NetSchemaInterfaceFromNetInterface(i *net.Interface) *NetSchemaInterface {
	return &NetSchemaInterface{i: i, Addresses: make(map[string]NetSchemaAddress)}
}

// NewNetSchemaInterfaceByName returns a new NetSchemaInterface from an interface name string.
func NewNetSchemaInterfaceByName(name string) (*NetSchemaInterface, error) {
	i, err := net.InterfaceByName(name)
	if err != nil {
		return nil, err
	}

	return &NetSchemaInterface{i: i, Addresses: make(map[string]NetSchemaAddress)}, nil
}
