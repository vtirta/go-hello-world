package main

import "fmt"

type Person struct {
	ID   string
	name string
}

func main() {
	fmt.Println("Hello World!")

	var person Person
	person.ID = "123"
	person.name = "Vic"

	fmt.Println(person)
}
