package main

import (
	"fmt"
)

func PrintAllFiles(files *[]string) {
	fmt.Println(" _______________Found_files_______________\n|")
	for i, file := range *files {
		fmt.Println("| ", i+1, ") ", GetFileName(file))
	}
	fmt.Println("|_________________________________________")
}
