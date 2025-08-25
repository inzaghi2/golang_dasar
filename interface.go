package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func sayName(value HasName) {
	fmt.Println("Hello", value.GetName())

}

func (person Person) GetName() string {
	return person.Name
}
func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "rifqi"}
	sayName(person)

	animal := Animal{Name: "kucing"}
	sayName(animal)

}
