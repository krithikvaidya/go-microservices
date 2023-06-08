package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	var max int = 10e7
	// calculateSumOfNumbers(max)
	calculateConcurrently(max)
}

func calculateConcurrently(max int) {
	s := GenerateNumbers(max)
	t := time.Now()
	ch := make(chan int64)
	numCpu := runtime.NumCPU()

	sizeOfParts := max / numCpu
	for i := 0; i < numCpu; i++ {
		start := i * sizeOfParts
		end := start + sizeOfParts

		part := s[start:end]
		go sum(part, ch)
	}
	var total int64
	for i := 0; i < numCpu; i++ {
		total += <-ch
	}
	fmt.Printf("With Channel Add, Sum: %d,  Time Taken: %s\n", total, time.Since(t))
}

func calculateSumOfNumbers(max int) {
	s := GenerateNumbers(max)
	t := time.Now()

	part1 := s[:len(s)/2]
	part2 := s[len(s)/2:]
	ch := make(chan int64)

	go sum(part1, ch)
	go sum(part2, ch)

	x := <-ch
	y := <-ch
	total := x + y

	fmt.Printf("Channel Add, Sum: %d,  Time Taken: %s\n", total, time.Since(t))
}

func sum(part []int, ch chan int64) {
	var sum int64
	for _, v := range part {
		sum += int64(v)
	}
	ch <- sum
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateNumbers - random number generation
func GenerateNumbers(max int) []int {
	numbers := make([]int, max)
	for i := 0; i < max; i++ {
		numbers[i] = rand.Intn(10)
	}
	return numbers
}
