package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Name     string `json:"name"`
	Compiler string `json:"compiler"`
	Path     string `json:"path"`
}

func (conf *Config) SetName(name string) {
	conf.Name = name
}

func (conf *Config) SetCompiler(compiler string) {
	conf.Compiler = compiler
}

func (conf *Config) SetPath() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: cannot find project's absolute path")
		return
	}

	conf.Path = path + "/" + conf.Name
}

func (conf Config) GetName() string {
	return conf.Name
}

func (conf Config) GetCompiler() string {
	return conf.Compiler
}

func (conf Config) GetPath() string {
	return conf.Path
}

func (conf Config) display() {
	fmt.Printf("________config:________\n| name: %s\n| compiler: %s\n| path: %s\n", conf.GetName(), conf.GetCompiler(), conf.GetPath())
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
