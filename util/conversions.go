package util

import (
	"fmt"
	"net"
	"netmgr/schema"
)

func isIPv4(addr string) bool {
	ip := net.ParseIP(addr)

	if ip == nil {
		ip, _, _ = net.ParseCIDR(addr)
		if ip == nil {
			return false
		}
	}

	return ip.To4() != nil
}

func isIPv6(addr string) bool {
	ip := net.ParseIP(addr)

	if ip == nil {
		ip, _, _ = net.ParseCIDR(addr)
		if ip == nil {
			return false
		}
	}

	return ip.To4() == nil && ip.To16() != nil
}

func ConvertStringToNetSchemaAddress(address string) (*schema.NetSchemaAddress, error) {
	var protocol schema.NetSchemaAddressProtocol

	switch {
	case isIPv4(address):
		protocol = schema.NetSchemaAddressIPv4
	case isIPv6(address):
		protocol = schema.NetSchemaAddressIPv6
	default:
		return nil, fmt.Errorf("invalid address: %s", address)
	}

	ip, ipNet, err := net.ParseCIDR(address)
	if err != nil {
		return nil, err
	}

	return &schema.NetSchemaAddress{
		Protocol: protocol,
		Host:     ip,
		Network:  ipNet,
	}, nil
}
