## Interfaces

Go is an object oriented language; we have methods on types, but Go does not support inheritance or sub-classes. Go supports polymorphism with the help of interfaces.

In Go, an interface is a set of method signatures. When a type provides definition for all the methods in the interface, it is said to implement the interface.

Go supports Duck Typing. "If it walks like a duck and it quacks like a duck, then it must be a duck". 

Duck typing is a concept related to dynamic typing, where the type or the class of an object is less important than the methods it defines. When you use duck typing, you do not check types at all. Instead, you check for the presence of a given method or attribute.


## Example 1:

Lets say we have an interface A and we want to implement it to struct B

```go
type A interface {
  foo() int64
}

type B struct {
  num int64
}
```

Go supports duck typing. So to make B implement A we have to make B look like A. 
That means, if you define a receiver function foo() on B, it will look like A.

```go
func (b B) foo() int64 {
  return 10
}

func main() {
  var v A = B{2} // Here B implements interface A?
  fmt.Println(v.foo())
}
```

You can also have another struct C

```go
type C struct {
}

// and now C also implements A
func (c C) foo() int64 {
  return 20
}

func main() {
  var v A = B{2} // Here B implements interface A?
  fmt.Println(v.foo())
  v = C{}
  fmt.Println(v.foo())
}
```


## Example 2:

```go
package main

import "fmt"

type Repository interface {
  FindAll()
}

type DbRepository struct {}
type StubRepository struct {}

func (r DbRepository) FindAll() {
  fmt.Println("I talk to DB and will get data from DB")
}

func (r StubRepository) FindAll() {
  fmt.Println("I am just a stubbed response, I don't talk to DB")
}

func main() {
  dbRepo := DbRepository{}
  getData(dbRepo)

  stubRepo := StubRepository{}
  getData(stubRepo)
}

func getData(repository Repository) {
  repository.FindAll()
}
```