package main

import (
	"log/slog"
	"netmgr/manager"
)

var (
	netConfigFile = "/etc/netplan/01-virtualbox.yaml"
	netInterface  = "enp0s3"
	netAddresses  = []string{
		"10.0.2.15/24",
		"10.0.2.16/24",
		"10.0.2.17/24",
		"10.0.2.18/24",
		"10.0.2.19/24",
		"10.0.2.20/24",
	}
)

func main() {
	nm := manager.NewNetPlanManager()

	err := nm.ReadConfig(netConfigFile)
	if err != nil {
		slog.Error("Failed to read netplan file", slog.String("error", err.Error()))
	}

	slog.Info("Read netplan file", slog.Any("schema", nm.Schema))

	err = nm.SetIPv4Addresses(netAddresses, netInterface)
	if err != nil {
		slog.Error("Failed to set IPv4 addresses", slog.String("error", err.Error()))
	}

	// slog.Info("Set IPv4 addresses", slog.Any("addrs", netAddresses))

	err = nm.WriteConfig(netConfigFile)
	if err != nil {
		slog.Error("Failed to write netplan file", slog.String("error", err.Error()))
	}
}
