package network_driver

import (
	"ignition-link/src/link"
	"ignition-link/src/system"
	"time"

	"github.com/fatih/color"
)

type NetworkDriver struct {
	NodeManager *link.NodeManager

	NetNodes []NetworkNode

	Scanner NetworkScanner
}

/*
 * Initializes a new network driver and returns it
 */
func NewNetworkDriver(NodeManager *link.NodeManager) NetworkDriver {
	driver := NetworkDriver{}
	driver.NodeManager = NodeManager
	driver.Scanner = NetworkScanner{Driver: &driver}

	system.Log("Network Driver enabled!", "driver")

	return driver
}

/*
 * Starts the NetworkScanner which discovers all Nodes on the local network
 */
func (this *NetworkDriver) StartScan() {
	system.Log(color.CyanString("network >> ")+"Starting to scan local network...", "driver")

	for {
		go this.Scan()
		time.Sleep(time.Duration(system.SCAN_INTERVAL) * time.Second)
	}
}

func (this *NetworkDriver) Scan() {
	this.Scanner.ScanLocal()

	for i := 0; i < len(this.Scanner.availableAddresses); i++ {
		addr := this.Scanner.availableAddresses[i]

		if this.GetNetworkNode(addr) == nil {
			system.Log("New Node found!", "debug")

			this.NewNetworkNode(addr)
		}
	}
}

func (this *NetworkDriver) NewNetworkNode(address string) *NetworkNode {
	node := NetworkNode{Address: address, Driver: this}
	this.NetNodes = append(this.NetNodes, node)

	go node.Connect()

	return &node
}

func (this *NetworkDriver) GetNetworkNode(address string) *NetworkNode {
	for i := 0; i < len(this.NetNodes); i++ {
		n := this.NetNodes[i]

		if n.Address == address {
			return &n
		}
	}

	return nil
}
