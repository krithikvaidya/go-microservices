package main

import "fmt"

// channel and buffered channel example
func main() {
	ch := make(chan string, 2)

	ch <- "hello"
	ch <- "world"
	msg := <-ch
	fmt.Println(msg)

	ch <- "third"

	msg = <-ch
	fmt.Println(msg)

	msg = <-ch
	fmt.Println(msg)
}
