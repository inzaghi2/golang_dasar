package main

import "fmt"

type UserName struct {
	Nama, Address string
	Age           int
}

func main() {
	user1 := UserName{
		Nama:    "rifqi inzaghi",
		Address: "Semarang",
		Age:     24,
	}
	user2 := &user1

	user2.Nama = "zaghi"

	*user2 = UserName{
		Nama:    "Monkey D luffi",
		Address: "East Blue",
		Age:     24,
	}

	fmt.Println(user1)
	fmt.Println(user2)
}
