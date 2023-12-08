package utils

import (
	"time"

	"github.com/fatih/color"
)

/*
 * Log is a function for printing a log message.
 */

func Log(level int, message string, newLine bool) {
	var c *color.Color
	now := time.Now()

	if level == 0 {
		c = color.New(color.BgRed)
		c.Print("ERROR")
	} else if level == 1 {
		c = color.New(color.BgYellow)
		c.Print("WARNING")
	} else {
		c = color.New(color.BgGreen)
		c.Print("LOG")
	}

	print(" ")

	c = color.New(color.BgBlue)
	// This date is standard in Go. Do not change it.
	c.Print(now.Format("2006/01/02 15:04:05"))

	print(" ")
	print(message)

	if newLine {
		println()
	}
}
