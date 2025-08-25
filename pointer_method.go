package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	alias := Man{Name: "rifqi"}
	alias.Married()

	fmt.Println(alias.Name)
}
