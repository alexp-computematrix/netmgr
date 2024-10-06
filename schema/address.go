package schema

import (
	"fmt"
	"net"
	"netmgr/netaddr"
)

const (
	InternetProtocolV4 NetSchemaAddressProtocol = "IPv4"
	InternetProtocolV6 NetSchemaAddressProtocol = "IPv6"
)

// NetSchemaAddressProtocol ...
type NetSchemaAddressProtocol string

// NetSchemaAddress ...
type NetSchemaAddress struct {
	// Protocol ...
	Protocol NetSchemaAddressProtocol

	// Host ...
	Host net.IP

	// Network ...
	Network *net.IPNet

	// // Interface ...
	// Interface *net.Interface
}

// String ...
func (nsa NetSchemaAddress) String() string {
	return nsa.Host.String()
}

// PrefixLength ...
func (nsa NetSchemaAddress) PrefixLength() int {
	prefixLen, _ := nsa.Network.Mask.Size()
	return prefixLen
}

// BitMask ...
func (nsa NetSchemaAddress) BitMask() int {
	_, bits := nsa.Network.Mask.Size()
	return bits
}

// CIDRFormat ...
func (nsa NetSchemaAddress) CIDRFormat() string {
	return fmt.Sprintf("%s/%d", nsa.Host.String(), nsa.PrefixLength())
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
// func (nsa NetSchemaAddress) GetInterface() *net.Interface {
// 	return nsa.Interface
// }

func NewNetSchemaAddress(address string) (NetSchemaAddress, error) {
	ip, ipNet, err := net.ParseCIDR(address)
	if err != nil {
		return NetSchemaAddress{}, err
	}

	var protocol NetSchemaAddressProtocol
	switch {
	case netaddr.IsIPv4(ip):
		protocol = InternetProtocolV4
	case netaddr.IsIPv6(ip):
		protocol = InternetProtocolV6
	default:
		return NetSchemaAddress{}, fmt.Errorf("invalid address: %s", address)
	}

	return NetSchemaAddress{
		Protocol: protocol,
		Host:     ip,
		Network:  ipNet,
	}, nil
}
