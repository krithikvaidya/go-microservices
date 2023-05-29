package main

import "fmt"

func main() {
	// functions as first citizen example
	v := calc(2, 3, sum)
	fmt.Println(v)

	v = calc(2, 3, multiply)
	fmt.Println(v)
}

func calc(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func multiply(a, b int) int {
	return a * b
}

func sum(a, b int) int {
	return a + b
}
