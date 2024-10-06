package file

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log/slog"
	"net"
	"netmgr/netdev"
	"netmgr/schema"
)

type NetPlanRoute struct {
	From                    string `yaml:"from,omitempty"`
	To                      string `yaml:"to,omitempty"`
	Via                     string `yaml:"via,omitempty"`
	OnLink                  bool   `yaml:"on-link,omitempty"`
	Metric                  int    `yaml:"metric,omitempty"`
	Type                    string `yaml:"type,omitempty"`
	Scope                   string `yaml:"scope,omitempty"`
	Table                   int    `yaml:"table,omitempty"`
	MTU                     int    `yaml:"mtu,omitempty"`
	CongestionWindow        int    `yaml:"congestion-window,omitempty"`
	AdvertisedReceiveWindow int    `yaml:"advertised-receive-window,omitempty"`
	AdvertisedMSS           int    `yaml:"advertised-mss,omitempty"`
}

type NetPlanEthernetMatch struct {
	Name       string `yaml:"name"`
	MacAddress string `yaml:"macaddress"`
	Driver     string `yaml:"driver"`
}

type NetPlanEthernet struct {
	Match       NetPlanEthernetMatch `yaml:"match"`
	DHCP4       bool                 `yaml:"dhcp4,omitempty"`
	Addresses   []string             `yaml:"addresses,omitempty"`
	Routes      []NetPlanRoute       `yaml:"routes,omitempty"`
	Nameservers map[string][]string  `yaml:"nameservers,omitempty"`
}

type NetPlanEthernets map[string]NetPlanEthernet

type NetPlanYAML struct {
	Network struct {
		Version          int              `yaml:"version,omitempty"`
		Renderer         string           `yaml:"renderer,omitempty"`
		Bonds            interface{}      `yaml:"bonds,omitempty"`             // TODO(alexp): Create type
		Bridges          interface{}      `yaml:"bridges,omitempty"`           // TODO(alexp): Create type
		DummyDevices     interface{}      `yaml:"dummy-devices,omitempty"`     // TODO(alexp): Create type
		Ethernets        NetPlanEthernets `yaml:"ethernets,omitempty"`         // TODO(alexp): Update type with full support
		Modems           interface{}      `yaml:"modems,omitempty"`            // TODO(alexp): Create type
		Tunnels          interface{}      `yaml:"tunnels,omitempty"`           // TODO(alexp): Create type
		VirtualEthernets interface{}      `yaml:"virtual-ethernets,omitempty"` // TODO(alexp): Create type
		VLANs            interface{}      `yaml:"vlans,omitempty"`             // TODO(alexp): Create type
		VRFs             interface{}      `yaml:"vrfs,omitempty"`              // TODO(alexp): Create type
		WIFIs            interface{}      `yaml:"wifis,omitempty"`             // TODO(alexp): Create type
		NMDevices        interface{}      `yaml:"nm-devices,omitempty"`        // TODO(alexp): Create type
	}
}

func (y *NetPlanYAML) HasEthernet(eth string) bool {
	_, ok := y.Network.Ethernets[eth]
	return ok
}

type NetPlanSerializer struct {
	yaml    *NetPlanYAML
	matches map[string]string
}

func (s *NetPlanSerializer) HasMatch(eth string) bool {
	_, ok := s.matches[eth]
	return ok
}

func (s *NetPlanSerializer) Serialize(schema *schema.NetSchema) ([]byte, error) {
	// TODO(alexp): finish / fix this implementation

	for _, nsi := range schema.Interfaces {
		eth := nsi.Name()
		if s.HasMatch(nsi.Name()) {
			eth = s.matches[nsi.Name()]
		}

		// create address list
		var ethAddrs []string
		for _, addr := range nsi.Addresses {
			ethAddrs = append(ethAddrs, addr.CIDRFormat())
		}

		// TODO(alexp): Implement cases for all supported interface types
		switch {
		case netdev.IsEthernet(nsi.Interface()):
			if s.yaml.HasEthernet(eth) {
				// TODO(alexp): set ethernet addresses from schema
				ethConfig := s.yaml.Network.Ethernets[eth]
				ethConfig.Addresses = ethAddrs
				s.yaml.Network.Ethernets[eth] = ethConfig
			} else {
				s.yaml.Network.Ethernets[eth] = NetPlanEthernet{
					Addresses: ethAddrs,
				}
			}
		}
	}

	return yaml.Marshal(s.yaml)
}

func (s *NetPlanSerializer) Deserialize(data []byte) (*schema.NetSchema, error) {
	s.matches = make(map[string]string)
	if err := yaml.Unmarshal(data, &s.yaml); err != nil {
		return nil, err
	}

	netSchema := schema.NewNetSchema()

	for eth, ethConfig := range s.yaml.Network.Ethernets {

		nsi, err := netSchema.AssociateInterfaceByName(eth)
		if err != nil {
			var ethInterface *net.Interface
			if ethInterface, err = InterfaceByNetPlanEthernetMatch(ethConfig.Match); err != nil {
				return nil, err
			}

			nsi, err = netSchema.AssociateInterface(ethInterface)
			if err != nil {
				return nil, err
			}

			// Create interface match mapping
			s.matches[ethInterface.Name] = eth
		}

		for _, addr := range ethConfig.Addresses {
			var nsa schema.NetSchemaAddress
			nsa, err = nsi.AssociateIPAddress(addr)
			if err != nil {
				return nil, err
			}

			slog.Info("Interface associated IP address",
				slog.String("dev", nsi.Name()),
				slog.String("addr", nsa.CIDRFormat()))
		}
	}

	return netSchema, nil
}

func InterfaceByNetPlanEthernetMatch(match NetPlanEthernetMatch) (*net.Interface, error) {
	ifaces, err := InterfacesByNetPlanEthernetMatch(match)
	if err != nil {
		return nil, err
	}
	if len(ifaces) > 1 {
		return nil, fmt.Errorf("ambiguous interface match: %v", match)
	}
	return ifaces[0], nil
}

func InterfacesByNetPlanEthernetMatch(match NetPlanEthernetMatch) ([]*net.Interface, error) {
	if ifaces, err := netdev.InterfacesByMatch(match.Name); err == nil {
		return ifaces, nil
	}
	if iface, err := netdev.InterfaceByMAC(match.MacAddress); err == nil {
		return []*net.Interface{iface}, nil
	}
	return netdev.InterfacesByDriver(match.Driver)
}
