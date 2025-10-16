package main

import (
	"fmt"
	"os"
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
