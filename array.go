package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "rifqi"
	names[1] = "inzaghi"
	names[2] = "s.kom"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//array langsung
	var values = [3]int{
		90,
		80,
		70,
	}
	fmt.Println(values)

}
