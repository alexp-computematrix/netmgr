package file

type NetPlanConfig struct {
	Network struct {
		Version          int         `yaml:"version,omitempty"`
		Renderer         string      `yaml:"renderer,omitempty"`
		Bonds            interface{} `yaml:"bonds,omitempty"`             // TODO(alexp): Create type
		Bridges          interface{} `yaml:"bridges,omitempty"`           // TODO(alexp): Create type
		DummyDevices     interface{} `yaml:"dummy-devices,omitempty"`     // TODO(alexp): Create type
		Ethernets        struct{}    `yaml:"ethernets,omitempty"`         // TODO(alexp): Create type
		Modems           interface{} `yaml:"modems,omitempty"`            // TODO(alexp): Create type
		Tunnels          interface{} `yaml:"tunnels,omitempty"`           // TODO(alexp): Create type
		VirtualEthernets interface{} `yaml:"virtual-ethernets,omitempty"` // TODO(alexp): Create type
		VLANs            interface{} `yaml:"vlans,omitempty"`             // TODO(alexp): Create type
		VRFs             interface{} `yaml:"vrfs,omitempty"`              // TODO(alexp): Create type
		WIFIs            interface{} `yaml:"wifis,omitempty"`             // TODO(alexp): Create type
		NMDevices        interface{} `yaml:"nm-devices,omitempty"`        // TODO(alexp): Create type
	}
}

type NetPlanSerializer struct {
	NetConfigFileSerializer
}

func (s *NetPlanSerializer) Serialize() ([]byte, error) {
	return nil, nil
}

func (s *NetPlanSerializer) Deserialize(bytes []byte) error {
	return nil
}

func (s *NetPlanSerializer) Bytes() []byte {
	if bytes, err := s.Serialize(); err == nil {
		return bytes
	}
	return []byte{}
}

func (s *NetPlanSerializer) String() string {
	if bytes, err := s.Serialize(); err == nil {
		return string(bytes)
	}
	return ""
}
