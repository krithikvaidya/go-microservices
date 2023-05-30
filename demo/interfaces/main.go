package main

import "fmt"

type Repository interface {
	FindAll()
}

type DbRepository struct{}
type StubRepository struct{}

func (r DbRepository) FindAll() {
	fmt.Println("I talk to DB and will get data from DB")
}

func (r StubRepository) FindAll() {
	fmt.Println("I am just a stubbed response, I don't talk to DB")
}

func main() {
	// dbRepo := DbRepository{}
	stub := StubRepository{}
	getData(stub)

}

func getData(repository Repository) {
	repository.FindAll()
}
