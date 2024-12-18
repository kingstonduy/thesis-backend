package main

import "fmt"

func main() {
	// Create
	var person Person
	person.Name = "John Doe"
	person.Age = 30

	// Print
	fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)

	// Update
	person.Age = 31
	fmt.Printf("Updated Age: %d\n", person.Age)

	// Delete
	person = Person{}
	fmt.Println(person.Name, person.Age)
}

type Person struct {
	Name string
	Age  int
}
