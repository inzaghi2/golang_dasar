package main

import (
	"fmt"
)

func stopApp() {
	fmt.Println("applikasi selesai")
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApp(error bool) {
	defer stopApp()
	if error {
		panic("APLIKASI ERROR")
	}
	// contoh salah harusnya di masukan kedalam defer agar pesan itu muncul jika terjadi error
	// message := recover()
	// fmt.Println("terjadi panic", message)
}

func main() {
	runApp(true)

}
