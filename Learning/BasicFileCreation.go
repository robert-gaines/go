package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "TestFile.txt"

	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("This is sample content")

	if err != nil {
		fmt.Println(err)
		return
	}

	err = writer.Flush()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Successfully created and populated the file: %s", fileName)
}
