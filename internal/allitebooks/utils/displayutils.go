package utils

import (
	"strings"

	tm "github.com/buger/goterm"
)

func DisplayMessage(message string, color int) {
	tm.MoveCursor(1, 4)
	tm.Printf("%s\r", strings.Repeat(" ", tm.Width()))
	tm.MoveCursor(1, 4)
	tm.Print("   ", tm.Color(message, color))
	tm.Flush()
}
