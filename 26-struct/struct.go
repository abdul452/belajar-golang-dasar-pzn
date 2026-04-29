package main

import "fmt"

type Customer struct { // nama struct harus diawali dengan huruf besar agar bisa diakses dari package lain
	Name    string // nama field harus diawali dengan huruf besar agar bisa diakses dari package lain
	Address string
	Age     int
}

// ini struct method, method adalah function yang memiliki receiver, receiver adalah parameter yang digunakan untuk mengakses field struct
func (customer Customer) sayHello(name string) {
	fmt.Printf("Hello %s, my name is %s\n", name, customer.Name)
}

// Gunakan pointer (*) agar bisa mengubah data aslinya
func (c *Customer) ChangeName(newName string) {
	c.Name = newName // Perubahan ini akan berdampak ke data asli
}

func main() {
	var customer Customer
	customer.Name = "Eko"
	customer.Address = "Subang"
	customer.Age = 30

	fmt.Println(customer)
	fmt.Println(customer.Name)
	fmt.Println(customer.Address)
	fmt.Println(customer.Age)

	// struct literal
	customer2 := Customer{
		Name:    "Budi",
		Address: "Jakarta",
		Age:     25,
	}
	fmt.Println(customer2)

	// struct literal tanpa nama field
	customer3 := Customer{"Joko", "Bandung", 35}
	fmt.Println(customer3)

	customer.sayHello("Budi")

	customer.ChangeName("Joko")
	fmt.Println(customer.Name)
	customer.sayHello("Budi")
}
