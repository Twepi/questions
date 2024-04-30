package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {

			if i%2 == 0 {
				fmt.Println(i * i)
				return
			}

			fmt.Println(i)

			wg.Done()
		}(i)
	}

	wg.Wait()
}
