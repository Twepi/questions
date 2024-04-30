package main

// 1. используя функцию printNumber, последовательно напечатать числа от 1 до 10
// 2. распараллелить печать
// 3. изменить программу так, чтоб одновременно выводились 5 чисел

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	maxGoroutines := 5
	end := 10
	cur := 1

	for cur <= end {
		aliveRoutines := 0
		for aliveRoutines < maxGoroutines {
			aliveRoutines++

			wg.Add(1)
			go func(i int) {
				printNumber(i)
				wg.Done()
			}(cur)
			cur++
		}
		wg.Wait()
	}

}

func printNumber(n int) {
	time.Sleep(time.Second)
	fmt.Println(n)
}
