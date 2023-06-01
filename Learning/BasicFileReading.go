package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	filename := "ReaderFile.txt"

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error reading the file", err)
		return
	}

	trimmedContent := bytes.Trim(content, "\x00")
	fileString := string(trimmedContent)
	fmt.Println("File content: ", string(fileString))

}
