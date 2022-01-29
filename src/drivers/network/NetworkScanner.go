package network_driver

import (
	"ignition-link/src/system"
	"net"

	"fmt"
	"time"
)

type NetworkScanner struct {
	Driver *NetworkDriver

	p1i uint8
	p2i uint8

	availableAddresses []string
}

/*
 * Tries to connect to a local ip address on port 52133. If successful, the full address will be appended to
 * the array of availableAddresses
 */
func (this *NetworkScanner) CheckAddress(p1 int, p2 int) {
	// Address with format: "192.168.{p1}.{p2}:52133"
	addr := "192.168." + fmt.Sprint(p1) + "." + fmt.Sprint(p2) + ":52133"

	if addr == "192.168.1.119:52133" {
		fmt.Println(addr)
		fmt.Print("NetworkScanner.go > ")
		fmt.Println(*this.Driver)
	}

	if system.IsStringInArray(addr, this.Driver.AllConnectedAddresses()) {
		return
	}

	d := net.Dialer{Timeout: time.Millisecond * 500}
	conn, err := d.Dial("tcp", addr)

	// Connection did not time out
	if err == nil {
		this.availableAddresses = append(this.availableAddresses, addr)

		conn.Close()
	}
}

/*
 * Scans the local network for Ignition Nodes
 */
func (this *NetworkScanner) ScanLocal() {
	this.availableAddresses = nil

	for i1 := 0; i1 < 16; i1++ {
		for i2 := 0; i2 < 256; i2++ {
			go this.CheckAddress(i1, i2)
		}

		time.Sleep(time.Millisecond * 300)
	}
}
