package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Alice",
		"country": "USA",
	}

	fmt.Println("Nama:", person["name"])
	fmt.Println("Negara:", person["country"])
	fmt.Println(person)

	// nambahin data ke map
	person["age"] = "30"
	fmt.Println("Usia:", person["age"])

	// Deklarasi map dengan tipe data string sebagai key dan int sebagai value
	ages := make(map[string]int)

	// Menambahkan data ke dalam map
	ages["Alice"] = 30
	ages["Bob"] = 25

	// Mengakses data dari map
	aliceAge := ages["Alice"]
	bobAge := ages["Bob"]

	fmt.Println("Usia Alice:", aliceAge)
	fmt.Println("Usia Bob:", bobAge)

	// Menghapus data dari map
	delete(ages, "Alice")

	books := make(map[string]string)
	books["judul"] = "Belajar Golang"
	books["penulis"] = "Abdul"
	books["ups"] = "salah"
	fmt.Println(books)
	delete(books, "ups")
	fmt.Println(books)
}
