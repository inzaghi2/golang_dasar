package main

import "fmt"

func logging() {
	fmt.Println("selesia memanggil function")

}
func runApplication() {
	defer logging()
	fmt.Println("run application")
}

func main() {

}
