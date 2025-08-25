package main

import "fmt"

func main() {
	name := "rifqi inzaghi"

	switch name {
	case "rifqi":
		fmt.Println("hello rifqi")
	case "inzaghi":
		fmt.Println("hello inzaghi")
	default:
		fmt.Println("tidak kenal")
	}

	switch lenghth := len(name); lenghth > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	default:
		fmt.Println("nama sudah benar")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")

	}
}
