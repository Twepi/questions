package main

import "fmt"

func main() {

}

func exp1() {
	ch := make(chan int)
	ch <- 1
	fmt.Println("1 ended")
}

func exp2() {
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println("2 ended")
}
