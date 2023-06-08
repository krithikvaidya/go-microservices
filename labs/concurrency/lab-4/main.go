package main

import (
	"time"
)

func main() {

}

func getMessage(ch chan string) {
	for i := 1; i < 5; i++ {
		ch <- time.Now().String()
		time.Sleep(time.Millisecond * 500)
	}
}
