package main

import "fmt"

// channel and buffered channel example
func main() {
	ch := make(chan string)

	ch <- "hello"

	msg := <-ch
	fmt.Println(msg)
}
