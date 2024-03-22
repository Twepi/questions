package main

import (
	"fmt"
	"sort"
)

func main() {

	sl := []int{7, 9, 15, 6}

	sortSlice(sl)

	fmt.Println(sl)

}

func sortSlice(sl []int) {

	sort.Slice(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})

}
