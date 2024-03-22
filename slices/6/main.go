package main

import (
	"fmt"
)

func main() {

	sl := []int{1, 2, 3, 4, 5}

	addToSlice(sl)

	fmt.Println(sl)

}

func addToSlice(sl []int) {

	sl = append(sl, 6)

}
