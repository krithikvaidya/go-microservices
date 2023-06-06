package main

import (
	"fmt"
	"time"
)

func main() {
	printMessage("hello")
	printMessage("world")
}

func printMessage(msg string) {
	for i := 1; true; i++ {
		fmt.Println(i, msg)
		time.Sleep(time.Millisecond * 500)
	}
}
