package main

import "fmt"

type User struct {
	Nama, Address string
	Age           int
}

func (user User) sayHello() {
	fmt.Println("hello", user.Nama)

}

func main() {
	Rully := User{Nama: "rifqi inzaghi"}

	Rully.sayHello()
}
