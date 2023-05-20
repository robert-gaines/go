package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	files, err := ioutil.ReadDir("C:\\")

	if err != nil {
		fmt.Println("Error", err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

}
