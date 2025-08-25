package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "rifqi inzaghi",
		"address": "semarang",
	}

	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "buku golang"
	book["author"] = "rifqi inzaghi"
	book["ups"] = "salah"
	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)

}
