package main

import "fmt"

func main() {
	var name string

	name = "Belajar Golang"
	fmt.Println("name =", name)

	name = "Belajar Golang Dasar"
	fmt.Println("name =", name)

	var friendName = "Belajar Golang"
	fmt.Println("friendName =", friendName)

	var age = 20
	fmt.Println("age =", age)

	country := "Indonesia"
	fmt.Println("country =", country)

	var (
		firstName = "Abdul"
		lastName  = "Salim"
	)
	fmt.Println("firstName =", firstName)
	fmt.Println("lastName =", lastName)
}
