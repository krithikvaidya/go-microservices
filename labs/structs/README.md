## Structs

Go supports what we call *compound* types, that is, types that are *composed* of other types.  They’re useful for grouping data together.

These are called struct (for structure) types. We declare a struct like this:

```go
type Point struct {
  X int
  Y int
}
```

Point is a position in two dimensional space, it has two fields, `X` and `Y`.

Another example:

```go
type person struct {
  name string
  age  int
}

func newPerson(name string) *person {
  p := person{name: name}
  p.age = 42
  return &p
}

func main() {
  // This syntax creates a new struct.
  r := person{name: "Rob", age: 50}

  // Access struct fields with a dot.
  fmt.Println(r.name)

  r.age = 51
  fmt.Println(r.age)

  // It’s idiomatic to encapsulate new struct creation in constructor functions
  fmt.Println(newPerson("John"))
}
```

## String method

Any type that implements a `String() string` method will be used by the `fmt` package
when it prints the value.

## Lab:

You will build the [ToDo project](../project) using structs