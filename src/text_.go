package main

import (
	"regexp"
	"strings"
)

func ParseInputCompile(input string) []string {
	result := strings.Split(input, " ")
	outputfile := result[len(result)-1]

	if notCfile(outputfile) {
		result[len(result)-1] = "-o"

		result = append(result, outputfile)
		return result
	}

	return result
}

func notCfile(filename string) bool {
	res_c, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+\.cpp`, filename)
	res_cpp, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+\.c`, filename)
	return !(res_c || res_cpp)
}
