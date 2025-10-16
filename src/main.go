package main

import (
	"fmt"
	"os"
	// "os/exec"
)

func main() {
	argv := os.Args
	argc := len(argv)
	if argc < 2 {
		fmt.Println("Error: missing arguments\n| use:   $ cls help    for more information")
		return
	}

	switch argv[1] {
	case "help":
		fmt.Println("_______________________cls_options______________________\n| new <project_name>   creates new directory and initializes git repository with simple structure and default hello world app:")
		fmt.Println("|\t\t\t<project_name> -> src/ -> main.cpp\n|\n| run [list of C/CPP files]  builds and executes all source files from list\n|\t\tWithout list of files executes project from root or inner directory\n|")
		fmt.Println("| build [list of C/CPP files] [output file name]  builds all files from list with output name(default main)\n|\t\t Without arguments build project from root or inner directory")
		//fmt.Println("| ")
	}

}
