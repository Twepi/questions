package main

import (
	"fmt"
)

func main() {
	// var wg sync.WaitGroup

	go func() {
		fmt.Println("Hello 1!")
	}()

	go func() {
		fmt.Println("Hello 2!")
	}()

	go func() {
		fmt.Println("Hello 3!")
	}()

	fmt.Println("programm ended")
}
