package main

import (
	"encoding/json"
	"fmt"
	"os"
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

	fmt.Printf("________config:________\n| name: %s\n| compiler: %s\n| path: %s\n", config.GetName(), config.GetCompiler(), config.GetPath())
	jsonFile, err := json.Marshal(config)
	fmt.Println(string(jsonFile))
	if err != nil {
		return err
	}

	configFile.Write(jsonFile)
	return nil
}
