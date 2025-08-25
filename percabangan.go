package main

import "fmt"

func main() {
	name := "rifqi inzaghi"
	fmt.Println(name)

	if name == "rifqi inzaghi" {
		fmt.Println("hello rifqi inzaghi")
	} else {
		fmt.Println("tidak kenal")
	}

	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}
