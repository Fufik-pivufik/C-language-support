package main

import (
	"fmt"
	"time"
)

func PrintAllFiles(files *[]string) {
	fmt.Println(" _______________Found_files_______________\n|")
	for i, file := range *files {
		fmt.Println("| ", i+1, ") ", GetFileName(file))
	}
	fmt.Println("|_________________________________________")
}


func GettingAnimation(done <-chan struct{}) {
	loading := []string{ "[--------]", "[ ------]", "[- -----]", "[-- ----]", "[--- ---]", "[---- --]", "[----- -]", "[------ ]"}
	i := 0
	for {
		select {
		case <-done:
			fmt.Print("\rDone!")
			return
		default:
			fmt.Printf("\r%s Downloading...", loading[i % 8])
			i++
			time.Sleep(100 * time.Millisecond)
		}
	}
}
