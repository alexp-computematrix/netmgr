package schema

import "net"

const (
	NetSchemaAddressIPv4 NetSchemaAddressProtocol = iota
	NetSchemaAddressIPv6
)

// NetSchema ...
type NetSchema struct {
	// Addresses ...
	Addresses []NetSchemaAddress

	// Interfaces ...
	Interfaces []NetSchemaInterface

	// Routes ...
	Routes []NetSchemaRoute
}

// NetSchemaAddressProtocol ...
type NetSchemaAddressProtocol int64

// NetSchemaAddress ...
type NetSchemaAddress struct {
	// Protocol ...
	Protocol NetSchemaAddressProtocol

	// Host ...
	Host net.IP

	// Network ...
	Network *net.IPNet

	// Interface ...
	Interface *net.Interface
}

// String ...
func (nsa NetSchemaAddress) String() string {
	return nsa.Host.String()
}

// GetProtocol ...
func (nsa NetSchemaAddress) GetProtocol() NetSchemaAddressProtocol {
	return nsa.Protocol
}

// GetHost ...
func (nsa NetSchemaAddress) GetHost() net.IP {
	return nsa.Host
}

// GetNetMask ...
func (nsa NetSchemaAddress) GetNetMask() net.IPMask {
	return nsa.Network.Mask
}

// GetNetwork ...
func (nsa NetSchemaAddress) GetNetwork() *net.IPNet {
	return nsa.Network
}

// GetInterface ...
func (nsa NetSchemaAddress) GetInterface() *net.Interface {
	return nsa.Interface
}

// NetSchemaInterface ...
// TODO(alexp)
type NetSchemaInterface struct{}

// NetSchemaRoute ...
// TODO(alexp)
type NetSchemaRoute struct{}
