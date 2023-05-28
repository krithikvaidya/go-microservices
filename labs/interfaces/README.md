## Interfaces

Go is an object oriented language; we have methods on types, but Go does not support inheritance or sub-classes. Go supports polymorphism with the help of interfaces.

In Go, an interface is a set of method signatures. When a type provides definition for all the methods in the interface, it is said to implement the interface.

Go supports Duck Typing. "If it walks like a duck and it quacks like a duck, then it must be a duck". 

Duck typing is a concept related to dynamic typing, where the type or the class of an object is less important than the methods it defines. When you use duck typing, you do not check types at all. Instead, you check for the presence of a given method or attribute.

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