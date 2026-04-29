package main

import "fmt"

func main() {
	const firstName string = "Abdul"
	const lastName = "Salim"
	const age = 30

	// firstName

	fmt.Println("firstName =", firstName)
	fmt.Println("lastName =", lastName)
	fmt.Println("age =", age)

	const (
		country  = "Indonesia"
		province = "Jawa Barat"
	)
	fmt.Println("country =", country)
	fmt.Println("province =", province)
}
