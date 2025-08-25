package main

import "fmt"

func endApp() {
	fmt.Println("applikasi selesai")

}
func startApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}

}

func main() {
	startApp(true)
	fmt.Println("rifqi")
}
