package main

import (
	"errors"
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
