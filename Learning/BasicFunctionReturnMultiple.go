package main

import "fmt"

func swap(x string, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("first", "second")
	fmt.Println(a, b)
}
