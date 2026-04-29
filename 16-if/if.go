package main

import "fmt"

func main() {
	name := "Abdul"

	if name == "Abdul" {
		fmt.Println("Nama adalah Abdul")
	} else if name == "Budi" {
		fmt.Println("Nama adalah Budi")
	} else {
		fmt.Println("Nama bukan Abdul atau Budi")
	}

	if length := len(name); length > 5 { // short statement
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama cukup pendek")
	}
}
