package main

import (
	"fmt"
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

func PrintAllFiles(files *[]string) {
	fmt.Println(" _______________Found_files_______________")
	for i, file := range *files {
		fmt.Println("| ", i+1, " ", GetFileName(file))
	}
	fmt.Println("|_________________________________________")
}

func GetFileName(filepath string) string {
	result := ""

	for _, char := range filepath {
		if char == '/' {
			result = ""
		} else {
			result += string(char)
		}
	}

	return result
}
