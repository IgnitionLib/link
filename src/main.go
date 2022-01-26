package main

import (
	network_driver "ignition-link/src/drivers/network"
	"ignition-link/src/link"
	"ignition-link/src/system"

	"github.com/fatih/color"
)

func main() {
	system.Log(color.GreenString("Starting IgnitionLink..."), "info")

	nodeManager := link.NodeManager{}

	networkDriver := network_driver.NewNetworkDriver(&nodeManager)
	go networkDriver.StartScan()

	// Continuous loop to prevent Network Scan goroutine from being stopped
	for {

	}
}
