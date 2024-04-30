package main

import (
	"fmt"
	"sync"
)

type SafeNumber struct {
	shared int
	mu     sync.Mutex
}

func (s *SafeNumber) Store(num int) {
	s.mu.Lock()
	s.shared = num
	s.mu.Unlock()
}

func (s *SafeNumber) Add(num int) {
	s.mu.Lock()
	s.shared += num
	s.mu.Unlock()
}

func (s *SafeNumber) Load() int {
	s.mu.Lock()
	result := s.shared
	s.mu.Unlock()

	return result
}

func main() {
	shared := SafeNumber{
		shared: 0,
		mu:     sync.Mutex{},
	}

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
