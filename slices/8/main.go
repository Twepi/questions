package main

import (
	"fmt"
)

func main() {

	sl := []int{1, 2, 3, 4, 5}

	changeSlice(sl)

	fmt.Println(sl)

}

func changeSlice(sl []int) {

	sl[2] = 99

}
