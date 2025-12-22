package main

import ()

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

func GetDirPath(filepath string) string {
	result := ""
	str := ""

	for _, char := range filepath {
		if char == '/' {
			str += string(char)
			result += str
			str = ""
		} else {
			str += string(char)
		}
	}

	return result[:len(result)-1]
}
