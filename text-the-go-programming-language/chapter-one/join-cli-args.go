package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Join all args after the program name using a space as the delimiter
	fmt.Println(strings.Join(os.Args[1:], " "))
}
