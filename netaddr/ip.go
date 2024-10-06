package netaddr

import (
	"net"
)

func IsIPv4(ip net.IP) bool {
	return ip.To4() != nil
}

func IsIPv6(ip net.IP) bool {
	return ip.To4() == nil && ip.To16() != nil
}
