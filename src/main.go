package main

import (
	"fmt"
	"os"
	// "os/exec"
)

func main() {
	argv := os.Args
	flags := []byte{}

	for i := 0; i < len(argv); {
		if argv[i][0] == '-' {
			flags = append(flags, argv[i][1])
			if argv[i][1] == 'v' {
				fmt.Printf("C language support(cls) %s\n", Version)
				os.Exit(0)
			}

			argv = append(argv[:i], argv[i+1:]...)
		} else if len(flags) > 0 {
			break
		} else {
			i++
		}
	}

	argc := len(argv)
	ArgsCheck(argc, 2)

	switch argv[1] {
	case "help":
		if argc < 3 {
			fmt.Println("_______________________cls_options______________________")
			fmt.Println("| 'version'              shows current cls version\n|")
			fmt.Println("| 'new <project_name>'   creates new directory with simple structure and default hello world app:")
			fmt.Println("|\t\t\t<project_name> -> src/ -> main.cpp\n|")
			fmt.Println("| 'build'  builds all files from list with output name(default main or project name)\n|\t\t Without arguments build project from root or inner directory")
			fmt.Println("| 'run'    the same thing as build. Just runs executable file after building")
			fmt.Println("|                                          ")
			fmt.Println("| 'config <show/name/compiler> < /new_name/new_compiler>' you don't have to edit config by  yourself\n|            'show' shows current configuration")
			fmt.Println("|            'name' allows you change name for your project(doesn't change directory name)")
			fmt.Println("|            'compiler' allows you change compiler for your project")
			fmt.Println("|                                          ")
			fmt.Println("| 'test <create/run/path> < / /full_path_to_test>' you can create your test(but only with main function.)")
			fmt.Println("|            'create' creates base test file with default path: <project>/test/test.cpp")
			fmt.Println("|            'path' + <full_path_to_test> you can include test from another file")
		}

	case "version":
		fmt.Printf("C language support(cls) %s\n", Version)

	case "new":
		ArgsCheck(argc, 3)
		Cproj := false
		for _, flag := range flags {
			switch flag {
			case 'c':
				Cproj = true
			default:
				fmt.Println("Unknown flag: -", flag)
			}
		}

		err := os.Mkdir(argv[2], 0777)
		DirCreationCheck(err)

		srcPath := argv[2] + "/src"
		err = os.Mkdir(srcPath, 0777)
		DirCreationCheck(err)

		headPath := argv[2] + "/headers"
		err = os.Mkdir(headPath, 0777)
		DirCreationCheck(err)

		if !Cproj {

			hppPath := headPath + "/include.hpp"
			hppFile, err := os.Create(hppPath)
			CreationCheck(err)

			err = DefaultHppFile(hppFile)
			DefaultCodeCheck(err)

			mainPath := srcPath + "/main.cpp"
			mainFile, err := os.Create(mainPath)
			CreationCheck(err)

			err = DefaultCppFile(mainFile)
			DefaultCodeCheck(err)
		} else {
			hPath := headPath + "/include.h"
			hFile, err := os.Create(hPath)
			CreationCheck(err)

			err = DefaultHFile(hFile)
			DefaultCodeCheck(err)

			mainPath := srcPath + "/main.c"
			mainFile, err := os.Create(mainPath)
			CreationCheck(err)

			err = DefaultCFile(mainFile)
			DefaultCodeCheck(err)
		}

		err = CreateConfig(argv[2], Cproj)
		ConfigCreationCheck(err)

		err = Execute(false, "git", "init", argv[2])
		CompilationCheck(err)

	case "build", "run":
		displ := true

		for _, flag := range flags {
			switch flag {
			case 'h':
				displ = false
			default:
				fmt.Println("Unknown flag: -", flag)
			}
		}
		// compiling lsit of files
		// if argc > 2 && argv[1] == "build" {
		// 	compileArgs := ParseInputCompile(os.Args[2:])
		// 	cmd := exec.Command("g++", compileArgs...)
		// 	cmd.Stdout = os.Stdout
		// 	cmd.Stderr = os.Stderr
		//
		// 	fmt.Println("Compiling...")
		//
		// 	err := cmd.Run()
		// 	CompilationCheck(err)
		// 	break
		// }

		//building for project
		if displ {
			fmt.Println("\n\n ____________Finging_config.json..._______")
		}

		config := GetConfig()
		if displ {
			fmt.Println("| config file found at ", config.GetPath())
			fmt.Println("| Reading configuration file...\n|_________________________________________\n")
		}

		files := GetFiles(string(config.GetPath() + "/src"))

		if displ {
			PrintAllFiles(&files)
		}
		files = append(files, "-o", config.GetName())

		if config.MainLangCPP() && config.GetCXXversion() != "" {
			standart := "-std=" + config.GetCXXversion()
			files = append(files, standart)
		}
		if displ {
			fmt.Println("\n\n\t\tCompiling project...")
		}

		// Compilation
		err := Execute(true, config.GetCompiler(), files...)
		CompilationCheck(err)

		// move to root dir
		if config.ExeNInRoot() {
			err = Execute(true, "mv", config.GetName(), config.GetPath())
			CompilationCheck(err)
		}

		if argv[1] == "build" {
			if displ {
				fmt.Printf("\n _______Compilation_ complete!_______\n| Used %s\n| Executable file \033[32m%s\033[0m\n", config.GetCompiler(), config.GetName())
			}
		} else {
			err := Execute(true, config.GetPath()+"/"+config.GetName(), argv[2:]...)
			CompilationCheck(err)
		}

	case "config":
		ArgsCheck(argc, 3)

		config := GetConfig()
		switch argv[2] {
		case "show":
			config.display()

		case "name":
			ArgsCheck(argc, 4)

			config.Name = argv[3]
			fmt.Println("Project's name succesfully updated")

		case "compiler":
			ArgsCheck(argc, 4)

			config.SetCompiler(argv[3])
			fmt.Println("Project's compiler succesfully updated")

		case "std":
			ArgsCheck(argc, 4)
			config.SetCXXversion(argv[3])
			fmt.Println("Language standart succesfully updated")

		default:
			fmt.Println("| Error: unknown argument for 'config'\n| try    $ cls help   for more information")
			os.Exit(1)
		}
		err := config.Update()
		UpdateCheck(err)

	case "test":
		ArgsCheck(argc, 3)

		switch argv[2] {

		case "create":

			config := GetConfig()
			err := config.CreateTest()
			CreationCheck(err)

			err = config.Update()
			UpdateCheck(err)

		case "run":
			config := GetConfig()
			TestExistCheck(config)

			files := GetFiles(string(config.GetPath() + "/src"))
			mainFile := config.GetMainFile()
			mainIndx := FindMainFile(&files, &mainFile)
			files[mainIndx] = config.GetTestPath()

			PrintAllFiles(&files)
			files = append(files, "-o", "test_outputxyz")

			if config.MainLangCPP() && config.GetCompiler() == "g++" {
				standart := "-std=" + config.GetCXXversion()
				files = append(files, standart)
			}

			err := Execute(true, config.GetCompiler(), files...)
			CompilationCheck(err)

			err = Execute(true, "./test_outputxyz")
			CompilationCheck(err)

			err = Execute(true, "rm", "test_outputxyz")
			CompilationCheck(err)
		// case "path":
		// 	ArgsCheck(argc, 4)
		//
		// 	config := GetConfig()
		// 	config.SetTestPath(argv[3])
		// 	err := config.Update()
		// 	UpdateCheck(err)

		default:
			fmt.Println("| Error: unknown argument for 'test'\n| try    $ cls help   for more information")
			os.Exit(1)
		}

	default:
		fmt.Println("| Error: unknown argument\n| try   $ cls help    for more information")
		os.Exit(1)
	}
}
