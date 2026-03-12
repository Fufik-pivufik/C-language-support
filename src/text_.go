package main

import (
	"strings"
	"os"
	"bufio"
	"os/exec"
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


func GetDistroName() string {
	release, err := os.Open("/etc/os-release")
	if err != nil {
		return ""
	}
	defer release.Close()

	scanner := bufio.NewScanner(release)
	distro := ""

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "NAME=") {
			distro = line[6:len(line) - 1]
			return distro
		}
	}
	return distro
}

func GetGPPVersion() string {
	cmd := exec.Command("g++",  "--version")

	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	lines := strings.Split(string(output), "\n")

	return lines[0]
}
