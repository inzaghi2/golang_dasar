package main

import "fmt"

func main() {
	for index := 0; index <= 10; index++ {
		if index == 5 {
			break

		}
		fmt.Println("perulangan ke", index)

	}

	for num := 0; num < 10; num++ {
		if num%2 == 0 {
			continue
		}
		fmt.Println("perulangan ke", num)

	}
}
