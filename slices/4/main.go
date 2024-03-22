package main

import "fmt"

func main() {

	x := make([]int, 0) // массив, длина слайса, объем массива

	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)

	y := make([]int, len(x))
	copy(y, x)
	y = append(y, 4)

	z := make([]int, len(x))
	copy(z, x)
	z = append(z, 5)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
