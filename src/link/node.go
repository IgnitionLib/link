package link

import (
	"ignition-link/src/system"

	"github.com/fatih/color"
)

type Node struct {
	Driver  string
	Address string

	Name string
	Id   string

	Version  string
	Platform string
}

/*
 * Logs some info about the Node
 * This includes Driver, Address, Node Name, Node ID, Version and Platform
 */
func (this *Node) LogInfo() {
	system.Log(color.CyanString("--------[ Node Info - "+this.Driver+"@"+this.Address+" ]--------"), "node-info")
	system.Log(color.CyanString(this.Driver+"@"+this.Address)+color.BlackString(" > ")+"Node Name : "+this.Name, "node-info")
	system.Log(color.CyanString(this.Driver+"@"+this.Address)+color.BlackString(" > ")+"Node ID   : "+this.Id, "node-info")
	system.Log(color.CyanString(this.Driver+"@"+this.Address)+color.BlackString(" > ")+"Version   : "+this.Version, "node-info")
	system.Log(color.CyanString(this.Driver+"@"+this.Address)+color.BlackString(" > ")+"Platform  : "+this.Platform, "node-info")
}
