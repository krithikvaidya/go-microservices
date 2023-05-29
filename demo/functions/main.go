package main

import "fmt"

func swap(a, b int, name string) (int, int) {
	return b, a
}

func Write(b []byte) (numberOfBytes int, err error) {
	return 0, nil
}

func sum(a, b int) int {
	return a + b
}

func print(msg string) {
	fmt.Println(msg)
}

func main() {
	defer releaseResources()
	fmt.Println("testing defer statement")
}

func releaseResources() {
	fmt.Println("closing and releasinng connections....")
}
