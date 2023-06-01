package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("tasklist")
	} else {
		cmd = exec.Command("ps", "-ax")
	}
	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(out))
}
