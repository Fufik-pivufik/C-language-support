package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
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

func DefaultHFile(file *os.File) error {
	defer file.Close()
	_, err := file.Write([]byte(BaseHFile))
	if err != nil {
		return err
	}
	return nil
}

func DefaultCFile(file *os.File) error {
	defer file.Close()
	_, err := file.Write([]byte(BaseCFile))
	if err != nil {
		return err
	}
	return nil
}

func DefaultClangdFile(file *os.File, path string, std string) error {
	defer file.Close()
	absPath, err := os.Getwd()
	if err != nil {
		return err
	}

	path = absPath + "/" + path
	includep := "-I" + path + "/include"
	externalp :="-I" + path + "/extend"
	
	_, err = file.Write([]byte("CompileFlags:\n\tAdd: [" +includep + ", " + externalp + ", -std=" + std + "]\n"))
	if err != nil {
		return err
	}

	return nil
}

func Execute(display bool, command string, attributes ...string) error {
	cmd := exec.Command(command, attributes...)
	if display {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
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


func GetHeaders(incPath string) []string {
result := make([]string, 0)
files, err := os.ReadDir(incPath)
if err != nil {
	fmt.Println("Error cannot read files: ", err)
	return result
}

for _, file := range files {
	filename := file.Name()
	res_c, _ := regexp.MatchString(`\.hpp$`, filename)
	res_cpp, _ := regexp.MatchString(`\.h$`, filename)
	if res_c || res_cpp {
		result = append(result, filepath.Join(incPath, filename))
	} else if file.IsDir() {
		dirPath := filepath.Join(incPath, filename)
		result = append(result, GetFiles(dirPath)...)
	}
}

return result

}


func GenInstall(conf *Config) error {
	var script string
	installFile, err := os.Create(conf.GetPath() + "/install.sh")
	if err != nil {
		return err
	}
	defer installFile.Close()

	script = conf.GetCompiler() + " -Iinclude -Iextend \\\n"
	
	files := GetFiles(conf.GetPath() + "/src")
	for _, file := range files {
		file = strings.Replace(file, conf.GetPath(), ".", 1)
		script += "\t" + file + " \\\n"
	}
	script += "\t-o " + conf.GetName()
	installFile.Write([]byte(script))
	installFile.Chmod(0777)
	return nil
}
