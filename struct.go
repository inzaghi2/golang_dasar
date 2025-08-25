package main

import "fmt"

type Customers struct {
	Nama, Address string
	Age           int
}

func main() {
	var rifqi Customers
	rifqi.Nama = "rifqi inzaghi"
	rifqi.Address = "Semarang"
	rifqi.Age = 24

	fmt.Println(rifqi.Nama)
	fmt.Println(rifqi.Address)
	fmt.Println(rifqi.Age)

}
