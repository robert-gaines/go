package main

import (
	"fmt"
	"math/cmplx"
)

// Example of a block based variable declaration

var (
	BooleanOne     bool       = false
	MaximumInteger uint64     = 1<<64 - 1
	z              complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v \n", BooleanOne, BooleanOne)
	fmt.Printf("Type: %T Value: %v \n", MaximumInteger, MaximumInteger)
	fmt.Printf("Type: %T Value: %v \n", z, z)
}
