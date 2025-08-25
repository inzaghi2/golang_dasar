package main

import "fmt"

type alamat struct {
	City, Province, Country string
}

func main() {
	var alamat1 *alamat = new(alamat)
	var alamat2 *alamat = alamat1

	alamat2.City = "Solo"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
