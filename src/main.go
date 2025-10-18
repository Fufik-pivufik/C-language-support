package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	argv := os.Args
	argc := len(argv)
	ArgsCheck(argc, 2)

	switch argv[1] {
	case "help":
		if argc < 3 {
			fmt.Println("_______________________cls_options______________________")
			fmt.Println("| new <project_name>   creates new directory with simple structure and default hello world app:")
			fmt.Println("|\t\t\t<project_name> -> src/ -> main.cpp\n|")
			fmt.Println("| build [list of C/CPP files] [output file name]  builds all files from list with output name(default main or project name)\n|\t\t Without arguments build project from root or inner directory")
			fmt.Println("| config <show/name/compiler/path> < /new_name/new_compiler/ > you don't have to edit config by  yourself, 'show' shows current configuration")
			fmt.Println("|                                                                                                             'name' allows you change name for your project(doesn't change directory name)")
			fmt.Println("|                                                                                                             'compiler' allows you change compiler for your project")
			fmt.Println("|                                                                                                             'path' updates path to current")
			fmt.Println("| test <create/run/path> < / /full_path_to_test> you can create your test(but only with main function. You can find example in readme file)")
			fmt.Println("|                                                'create' creates base test file with default path: <project>/test/test.cpp")
			fmt.Println("|                                       'path' + <full_path_to_test> you can include test from another file")
		}

	case "new":
		ArgsCheck(argc, 3)

		err := os.Mkdir(argv[2], 0777)
		DirCreationCheck(err)

		srcPath := argv[2] + "/src"
		err = os.Mkdir(srcPath, 0777)
		DirCreationCheck(err)

		mainPath := srcPath + "/main.cpp"
		mainFile, err := os.Create(mainPath)
		CreationCheck(err)

		err = DefaultCppFile(mainFile)
		DefaultCodeCheck(err)

		err = CreateConfig(argv[2])
		ConfigCreationCheck(err)

	case "build":

		// compiling lsit of files
		if argc > 2 {
			compileArgs := ParseInputCompile(os.Args[2:])
			cmd := exec.Command("g++", compileArgs...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			fmt.Println("Compiling...")

			err := cmd.Run()
			CompilationCheck(err)
			break
		}

		//building for project
		fmt.Println("\n\n ____________Finging_config.json..._______")

		config := GetConfig()

		fmt.Println("| config file found at ", config.GetPath())
		fmt.Println("| Reading configuration file...\n|_________________________________________\n")

		files := GetFiles(string(config.GetPath() + "/src"))

		PrintAllFiles(&files)

		files = append(files, "-o", config.GetName())

		fmt.Println("\n\nCompiling project...")
		// Compilation
		err := Execute(config.GetCompiler(), files...)
		CompilationCheck(err)

		// move to root dir
		err = Execute("mv", config.GetName(), config.GetPath())
		CompilationCheck(err)

		fmt.Printf("\n ____________Compilation_ complete!__________\n| Used %s\n| Executable file \033[32m%s\033[0m\n", config.GetCompiler(), config.GetName())

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

			config.Compiler = argv[3]
			fmt.Println("Project's compiler succesfully updated")

		case "path":
			config.SetPath()
			fmt.Println("Path to project succesfully updated")

		default:
			fmt.Println("Error: unknown argument for 'config'\n| try    $ cls help   for more information")
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

		case "path":
			ArgsCheck(argc, 4)

			config := GetConfig()
			config.SetTestPath(argv[3])
			err := config.Update()
			UpdateCheck(err)

		default:
			fmt.Println("Error: unknown argument for 'test'\n| try    $ cls help   for more information")
			os.Exit(1)
		}

	default:
		fmt.Println("Error: unknown argument\n| try   $ cls help    for more information")
		os.Exit(1)
	}
}
