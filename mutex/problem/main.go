package main

import (
	"fmt"
	"sync"
)

func main() {
	shared := 0
	var wg sync.WaitGroup

	for i := 0; i <= 100_000; i++ {
		wg.Add(1)
		go func() {
			shared++
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("result: ", shared)
}
