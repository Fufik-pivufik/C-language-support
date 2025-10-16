package main

import (
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
