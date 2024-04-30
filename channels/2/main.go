package main

import "fmt"

func main() {
	resCh := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("[%d]: ended\n", i)
			res := i * i
			resCh <- res
		}
	}()

	for res := range resCh {
		fmt.Printf("result: %d\n", res)
	}

	fmt.Println("programm ended")
}
