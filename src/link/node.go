package link

import (
	"ignition-link/src/link/protocol"
	"ignition-link/src/system"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type Node struct {
	Driver  string
	Address string
	UID     uint64

	ProjectName string
	ProjectId   string

	Version  string
	Platform string
}

/*
 * Logs some info about the Node
 * This includes Driver, Address, Node Name, Node ID, Version and Platform
 */
func (this *Node) LogInfo() {
	head := "--------[ Node Info - " + this.Driver + "@" + this.Address + " - " + strconv.FormatUint(this.UID, 16) + " ]--------"
	system.Log(color.CyanString(head), "node-info")
	system.Log(color.CyanString(this.Driver+"@"+this.Address)+color.BlackString(" > ")+"Unique ID    : "+strconv.FormatUint(this.UID, 16), "node-info")
	system.Log(color.CyanString(this.Driver+"@"+this.Address)+color.BlackString(" > ")+"Project Name : "+this.ProjectName, "node-info")
	system.Log(color.CyanString(this.Driver+"@"+this.Address)+color.BlackString(" > ")+"Project ID   : "+this.ProjectId, "node-info")
	system.Log(color.CyanString(this.Driver+"@"+this.Address)+color.BlackString(" > ")+"Version      : "+this.Version, "node-info")
	system.Log(color.CyanString(this.Driver+"@"+this.Address)+color.BlackString(" > ")+"Platform     : "+this.Platform, "node-info")
	system.Log(color.CyanString(strings.Repeat("-", len(head))), "node-info")
}

func (this *Node) RecvData(packet *protocol.Packet) {

}
