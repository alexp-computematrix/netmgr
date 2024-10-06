package manager

import "netmgr/serialize/file"

type NetPlanManager struct {
	BaseNetworkManager
}

func NewNetPlanManager() *NetPlanManager {
	return &NetPlanManager{
		BaseNetworkManager{
			Serializer: &file.NetPlanSerializer{},
		},
	}
}
