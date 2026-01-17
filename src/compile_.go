package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

func DefaultCppFile(file *os.File) error {
	defer file.Close()
	_, err := file.Write([]byte(BaseCppFile))
	if err != nil {
		return err
	}
	return nil
}

func DefaultHppFile(file *os.File) error {
	defer file.Close()
	_, err := file.Write([]byte(BaseIncludeHPP))
	if err != nil {
		return err
	}
	return nil
}

func Execute(display bool, command string, attributes ...string) error {
	cmd := exec.Command(command, attributes...)
	if display {
		cmd.Stdout = os.Stdout
	}
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	return err
}

func notCfile(filename string) bool {
	res_c, _ := regexp.MatchString(`^[a-zA-Z0-9/._%+-]+\.cpp`, filename)
	res_cpp, _ := regexp.MatchString(`^[a-zA-Z0-9/._%+-]+\.c`, filename)
	return !(res_c || res_cpp)
}

func GetFiles(srcPath string) []string {

	result := make([]string, 0)
	files, err := os.ReadDir(srcPath)
	if err != nil {
		fmt.Println("Error cannot read files: ", err)
		return result
	}

	for _, file := range files {
		filename := file.Name()
		res_c, _ := regexp.MatchString(`\.cpp$`, filename)
		res_cpp, _ := regexp.MatchString(`\.c$`, filename)
		if res_c || res_cpp {
			result = append(result, filepath.Join(srcPath, filename))
		} else if file.IsDir() {
			dirPath := filepath.Join(srcPath, filename)
			result = append(result, GetFiles(dirPath)...)
		}
	}

	return result
}

func FindMainFile(files *[]string, mainname *string) int {

	for i, file := range *files {
		if GetFileName(file) == *mainname {
			return i
		}
	}

	return -1
}
