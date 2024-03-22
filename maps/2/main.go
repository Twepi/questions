package main

import "fmt"

func main() {

	set := map[int]struct{}{} // struct{} - ничего весит 0 байт поэтому его используют для сетов

	arr := []int{1, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3}

	for _, el := range arr {
		set[el] = struct{}{}
	}

	fmt.Println("длина -- ", len(set))

	for key := range set {
		fmt.Println(key)
	}
}
