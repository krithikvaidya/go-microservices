## Maps

Maps allow you to store items similar to a dictionary. You can think of the `key` as the word and the `value` as the definition. Go has a built in HashMap type, called a `map`.

```go
// defines a map having key of type string and values of type int
var m map[string]int
```

Just like making a `slice`, making a `map` can be accomplished with the built-in `make`.

```go
package main

import "fmt"

func main() {
  m := make(map[string]int)
  fmt.Println(m)
}
```

## Adding values to a map
Adding values to a map looks similar to assigning a value to a slice element:

```go
package main

import "fmt"

func main() {
  days := make(map[int]string)
  days[1] = "Monday"
  days[2] = "Tuesday"
  days[3] = "Wednesday"
  days[4] = "Thursday"
  days[5] = "Friday"
  days[6] = "Saturday"
  days[7] = "Sunday"
  fmt.Println(days)
}
```
If an entry already exits with that key, it will be overwritten.

## Compact literal initalisation
Just like slices, maps support compact literal initalisation, which declares and initalises the
map.

```go
package main

import "fmt"

func main() {
  days := map[int]string{
    1: "Monday",
    2: "Tuesday",
    3: "Wednesday",
    4: "Thursday",
    5: "Friday",
    6: "Saturday",
    7: "Sunday",
  }
  fmt.Println(days)
}
```

## Retrieving values from a map

Just like a slice, you can retrieve the value stored in a map with the syntax `m[key]`.

If it is present the value will be returned, if not the `zero value` will be returned.

```go
package main

import "fmt"

func main() {
  // map of planets to their number of moons
  moons := map[string]int {
    "Mercury": 0,
    "Venus": 0,
    "Earth": 1,
    "Mars": 2,
    "Jupiter": 67,
  }
  fmt.Println("Earth:", moons["Earth"])
  fmt.Println("Neptune:", moons["Neptune"])
}
```

## Checking if a map value exists
In the previous slide we saw that `moons["Neptune"]` returned `0`.

How can we tell if Neptune actually has no moons, or if 0 was returned because there is no entry for Neptune?

Map lookup support a second syntax:

```go
package main

import "fmt"

func main() {
  moons := map[string]int{"Mercury": 0, "Venus": 0, "Earth": 1, "Mars": 2, "Jupiter": 67}

  if n, ok := moons["Earth"]; ok {
    fmt.Println("Earth:", n, "Found:", ok)
  }

  n, found := moons["Neptune"]
  fmt.Println("Neptune:", n, "Found:", found)
}
```

## Deleting a value from map

To delete a value from a map, you use the built in `delete` function.

```go
package main

import "fmt"

func main() {
  moons := map[string]int{"Mercury": 0, "Venus": 0, "Earth": 1, "Mars": 2, "Jupiter": 67}
  const planet = "Mars"
  
  n, found := moons[planet]
  fmt.Println(planet, n, "Found:", found)
  delete(moons, planet)
  
  n, found = moons[planet]
  fmt.Println(planet, n, "Found:", found)
}
```

## Iterating over a map
If we wanted to print out all the values in a map we can use `range`.

```go
package main

import "fmt"

func main() {
  cities := map[string]int{
    "Tokyo": 42100000, "Delhi": 19500000, "Sau Paulo": 2300000, "New York": 18600000, "Moscow": 2150000,
  }
  
  for n, p := range cities {
    fmt.Println("City:", n, "Population", p)
  }
}
```

## Lab :

Assign every lowercase letter a value, from `1` for `a` to `26` for `z`. Given a string of lowercase letters, find the sum of the values of the letters in the string.

Write the program using maps:

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
  // create a map to keep alphabets from a to z
  letterValue := map[string]int{}
  letterValue[""] = 0
  for j, pos := 1, 'a'; pos <= 'z'; pos++ {
    char := fmt.Sprintf("%c", pos)
    letterValue[char] = j
    j++
  }

  // iterating over letters and getting the value of alphabet from map and doing sum
  sum := 0
  for _, char := range letters {
    c := fmt.Sprintf("%c", char)
    sum = sum + letterValue[c]
  }
  fmt.Println("sum is: ", sum)
}

```
</details>
<br>