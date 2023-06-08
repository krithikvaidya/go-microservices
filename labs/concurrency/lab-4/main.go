package main

import (
	"time"
)

func main() {

}

func getMessage(ch chan string) {
	for i := 1; true; i++ {
		ch <- time.Now().String()
		time.Sleep(time.Millisecond * 500)
	}
}
