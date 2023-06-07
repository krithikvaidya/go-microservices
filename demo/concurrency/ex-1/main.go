package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go printMessage("hello", &wg)
	go printMessage("world", &wg)
	// go func() {
	// 	printMessage("hello")
	// 	wg.Done()
	// }()

	// go func() {
	// 	printMessage("world")
	// 	wg.Done()
	// }()

	wg.Wait() // blocking
}

func printMessage(msg string, wg *sync.WaitGroup) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, msg)
		time.Sleep(time.Millisecond * 500)
	}
	wg.Done()
}
