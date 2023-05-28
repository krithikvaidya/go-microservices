## Arrays
Arrays allow you to store multiple elements of the same type in a variable in a particular order. Arrays in Go have fixed capacity which you define when you declare the variable. We can initialize an array in two ways:

- `[N]type{value1, value2, ..., valueN}` 
  
  e.g. `numbers := [5]int{1, 2, 3, 4, 5}`

- `[...]type{value1, value2, ..., valueN}` 

  e.g. `numbers := [...]int{1, 2, 3, 4, 5}`

Array's in Go are not frequently used, their most common purpose in Go is to hold storage for a slice.

## Range
`range` lets you iterate over an array. On each iteration, range returns two values - the `index` and the `value`. In the below example we are choosing to ignore the index value by using _ [blank identifier](../identifiers##unused-variable-declarations--blank-identifier).


```go
numbers := [5]int{1, 2, 3, 4, 5}
for _, v := range numbers {
  fmt.Println(v)
}
```

## Slices
A slice is an ordered collection of values of a single type. The syntax for declaring a slice variable is very similar to declaring a scalar variable.

```go
var i int   // an int called i
var j []int // a slice of ints called j
```
In this example,
- `i` is a variable of type `int`.
- `j` is a variable of type `[]int`, that is, a `slice` of `int`.

Slices are very important in Go programs

## How large is a slice?

If I declare a slice, `[]string`, how many items can it hold?

The zero value of a slice is empty, that is, it has a length of zero; it can hold 0 items.

```go
package main

import "fmt"

func main() {
  var s []string
  fmt.Println(len(s))
}
```
We can retrieve the length of a slice with the built-in `len` function.

## Making a slice
We can create a slice with space to hold items using the built-in make function.

```go
package main

import "fmt"

func main() {
  var cities []string
  cities = make([]string, 20)
  fmt.Println(len(s))
}
```

In this example, on the first line `var cities []string` declares cities to be a slice of string. On the second line, cities is assigned the result of `make([]string, 20)`.

## Lab:

Let's do a quick exercise to familarise yourself with using slices

```go
package main

import "fmt"

func main() {
  // Declare a variable called i which is a slice of 5 int.
  // Declare a variable called f which is a slice of 9 float64.
  // Declare a variable called s which is a slice of 4 string.

  fmt.Println(len(i), len(f), len(s))
}
```

- Does your program print the expected result, 5 9 4?
<details>
  <summary>Not sure how?</summary>

  ```go
    func main() {
      var i []int
      var f []float64
      var s []string
      
      i = make([]int, 5)
      f = make([]float64, 9)
      s = make([]string, 4)

      fmt.Println(len(i), len(f), len(s))
    }  
  ```
</details>
<br>

## Index expressions
- To access, or assign, the contents of a slice element at index `i`, use the form `s[i]`.

- Slices are zero indexed, so `s[0]` is the 1st element, `s[1]` is the second element, and so on.

- When the index expression appears on the left hand side of the equals operator, =

  `s[7] = 20`

- We are assigning the number 20 to the 8'th element of the slice.

- When the index expression appears on the right hand side of the equals operator, =

  `x := s[7]`

- We are assigning the value at the 8th element of `s` to the variable `x`.

## Slice zero value
We saw earlier that the zero value of the slice

`var s []int`

is an empty slice, a slice with length of zero.

What is the value of each of the elements of a newly created, with make, slice?

```go
package main

import "fmt"

func main() {
  x := make([]int, 5)

  for i := 0; i < len(x); i++ {
    fmt.Println(x[i])
  }

}
```

## Slice initialisation
We want to create an `[]int` slice of the first 10 prime numbers, how could we do this?

One solution could be to create the slice and assign a value to each element in the slice.

```go
package main

import "fmt"

func main() {
  primes := make([]int, 10)
  primes[0] = 2
  primes[1] = 3
  primes[2] = 5
  primes[3] = 7
  primes[4] = 11
  primes[5] = 13
  primes[6] = 17
  primes[7] = 19
  primes[8] = 23
  primes[9] = 29
  fmt.Println(primes)
}
```
Doing this manually is verbose and boring; how would you do this for the first 50 primes?

Go supports a method of assignment where we both declare and initalise the slice at once.

```go
package main

import "fmt"

func main() {
  
  primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

  fmt.Println(primes)
}
```
This is called the *composite literal* syntax.

## append
So far we've been using slices with a known length. You can extend the contents of a slice
with the built-in append function.

```go
package main

import "fmt"

func main() {
  primes := []int{2, 3, 5, 7, 11}
  fmt.Println(len(primes), primes)

  primes = append(primes, 13)  
  fmt.Println(len(primes), primes)

  primes = append(primes, 17, 19, 23)
  fmt.Println(len(primes), primes)
}
```
`append` increases the length of the slice to accommodate the new items, then returns a
new slice value.

You can `append` multiple values in one statement, providing they are all the same type.

## Further reading

[Arrays, slices (and strings): The mechanics of 'append'](https://go.dev/blog/slices)

## Subslices

```go
package main

import "fmt"

func main() {
  brothers := []string{"chico", "harpo", "groucho", "gummo", "zeppo"}
  fmt.Println(brothers)

  wellknown := brothers[0:3]
  fmt.Println(wellknown)
}
```
The expression brothers[0:3] evaluates to a slice of the 1st to 3rd Marx brother.

**Lab 1:**

```go

package main

import "fmt"

func main() {
	// These are the primes less than 200
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
		31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
		73, 79, 83, 89, 97, 101, 103, 107, 109,
		113, 127, 131, 137, 139, 149, 151, 157,
		163, 167, 173, 179, 181, 191, 193, 197, 199}
	fmt.Println(primes)

	// Write a program to print all the primes less than 10
	// loop through the slice of primes and test if the value
	// is less than 10. When you find a value that is 10 or more,
	// slice the list of primes at that point and print it.

	// Bonus: write a program to print only the two digit primes.
}

```

<details>
  <summary>Not sure how?</summary>


```go
package main

import "fmt"

func main() {
	// These are the primes less than 200
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
		31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
		73, 79, 83, 89, 97, 101, 103, 107, 109,
		113, 127, 131, 137, 139, 149, 151, 157,
		163, 167, 173, 179, 181, 191, 193, 197, 199}
	fmt.Println(primes)

	for i := 0; i < len(primes); i++ {
		if primes[i] > 9 {
			v := primes[0:i]
			fmt.Println(v)
			break
		}
	}

	// Bonus: write a program to print only the two digit primes.
	var i, j int
	for ; i < len(primes); i++ {
		if primes[i] > 9 {
			break
		}
	}
	for j = i; j < len(primes); j++ {
		if primes[j] > 99 {
			break
		}
	}
	fmt.Println(primes[i:j])
}
```
</details>
<br>

**Lab 2:**

Assign every lowercase letter a value, from `1` for `a` to `26` for `z`. Given a string of lowercase letters, find the sum of the values of the letters in the string.

```
sumOfLetters("")          => 0
sumOfLetters("a")         => 1
sumOfLetters("z")         => 26
sumOfLetters("cab")       => 6
sumOfLetters("excellent") => 100
```

<details>
  <summary>Not sure how?</summary>

```go
package main

import "fmt"

func main() {
  // find the sum of "cba"
  letters := "cba"
  sumOfLetters(letters)
}

func sumOfLetters(letters string) {
  // create a slice to keep alphabets from a to z
  letterValue := make([]string, 0)
  letterValue = append(letterValue, "")
  for pos := 'a'; pos <= 'z'; pos++ {
    char := fmt.Sprintf("%c", pos)
    letterValue = append(letterValue, char)
  }

  // iterating over letters and calling findValue to find the value of alphabet and doing sum
  sum := 0
  for _, char := range letters {
    c := fmt.Sprintf("%c", char)
    sum = sum + findValue(c, letterValue)
  }
  fmt.Println("sum is: ", sum)
}

// find letter value
func findValue(char string, letterValue []string) int {
  for i, v := range letterValue {
    if char == v {
      return i
    }
  }
  return 0
}

```
</details>
<br>