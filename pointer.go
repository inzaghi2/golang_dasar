package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{
		City:     "Semarang",
		Province: "Jawa Tengah",
		Country:  "Indonesia",
	}

	address2 := address1 //copy value

	//operator & sebagai penunjuk address1 bukan lagi mencopy data
	//address2 := &address1

	address2.City = "Solo"

	fmt.Println(address1) //tidak berubah datanya
	fmt.Println(address2) //data city berubah dari semarang menjadi solo

}
