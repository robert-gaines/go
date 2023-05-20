package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "example.txt"

	if fileInfo, err := os.Stat(filename); err == nil {
		if fileInfo.Mode().IsRegular() {
			fmt.Println("File exists!")
		} else {
			fmt.Println("File is not a regular file")
		}
	} else if os.IsNotExist(err) {
		fmt.Println("File does not exist.")
	} else {
		fmt.Println("Error occurred: ", err)
	}
}
