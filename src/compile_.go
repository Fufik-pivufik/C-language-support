package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const baseCppFile string = "#include <iostream>\n\nint main()\n{\n\tstd::cout << " + `"Hello, World!"` + "<< std::endl;\n\treturn 0;\n}\n"

func DefaultCppFile(file *os.File) error {
	defer file.Close()
	_, err := file.Write([]byte(baseCppFile))
	if err != nil {
		return err
	}
	return nil
}

func CreateConfig(projectName string) error {
	configPath := projectName + "/config.json"
	configFile, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer configFile.Close()

	config := new(Config)
	config.SetName(projectName)
	config.SetCompiler("g++")
	config.SetPath()

	fmt.Printf("________Created_project_%s________\n\n", projectName)
	fmt.Printf("________config:________\n| name: %s\n| compiler: %s\n| path: %s\n", config.GetName(), config.GetCompiler(), config.GetPath())
	jsonFile, err := json.Marshal(config)
	if err != nil {
		return err
	}

	configFile.Write(jsonFile)
	return nil
}

func ConfigExists() (bool, string) {
	file := "config.json"
	var err error = os.ErrNotExist
	home, _ := os.UserHomeDir()
	home, _ = filepath.Abs(home)
	lastPath := ""
	filePath, _ := os.Getwd()
	for errors.Is(err, os.ErrNotExist) {
		lastPath = filePath
		_, err = os.Stat(filepath.Join(filePath, file))

		filePath = filepath.Dir(filePath)

		if filePath == home {
			return false, ""
		}

	}
	return true, filepath.Join(lastPath, file)
}
