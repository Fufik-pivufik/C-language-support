package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	// Version  string `json:"cls-version"`
	Name     string `json:"name"`
	MainFile string `json:"main-file"`
	Compiler string `json:"compiler"`
	Path     string `json:"path"`
	TestPath string `json:"test-path"`
}

func (conf *Config) SetName(name string) {
	conf.Name = name
}

func (conf *Config) SetCompiler(compiler string) {
	conf.Compiler = compiler
}

func (conf *Config) SetMainFile(filename string) {
	conf.MainFile = filename
}

func (conf *Config) SetTestPath(filepath string) {
	conf.TestPath = filepath
}

func (conf *Config) SetPath() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: cannot find project's absolute path")
		return
	}

	conf.Path = path + "/" + conf.Name
}

func (conf *Config) GetName() string {
	return conf.Name
}

func (conf *Config) GetCompiler() string {
	return conf.Compiler
}

func (conf *Config) GetPath() string {
	return conf.Path
}

func (conf *Config) GetMainFile() string {
	return conf.MainFile
}

func (conf *Config) GetTestPath() string {
	return conf.TestPath
}

func (conf Config) display() {
	if conf.TestPath == "" {
		fmt.Printf("________config:________\n| name: %s\n| main file: %s\n| compiler: %s\n| path: %s\n", conf.GetName(), conf.GetMainFile(), conf.GetCompiler(), conf.GetPath())
		return
	}

	fmt.Printf("________config:________\n| name: %s\n| main file: %s\n| test file: %s\n| compiler: %s\n| path: %s\n", conf.GetName(), conf.GetMainFile(), conf.GetTestPath(), conf.GetCompiler(), conf.GetPath())
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

func ReadConfig(path string) *Config {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error cannot onpen config file: ", err)
		return &Config{}
	}

	var result Config
	err = json.Unmarshal(file, &result)
	if err != nil {
		fmt.Println("Error cannot parse json file: ", err)
		return &Config{}
	}

	return &result
}

func CreateConfig(projectName string) error {
	configPath := projectName + "/config.json"
	configFile, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer configFile.Close()

	config := new(Config)

	// default configuration for project
	config.SetName(projectName)
	config.SetCompiler("g++")
	config.SetPath()
	config.SetMainFile("main.cpp")
	config.SetTestPath("")

	fmt.Printf("________Created_project_%s________\n\n", projectName)
	config.display()
	jsonFile, err := json.Marshal(config)
	if err != nil {
		return err
	}

	configFile.Write(jsonFile)
	return nil
}

func ConfigUpdate(conf *Config) error {
	textUpdate, err := json.Marshal(conf)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(conf.GetPath(), "config.json"), textUpdate, 0777)
	if err != nil {
		return err
	}

	return nil
}
