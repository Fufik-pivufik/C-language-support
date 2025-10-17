package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func ParseInputCompile(result []string) []string {
	outputfile := result[len(result)-1]

	if notCfile(outputfile) {
		result[len(result)-1] = "-o"

		result = append(result, outputfile)
		return result
	}

	result = append(result, "-o")
	result = append(result, "main")
	return result
}

func notCfile(filename string) bool {
	res_c, _ := regexp.MatchString(`^[a-zA-Z0-9/._%+-]+\.cpp`, filename)
	res_cpp, _ := regexp.MatchString(`^[a-zA-Z0-9/._%+-]+\.c`, filename)
	return !(res_c || res_cpp)
}

func GetFiles(projectPath string) []string {
	srcPath := projectPath + "/src"

	result := make([]string, 0)
	files, err := os.ReadDir(srcPath)
	if err != nil {
		fmt.Println("Error cannot read files: ", err)
		return result
	}

	for _, file := range files {
		filename := file.Name()
		res_c, _ := regexp.MatchString(`^[a-zA-Z0-9/._%+-]+\.cpp`, filename)
		res_cpp, _ := regexp.MatchString(`^[a-zA-Z0-9/._%+-]+\.c`, filename)
		if res_c || res_cpp {
			result = append(result, filepath.Join(srcPath, filename))
		}
	}

	return result
}

func Print_all_files(files *[]string) {
	fmt.Println(" _______________Found_files_______________")
	for i, file := range *files {
		fmt.Println("| ", i+1, " ", file)
	}
	fmt.Println("|_________________________________________")
}
