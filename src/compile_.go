package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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
