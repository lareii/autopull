package utils

import (
	"os"
	"os/exec"
	"runtime"
)

/*
 * Clear is a function for clearing the terminal screen.
 */

func Clear() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
