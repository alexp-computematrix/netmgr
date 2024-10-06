package manager

import (
	"netmgr/schema"
	"netmgr/serialize"
	"netmgr/serialize/file"
)

type INetworkManager interface {

	// ReadConfig ...
	ReadConfig(path string) error

	// WriteConfig ...
	WriteConfig(path string) error

	// SetIPv4Addresses ...
	SetIPv4Addresses(ipv4 []string, dev string) error

	// Stop() error
	// Start() error
	// Restart() error
	// Commit() error
	// Rollback() error
	// SetDHCP() error
	// SetStatic() error
}

type BaseNetworkManager struct {
	Schema     *schema.NetSchema
	Serializer serialize.NetConfigSerializer
}

// ReadConfig ...
func (n *BaseNetworkManager) ReadConfig(path string) error {
	s, err := file.NetSchemaFromFile(path, n.Serializer)
	if err == nil {
		n.Schema = s
	}
	return err
}

func (n *BaseNetworkManager) WriteConfig(path string) error {
	return file.NetSchemaToFile(path, n.Serializer, n.Schema)
}

func (n *BaseNetworkManager) SetIPv4Addresses(ipv4 []string, dev string) error {
	iface, err := n.Schema.GetInterface(dev)
	if err != nil {
		return err
	}
	for _, ip := range ipv4 {
		_, err = iface.AssociateIPAddress(ip)
		if err != nil {
			return err
		}
	}
	return nil
}

func (n *BaseNetworkManager) Stop() error {
	panic("implement me")
}

func (n *BaseNetworkManager) Start() error {
	panic("implement me")
}

func (n *BaseNetworkManager) Restart() error {
	panic("implement me")
}

func (n *BaseNetworkManager) Commit() error {
	panic("implement me")
}

func (n *BaseNetworkManager) Rollback() error {
	panic("implement me")
}

func (n *BaseNetworkManager) SetDHCP() error {
	panic("implement me")
}

func (n *BaseNetworkManager) SetStatic() error {
	panic("implement me")
}
