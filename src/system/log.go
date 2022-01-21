package system

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func Log(message string) {
	msg := ""

	dt := time.Now()
	date := dt.Format("<01-02-2006 15:04:05> ")
	msg += color.WhiteString(date)

	msg += message

	fmt.Println(msg)
}
