package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke", counter)
		counter++
	}
	fmt.Println("selesai")

	for code := 1; code <= 10; code++ {
		fmt.Println("code ke", code)
	}
	fmt.Println("selesai")
}
