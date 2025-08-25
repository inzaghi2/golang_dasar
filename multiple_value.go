package main

import "fmt"

func getFullName() (string, string) {
	return "rifqi", "inzaghi"

}

func main() {
	fmt.Println(getFullName())

	// firstName, lastName := getFullName()
	// fmt.Println(firstName, lastName)

	firstName, _ := getFullName()
	fmt.Println(firstName)

}
