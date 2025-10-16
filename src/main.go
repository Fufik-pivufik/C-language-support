package main

import (
	"fmt"
	"os"
	// "os/exec"
)

func main() {
	argc := len(os.Args)
	if argc < 2 {
		fmt.Println("Error: missing arguments\n| use: cls help for more information")
		return
	}

}
