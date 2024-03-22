package main

import (
	"fmt"
)

func main() {

	sl := []int{1, 2, 3, 4, 5}

	toRemove := 2

	partOne := sl[:toRemove]
	fmt.Println(partOne)

	partTwo := sl[toRemove+1:]
	fmt.Println(partTwo)

	result := append(partOne, partTwo...)
	fmt.Println(result)
}
