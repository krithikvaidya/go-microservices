package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go getMessage(ch)

	for msg := range ch {
		// msg, isOpen := <-ch
		// if !isOpen {
		// 	break
		// }
		fmt.Println(msg)
	}
}

func getMessage(ch chan string) {
	defer close(ch)
	for i := 1; i <= 5; i++ {
		ch <- time.Now().String()
		time.Sleep(time.Millisecond * 500)
	}

	time.Sleep(time.Second * 5)
}
