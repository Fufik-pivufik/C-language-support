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
	CXXstd   string `json:"c++ standart"`
	Path     string `json:"path"`
	TestPath string `json:"test-path"`
}

// Config Setters

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

func (conf *Config) SetCXXversion(standart string) {
	conf.CXXstd = standart
}

func (conf *Config) SetPath() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: cannot find project's absolute path")
		return
	}

	conf.Path = path
}

// Conifg Getters

func (conf *Config) GetName() string {
	return conf.Name
}

func (conf *Config) GetCompiler() string {
	return conf.Compiler
}

func (conf *Config) GetCXXversion() string {
	return conf.CXXstd
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

// Config Methods

func (conf *Config) MainLangCPP() bool {
	result := ""
	for _, ch := range conf.GetMainFile() {
		if ch == '.' {
			result = ""
		}

		result += string(ch)
	}
	return result == ".cpp"
}

func (conf *Config) display() {
	if conf.TestPath == "" {
		fmt.Printf("________config:________\n| name: %s\n| main file: %s\n| compiler: %s\n| standart c++: %s\n| path: %s\n", conf.GetName(), conf.GetMainFile(), conf.GetCompiler(), conf.GetCXXversion(), conf.GetPath())
		return
	}

	fmt.Printf("________config:________\n| name: %s\n| main file: %s\n| test file: %s\n| compiler: %s\n| standart c++: %s\n| path: %s\n", conf.GetName(), conf.GetMainFile(), conf.GetTestPath(), conf.GetCompiler(), conf.GetCXXversion(), conf.GetPath())
}

func (conf *Config) ExeNInRoot() bool {
	_, err := os.Stat(filepath.Join(conf.GetPath(), conf.GetName()))
	currentPath, _ := os.Getwd()
	return errors.Is(err, os.ErrNotExist) || currentPath != conf.GetPath()
}

func (conf *Config) CreateTest() error {
	conf.SetTestPath(filepath.Join(conf.GetPath(), "test/test.cpp"))
	err := os.Mkdir(GetDirPath(conf.GetTestPath()), 0777)
	if err != nil {
		return err
	}

	file, err := os.Create(conf.GetTestPath())
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(BaseTestCppFile))
	if err != nil {
		return err
	}

	fmt.Println("Test file succesfully created at: ", conf.GetTestPath())

	return nil
}

func (conf *Config) Update() error {
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

// Functions with Config

func ConfigExists() (bool, string) {
	file := "config.json"
	err := os.ErrNotExist
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
	config.SetCXXversion("c++20")
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

func GetConfig() *Config {

	isEx, configPath := ConfigExists()
	if !isEx {
		fmt.Println("| Error: config file does not exist")
		os.Exit(1)
	}

	config := ReadConfig(configPath)
	return config
}
