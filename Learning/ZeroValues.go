package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	// These variables are delcared without an explicit value
	// Their default values are set to 0 or false
	fmt.Printf("%v %v %v %q \n", i, f, b, s)
}
