package main

import "fmt"

func main() {
	var (
		nilaiA = 20
		nilaiB = 30
	)

	fmt.Println("penjumlahan = ", nilaiA+nilaiB)
	fmt.Println("pengurangan = ", nilaiA-nilaiB)

	//augmented assigments (versi lebih singkatnya)
	var index = 10
	index += 10
	fmt.Println(index)

	index -= 5
	fmt.Println(index)
	index *= 2
	fmt.Println(index)
	index /= 2
	fmt.Println(index)
	index %= 2
	fmt.Println(index)

	//unary operator

	var angka = 50
	angka++
	fmt.Println(angka)
	angka--
	fmt.Println(angka)
}
