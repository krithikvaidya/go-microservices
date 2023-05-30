package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string, age int) Person {
	p := Person{name: name, age: age}
	return p
}

func main() {
	p1 := newPerson("somename", 20)

	fmt.Println(p1)
}

func (p Person) String() string {
	return fmt.Sprintf("{name: %s, age: %d}", p.name, p.age)
}
