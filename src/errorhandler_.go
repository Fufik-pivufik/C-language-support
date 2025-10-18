package main

import (
	"fmt"
	"os"
)

func ArgsCheck(argc int, val int) {
	if argc < val {
		fmt.Println("Error: missing arguments\n| try:   $ cls help    for more information")
		os.Exit(1)
	}
}

func UpdateCheck(err error) {
	if err != nil {
		fmt.Println("Error config update: ", err)
		os.Exit(1)
	}
}

func CreationCheck(err error) {
	if err != nil {
		fmt.Println("Error: Failed to create file: ", err)
		os.Exit(1)
	}
}

func DirCreationCheck(err error) {
	if err != nil {
		fmt.Println("Error: cannot create directory ", err)
		os.Exit(1)
	}

}

func CompilationCheck(err error) {
	if err != nil {
		fmt.Println("Compile error: ", err)
		os.Exit(1)
	}
}

func DefaultCodeCheck(err error) {
	if err != nil {
		fmt.Println("Error: cannot create default code: ", err)
		os.Exit(1)
	}
}

func ConfigCreationCheck(err error) {
	if err != nil {
		fmt.Println("Error: cannot create default configuration: ", err)
		os.Exit(1)
	}
}
