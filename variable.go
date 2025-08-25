package main

import "fmt"

func main() {
	//menginisilaisasi tipe data
	var name string

	name = "rifqi"
	fmt.Println(name)

	name = "rifqi inzaghi"
	fmt.Println(name)

	//tanpa menentukan tipe data
	var inisial = "rifqi inzaghi"
	fmt.Println(inisial)

	//mnyerdahanakan lagi menggunakan (:=) hanya bisa pertama kali saja
	//dan tidak bisa diulang
	username := "Rifqi Inzaghi S.KOM"
	fmt.Println(username)

	//deklarasi multiple variable
	var (
		firstName = "rifqi"
		lastName  = "inzaghi"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)

}
