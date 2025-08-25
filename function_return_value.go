package main

import "fmt"

func pengembalian(name string) string {
	return "hello " + name

}

func main() {
	result := pengembalian("rifqi inzaghi")
	fmt.Println(result)

	fmt.Println(pengembalian("rifqi inzaghi"))
}
