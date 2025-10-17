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
		fmt.Println("Error: missing arguments\n| try:   $ cls help    for more information")
		return
	}

	switch argv[1] {
	case "help":
		if argc < 3 {
			fmt.Println("_______________________cls_options______________________")
			fmt.Println("| new <project_name>   creates new directory with simple structure and default hello world app:")
			fmt.Println("|\t\t\t<project_name> -> src/ -> main.cpp\n|")
			fmt.Println("| build [list of C/CPP files] [output file name]  builds all files from list with output name(default main or project name)\n|\t\t Without arguments build project from root or inner directory")
			fmt.Println("| config <display/name/compiler/path> < /new_name/new_compiler/ > you don't have to edit config by  yourself, 'display' shows current configuration")
			fmt.Println("|                                                                                                             'name' allows you change name for your project(doesn't change directory name)")
			fmt.Println("|                                                                                                             'compiler' allows you change compiler for your project")
			fmt.Println("|                                                                                                             'path' updates path to current")
			fmt.Println("| test <run/path> < /full_path_to_test> you can create your test(but only with main function. You can find example in readme file)")
			fmt.Println("|                                       'path' + <full_path_to_test> you can include test from another file")
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
			fmt.Println("Error: config file does not exist")
			return
		}
		fmt.Println("| config file found at ", configPath)
		go fmt.Println("Reading configuration file...")
		config := ReadConfig(configPath)
		files := GetFiles(config.GetPath())
		Print_all_files(&files)
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

		fmt.Printf("Compilation complete!\n| Used %s\n| Executable file \033[32m%s\033[0m\n", config.GetCompiler(), config.GetName())

	case "config":
		if argc < 3 {
			fmt.Println("Error: missing arguments\n| try:   $ cls help    for more information")
			return
		}

		isEx, configPath := ConfigExists()
		if !isEx {
			fmt.Println("Error: config file does not exist")
			return
		}

		config := ReadConfig(configPath)
		switch argv[2] {
		case "display":
			config.display()

		case "name":
			if argc < 4 {
				fmt.Println("Error: missing arguments\n| try:   $ cls help    for more information")
				return
			}

			config.Name = argv[3]
			fmt.Println("Project's name succesfully updated")

		case "compiler":
			if argc < 4 {
				fmt.Println("Error: missing arguments\n| try:   $ cls help    for more information")
				return
			}

			config.Compiler = argv[3]
			fmt.Println("Project's compiler succesfully updated")

		case "path":
			config.SetPath()
			fmt.Println("Path to project succesfully updated")

		default:
			fmt.Println("Error: unknown argument for 'config'\n| try    $ cls help   for more information")
			return
		}
		err := ConfigUpdate(config)
		if err != nil {
			panic(err)
		}

	case "test":
		if argc < 3 {
			fmt.Println("Error: missing arguments\n| try:   $ cls help    for more information")
			return
		}

		switch argv[2] {
		case "run":

		case "path":
			if argc < 4 {
				fmt.Println("Error: missing arguments\n| try:   $ cls help    for more information")
				return
			}
		default:
			fmt.Println("Error: unknown argument for 'test'\n| try    $ cls help   for more information")
		}

	default:
		fmt.Println("Error: unknown argument\n| try   $ cls help    for more information")
	}
}
