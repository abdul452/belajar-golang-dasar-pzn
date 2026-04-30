package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	// address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	// punya memory address1
	// address2 := address1 // copy dari address1 ke address2, pass by value
	// bikin memory baru namanya address2
	// address2.City = "Jakarta"

	// fmt.Println(address1)
	// fmt.Println(address2)

	var address1 Address = Address{"Bandung", "Jawa Barat", "Indonesia"}
	// bikin memory baru address1
	var address2 *Address = &address1 // pass by reference, address2 menyimpan alamat dari address1
	// var address2 adalah nama lain untuk akses ke address1
	// jadi jika ingin mengubah address1, bisa menggunakan address2 atau langsung address1
	address2.City = "Jakarta"
	fmt.Println(address1)
	fmt.Println(address2)

	address1.City = "Solo"
	fmt.Println(address1)
	fmt.Println(address2)
}
