package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var shared atomic.Int64
	shared.Store(0) // store задает значение

	var wg sync.WaitGroup

	for i := 0; i < 100_000; i++ {
		wg.Add(1)
		go func() {
			shared.Add(1) // add добавляет к текущему значению
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("result: ", shared.Load()) // load возвращает текущее значение
}
