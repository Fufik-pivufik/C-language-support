package main

import (
	"fmt"
	"os"
	"os/exec"
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
		if argc < 3 {
			fmt.Println("_______________________cls_options______________________\n| new <project_name>   creates new directory with simple structure and default hello world app:")
			fmt.Println("|\t\t\t<project_name> -> src/ -> main.cpp\n|\n| run [list of C/CPP files]  builds and executes all source files from list\n|\t\tWithout list of files executes project from root or inner directory\n|")
			fmt.Println("| build [list of C/CPP files] [output file name]  builds all files from list with output name(default main or project name)\n|\t\t Without arguments build project from root or inner directory")
			//fmt.Println("| ")
		}

	case "new":
		if argc < 3 {
			fmt.Println("Error: missing arguments for 'new'\n| try:   $ cls help new    for more information")
			return
		}

		err := os.Mkdir(argv[2], 0777)
		if err != nil {
			fmt.Println("Error: cannot create directory ", err)
			return
		}

		srcPath := argv[2] + "/src"
		err = os.Mkdir(srcPath, 0777)
		if err != nil {
			fmt.Println("Error: cannot create directory ", err)
			return
		}

		mainPath := srcPath + "/main.cpp"
		mainFile, err := os.Create(mainPath)
		if err != nil {
			fmt.Println("Error: cannot create file ", err)
			return
		}

		err = DefaultCppFile(mainFile)
		if err != nil {
			fmt.Println("Error: cannot create default code: ", err)
			return
		}
		err = CreateConfig(argv[2])
		if err != nil {
			fmt.Println("Error: cannot create default configuration: ", err)
			return
		}

	case "build":

		// compiling lsit of files
		if argc > 2 {
			compileArgs := ParseInputCompile(os.Args[2:])
			cmd := exec.Command("g++", compileArgs...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			go func() {
				fmt.Println("Compiling...")
			}()

			err := cmd.Run()
			if err != nil {
				fmt.Println("Compile error: ", err)
				os.Exit(1)
			}

			break
		}

		//building for project
		go fmt.Println("Finging config.json...")
		isEx, configPath := ConfigExists()
		if !isEx {
			fmt.Println("Error: config file does not exists")
			return
		}

		go fmt.Println("Reading configuration file...")
		config := ReadConfig(configPath)
		files := GetFiles(config.GetPath())
		files = append(files, "-o", config.GetName())
		cmd := exec.Command(config.GetCompiler(), files...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		go fmt.Println("Compiling project...")

		err := cmd.Run()
		if err != nil {
			fmt.Println("Compile error: ", err)
			os.Exit(1)
		}

		go func() {
			fmt.Printf("Compilation complete!\n| Used %s\n| Executable file %s\n", config.GetCompiler(), config.GetName())
		}()

	default:
		fmt.Println("Error: unknown argument\n| try   $ cls help    for more information")
	}
}
