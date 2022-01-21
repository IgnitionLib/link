package main

import (
	"ignition-link/src/system"

	"github.com/fatih/color"
)

func main() {
	system.Log(color.GreenString("Starting IgnitionLink..."))

	system.Log(color.GreenString("New Device Found: ") + color.CyanString("network >> 192.168.1.69"))
}
