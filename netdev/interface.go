package netdev

import (
	"fmt"
	"net"
	"path/filepath"
	"strings"
)

func IsEthernet(i *net.Interface) bool {
	return strings.HasPrefix(i.Name, "e")
}

func IsVirtual(i *net.Interface) bool {
	return strings.HasPrefix(i.Name, "v")
}

func IsBridge(i *net.Interface) bool {
	return strings.HasPrefix(i.Name, "b")
}

func IsTunnel(i *net.Interface) bool {
	return strings.HasPrefix(i.Name, "t")
}

func InterfacesByMatch(match string) ([]*net.Interface, error) {
	undeterminedErr := fmt.Errorf("undetermined interfaces match: %s", match)
	devicesPath := "/sys/class/net"

	var ifs []*net.Interface
	netDevices, err := filepath.Glob(filepath.Join(devicesPath, match))
	if err != nil {
		return nil, err
	}

	if netDevices == nil || len(netDevices) == 0 {
		return nil, undeterminedErr
	}

	for _, netDevice := range netDevices {
		var iface *net.Interface
		iface, err = net.InterfaceByName(filepath.Base(netDevice))
		if err != nil {
			return nil, err
		}

		ifs = append(ifs, iface)
	}

	return ifs, nil
}
func InterfaceByMAC(match string) (*net.Interface, error) {
	// TODO(alexp): Implement
	// Use /sys/class/net/$DEV/address
	return nil, nil
}
func InterfacesByDriver(match string) ([]*net.Interface, error) {
	// TODO(alexp): Implement
	return nil, nil
}
