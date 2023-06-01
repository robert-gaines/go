package main

import "fmt"

func main() {

	thisArray := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("Array length is: ", len(thisArray))

	for i := 0; i < len(thisArray); i++ {
		fmt.Println(thisArray[i])
	}

	fmt.Println()

	j := 0
	for j < len(thisArray) {
		fmt.Println(thisArray[j])
		j++
	}

	fmt.Println()

	k := 0
	for {
		fmt.Println(thisArray[k])
		k++
		if k == len(thisArray) {
			break
		}
	}

}
