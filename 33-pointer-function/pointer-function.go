package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func changeCityToJakarta(address *Address) { // disini pakai *
	address.City = "Jakarta"
}

func main() {
	address := Address{"Bandung", "Jawa Barat", "Indonesia"}

	// disini kita pakai &
	changeCityToJakarta(&address)
	fmt.Println(address)
}
