package main

import "fmt"

type Country struct {
	negara string
}

func ChangedCountry(country *Country) {
	country.negara = "Indonesia"
}

func main() {
	country := Country{}
	ChangedCountry(&country)

	fmt.Println(country)
}
