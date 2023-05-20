package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the file name-> ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered: ", input)
}
