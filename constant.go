package main

import "fmt"

func main() {
	//deklarasi konstanta
	//konstanta tidak bisa diubah
	const firstName string = "Rifqi"
	const lastName = "Inzaghi"

	//terjadi error jika diubah
	//firstName = "zaghi"

	//deklarasi secara multiple
	const (
		namaDepan    = "rifqi"
		namaBelakang = "inzaghi"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)

}

//konstan tidak bisa diubah variable bisa diubah
