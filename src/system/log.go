package system

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func Log(message string, messageType string) {
	head := ""

	dt := time.Now()
	date := dt.Format("<01-02-2006 15:04:05>")
	head += color.WhiteString(date)
	head += color.GreenString(" [" + messageType + "]")

	fmt.Println(fmt.Sprintf("%-52v ", head) + message)
}
