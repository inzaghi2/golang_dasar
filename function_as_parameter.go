package main

import "fmt"

func sayHelloWithFillter(name string, fillter func(string) string) {
	fillteredName := fillter(name)
	fmt.Println("hello", fillteredName)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFillter("rifqi", spamFilter)
	fillter := spamFilter
	sayHelloWithFillter("anjing", fillter)

}
