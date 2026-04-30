package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	argv := os.Args
	flags := []byte{}

	for i := 0; i < len(argv); {
		if argv[i] == "flag" {
			break
		}
		if argv[i][0] == '-' {
			flags = append(flags, argv[i][1])
			if argv[i][1] == 'v' {
				argv[1] = "version"
				break
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
			fmt.Printf(" C Language Support - simple project manager for c/c++.\n")
			fmt.Printf("_______________________cls_options______________________\n")
			fmt.Printf("| \033[38;2;%sm'version'/'-v'\033[0m             shows current cls version\n|\n", ColorHelp)
			fmt.Printf("| \033[38;2;%sm'new <project_name>'\033[0m  creates new directory with simple structure and default hello world app:\n", ColorHelp)
			fmt.Printf("|\t\t\t\033[38;2;%sm<project_name> -> src/ -> main.cpp\033[0m\n|\t\t\t                  \033[38;2;%sminclude/ -> include.hpp\033[0m\n|\t\t\t                  \033[38;2;%smcls.json\033[0m\n|\n", ColorHelp, ColorHelp, ColorHelp)
			fmt.Printf("| \033[38;2;%sm'new <project_name> -c'\033[0m   creates new directory with simple structure and default hello world app:\n", ColorHelp)
			fmt.Printf("|\t\t\t\033[38;2;%sm<project_name> -> src/ -> main.c\033[0m\n|\t\t\t                  \033[38;2;%sminclude/ -> include.h\033[0m\n|\t\t\t                  \033[38;2;%smcls.json\033[0m\n|\n", ColorHelp, ColorHelp, ColorHelp)
			fmt.Printf("| \033[38;2;%sm'build' \033[0m builds all files from list with output name(default main or project name)\n|\t\t Without arguments build project from root or inner directory\n", ColorHelp)
			fmt.Printf("| \033[38;2;%sm'run'\033[0m    the same thing as build. Just runs executable file after building\n", ColorHelp)
			fmt.Printf("| Flag '-h' for 'build' and 'run' hides all unneccessary information\n")
			fmt.Printf("| Flag \033[38;2;%sm'-l'\033[0m for 'build' builds your project as a \033[38;2;%smstatic library\033[0m with '.a' extension\n", ColorHelp, ColorHelp)
			fmt.Printf("|\n")
			fmt.Printf("| \033[38;2;%sm'config <show/name/compiler> < /new_name/new_compiler>'\033[0m you don't have to edit config by  yourself\n|            \033[38;2;%sm'show'\033[0m shows current configuration\n", ColorHelp, ColorHelp)
			fmt.Printf("|            \033[38;2;%sm'name'\033[0m allows you change name for your project(doesn't change directory name)\n", ColorHelp)
			fmt.Printf("|            \033[38;2;%sm'compiler'\033[0m allows you change compiler for your project\n", ColorHelp)
			fmt.Printf("|\n")
			fmt.Printf("| \033[38;2;%sm'test <create/run/path> < / /full_path_to_test>'\033[0m you can create your test(but only with main function.)\n", ColorHelp)
			fmt.Printf("|            \033[38;2;%sm'create'\033[0m creates base test file with default path: <project>/test/test.cpp\n", ColorHelp)
			fmt.Printf("|            \033[38;2;%sm'path' + <full_path_to_test>\033[0m you can include test from another file\n", ColorHelp)
			fmt.Printf("| \033[38;2;%sm'flag <add/remove/show> <flagname/filename/ >'\033[0m adds and removes flags for compilation\n|\n", ColorHelp)
			fmt.Printf("| \033[38;2;%sm'get <URL>'\033[0m downloads and installs library from \033[38;2;%smgithub.com\033[0m which was made with cls or cmake\n|\n", ColorHelp, ColorHelp)
		}

	case "version":
		dist := GetDistroName()
		comp := GetGPPVersion()
		fmt.Printf("C language support (cls) %s  (%s %s)\n", Version, dist, comp)
	

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

		headPath := argv[2] + "/include"
		err = os.Mkdir(headPath, 0777)
		DirCreationCheck(err)

		clangdf, err := os.Create(argv[2] + "/.clangd")
		CreationCheck(err)

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

			err = DefaultClangdFile(clangdf, argv[2], "c++20")
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

			err = DefaultClangdFile(clangdf, argv[2], "c23")
			DefaultCodeCheck(err)
		}

		
		


		err = CreateConfig(argv[2], Cproj)
		ConfigCreationCheck(err)

		err = Execute(false, "git", "init", argv[2])
		CompilationCheck(err)

	case "build", "run":
		buildflags := map[string]bool {
			"displ" : true,
			"raw" : false,
			"lib" : false,
		}
		for _, flag := range flags {
			switch flag {
			case 'h':
				buildflags["displ"] = false
			case 'r':
				buildflags["raw"] = true
			case 'l':
				buildflags["lib"] = true;
			default:
				fmt.Println("Unknown flag: -", flag)
			}
		}


		//building for project
		if buildflags["displ"] {
			fmt.Println("\n\n ____________Finging_config.json..._______")
		}

		config := GetConfig()
		if buildflags["displ"] {
			fmt.Println("| config file found at ", config.GetPath())
			fmt.Println("| Reading configuration file...\n|_________________________________________")
		}

		if buildflags["lib"] {
			if argv[1] == "run" {
				fmt.Printf("you cannot run lirary :(\n")
				os.Exit(2);
			}
			
			files := GetFiles(string(config.GetPath() + "/src"))
			incdir := "-I" + config.GetPath() + "/include"
				os.Mkdir(config.GetPath() + "/mod", 0777)
				modpath := config.GetPath() + "/mod"
			for i, file := range files {
				modname := "mod" + strconv.Itoa(i) + ".o"
				
				if buildflags["displ"] {
					fmt.Printf("|\n| compiling file %s to %s\n", GetFileName(file), modname)
				}


				err := Execute(true, "g++", "-c" ,incdir, file, "-o", modname)
				CompilationCheck(err)

				err = Execute(true, "mv", modname, modpath + "/" + modname)
				CompilationCheck(err)
			}

			if buildflags["displ"] {
				fmt.Printf("|\n| Compilation done!\n")
			}

			modules, err := os.ReadDir(modpath)
			if err != nil {
				fmt.Println("Error cannot read files: ", err)
				os.Exit(1)
			}

			mods := make([]string, 0)
			libname := config.GetName() + ".a"
			mods = append(mods, "rcs")
			mods = append(mods, libname)
			for _, module := range modules {
				var modname string = config.GetPath() + "/mod/" + module.Name()
				mods = append(mods, modname)
			}
			err = Execute(true, "ar", mods...)
			CompilationCheck(err)

			file, err := os.Create(config.GetName() + ".h")
			if err != nil {
				file, err = os.Open(config.GetPath() + "/" + config.GetName() + ".h")
				if err != nil {
					fmt.Println("| Cannot open file: ", err)
				}
			}
			defer file.Close()
			incPath := config.GetPath() + "/include"
			headers := GetHeaders(incPath)
			for _, header := range headers {
				file.Write([]byte("#include\""+header+"\"\n"))
			}

			currdir, err := os.Getwd()
			if err != nil {
				fmt.Println("Error cannot find your current path: ", err)
				os.Exit(1)
			}


			if currdir != config.GetPath() {
				err = Execute(true, "mv", config.GetName() + ".a", config.GetPath() + "/" + config.GetName() + ".a")
				err = Execute(true, "mv", config.GetName() + ".h", config.GetPath() + "/" + config.GetName() + ".h")
				CompilationCheck(err)
			}

			if buildflags["displ"] {
			fmt.Printf("|\n| lib file \033[38;2;100;150;255m%s\033[0m\n| header for include \033[38;2;100;150;255m%s\033[0m\n", config.GetName() + ".a", config.GetName() + ".h")
			}
			os.Exit(0)
		}

		files := GetFiles(string(config.GetPath() + "/src"))

		if buildflags["displ"] {
			PrintAllFiles(&files)
		}
		files = append(files, "-Iinclude")
		files = append(files, "-Iextend")
		files = append(files, "-o", config.GetName())

		if !buildflags["raw"] {
			files = append(files, config.Flags...)
		}

		if config.GetCXXversion() != "" {
			standart := "-std=" + config.GetCXXversion()
			files = append(files, standart)
		}
		if buildflags["displ"] {
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
			if buildflags["displ"] {
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
		case "version":
			ArgsCheck(argc, 4)
			config.SetVersion(argv[3])
			fmt.Println("Project's version succesfully updated")

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

			clangf, err := os.OpenFile(config.GetPath() + "/.clangd", os.O_RDWR|os.O_CREATE|os.O_TRUNC,0777)
			CreationCheck(err)
			err = DefaultClangdFile(clangf, argv[2], argv[3])
			DefaultCodeCheck(err)
			clangf.Close()

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
			files = append(files, config.Flags...)

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

	case "flag":
		ArgsCheck(argc, 3)
		
		config := GetConfig()
		TestExistCheck(config)

		switch argv[2] {
			case "show":
				config.FlagsShow()
			case "add":
				ArgsCheck(argc, 4)
				config.AddFlag(argv[3])
				fmt.Printf("Added flag: %s\n", argv[3])
				err := config.Update()
				UpdateCheck(err)

			case "remove":
				ArgsCheck(argc, 4)
				config.RemoveFlag(argv[3])
				fmt.Printf("Removed flag: %s\n", argv[3])
				err := config.Update()
				UpdateCheck(err)
				
			default:
			fmt.Println("| Error: unknown argument for 'flag'\n| try    $ cls help   for more information")
			os.Exit(1)
		}


	case "get":
		ArgsCheck(argc, 3)
		config := GetConfig()
		TestExistCheck(config)
		
		extPath := config.GetPath() + "/extend"
		os.Mkdir(extPath, 0777)
		err := os.Chdir(extPath)
		if err != nil {
			fmt.Println("|", err)
			os.Exit(1)	
		}
		errGo := make(chan error)
		go func(error_par chan<- error) {
			err := Execute(false, "git", "clone", argv[2])
			error_par <- err
		}(errGo)
		
		done := make(chan struct{})
		go GettingAnimation(done)
		err = <-errGo
		close(done)

		config.AddDependence(argv[2])
		config.Update()
		
		fmt.Printf("\r Downloaded repository \033[38;2;%sm%s\033[0m\n", ColorHelp, GetFileName(argv[2]))
		fmt.Printf("Checking if it's cls lib...\n")
		files, err := os.ReadDir(GetFileName(argv[2]))
		if err != nil {
			fmt.Println("|", err)
			os.Exit(1)
		}
		is_cls := false
		for _, file := range files {
			if file.Name() == "cls.json" {
				is_cls = true
				break
			}
		} 

		if is_cls {
				fmt.Printf("Library uses cls!\n")
				os.Chdir(GetFileName(argv[2]))

				err := Execute(true, "cls", "build", "-l")
				if err != nil {
					os.Exit(1)
				}
		}

		files, err = os.ReadDir(".")
		if err != nil {
			fmt.Println("|", err)
			os.Exit(1)
		}
		
		for _, file := range files {
			is_lib, _ :=regexp.MatchString(`\.a`, file.Name())
			if is_lib  {
				config.AddFlag(config.GetPath() + "/extend/" + GetFileName(argv[2]) + "/" + file.Name())
				config.Update()
				break
			}
		}


	case "release":
		config := GetConfig()
		GenInstall(config)

	default:
		fmt.Println("| Error: unknown argument\n| try   $ cls help    for more information")
		os.Exit(1)
	}
}
