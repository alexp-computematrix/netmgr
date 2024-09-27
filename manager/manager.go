package manager

import (
	"errors"
	"net"
	"netmgr/schema"
	"netmgr/serialize"
	"strings"
)

type INetworkManager interface {
	Stop() error
	Start() error
	Restart() error
	Commit() error
	Rollback() error
	SetDHCP() error
	SetStatic() error
}

type BaseNetworkManager struct {
	Schema     *schema.NetSchema
	Serializer serialize.NetConfigSerializer
}

func (n BaseNetworkManager) Stop() error {
	panic("implement me")
}

func (n BaseNetworkManager) Start() error {
	panic("implement me")
}

func (n BaseNetworkManager) Restart() error {
	panic("implement me")
}

func (n BaseNetworkManager) Commit() error {
	panic("implement me")
}

func (n BaseNetworkManager) Rollback() error {
	panic("implement me")
}

func (n BaseNetworkManager) SetDHCP() error {
	panic("implement me")
}

func (n BaseNetworkManager) SetStatic() error {
	panic("implement me")
}

func GetNetworkInterfaceByPrefix(prefix string) (*net.Interface, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range ifaces {
		if strings.HasPrefix(iface.Name, prefix) {
			return &iface, nil
		}
	}

	return nil, errors.New("no network interfaces found")
}
