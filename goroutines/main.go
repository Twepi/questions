package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	total := 0
	count := 0
	for i := 0; i < 1000; i++ {
		total += generateRandomNumber()
		count++
	}

	elapsed := time.Since(start)

	fmt.Printf("Count: %d\n", count)
	fmt.Printf("Total: %d\n", total)
	fmt.Printf("Elapsed Time: %s\n", elapsed)
}

func generateRandomNumber() int {
	time.Sleep(3 * time.Millisecond)
	return rand.Intn(100)
}
