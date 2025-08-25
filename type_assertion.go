package main

import "fmt"

func random() interface{} {
	return "ok"

}

func main() {
	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)

	default:
		fmt.Println("Unknown")
	}

	/*
		resulString := result.(string)
		fmt.Println(resulString)

		resultInt := result.(int) //panic
		fmt.Println(resultInt)
	*/
}
