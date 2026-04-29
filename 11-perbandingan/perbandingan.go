package main

import "fmt"

func main() {
	var a = "Abdul"
	var b = "Abdul"

	var result bool = a == b
	fmt.Println("Apakah a sama dengan b?", result)

	result = a != b
	fmt.Println("Apakah a tidak sama dengan b?", result)

	var c = 10
	var d = 20

	var result2 bool = c < d
	fmt.Println("Apakah c lebih kecil dari d?", result2)
}
