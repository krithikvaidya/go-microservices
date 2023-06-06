package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var max int = 10e7
	calculateSumOfNumbers(max)
}

func calculateSumOfNumbers(max int) {
	s := GenerateNumbers(max)
	t := time.Now()
	var sum int64
	for _, v := range s {
		sum += int64(v)
	}

	fmt.Printf("Without channel Add, Sum: %d,  Time Taken: %s\n", sum, time.Since(t))
}

// GenerateNumbers - random number generation
func GenerateNumbers(max int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, max)
	for i := 0; i < max; i++ {
		numbers[i] = rand.Intn(10)
	}
	return numbers
}
