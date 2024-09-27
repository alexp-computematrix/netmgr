package file

import (
	"gopkg.in/yaml.v3"
	"net"
	"netmgr/schema"
	"netmgr/util"
)

type NetPlanEthernet struct {
	DHCP4       bool                `yaml:"dhcp4,omitempty"`
	Addresses   []string            `yaml:"addresses,omitempty"`
	Gateway4    string              `yaml:"gateway4,omitempty"`
	Gateway6    string              `yaml:"gateway6,omitempty"`
	Nameservers map[string][]string `yaml:"nameservers,omitempty"`
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
	NetConfigFileSerializer
	yaml NetPlanYAML
}

func (s *NetPlanSerializer) Serialize(schema *schema.NetSchema) ([]byte, error) {
	// TODO(alexp): finish implementation

	for _, addr := range schema.Addresses {
		addrInterface := addr.Interface.Name

		var eth NetPlanEthernet
		if s.yaml.HasEthernet(addrInterface) {
			eth = s.yaml.Network.Ethernets[addrInterface]
		}
		eth.Addresses = append(eth.Addresses, addr.String())
	}

	return yaml.Marshal(s.yaml)
}

func (s *NetPlanSerializer) Deserialize(data []byte) (*schema.NetSchema, error) {
	// TODO(alexp): finish implementation

	if err := yaml.Unmarshal(data, &s.yaml); err != nil {
		return nil, err
	}

	netSchema := &schema.NetSchema{}

	for eth, ethConfig := range s.yaml.Network.Ethernets {
		ethInterface, err := net.InterfaceByName(eth)
		if err != nil {
			return nil, err
		}

		for _, addr := range ethConfig.Addresses {
			schemaIp, err := util.ConvertStringToNetSchemaAddress(addr)
			if err != nil {
				return nil, err
			}
			schemaIp.Interface = ethInterface
			netSchema.Addresses = append(netSchema.Addresses, *schemaIp)
		}
	}

	return netSchema, nil
}
