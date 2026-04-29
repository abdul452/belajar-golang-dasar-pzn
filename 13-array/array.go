package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Abdul"
	names[1] = "Budi"
	names[2] = "Cici"

	fmt.Println("Nama pertama:", names[0])
	fmt.Println("Nama kedua:", names[1])
	fmt.Println("Nama ketiga:", names[2])

	// Menggunakan array literal
	numbers := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Angka pertama:", numbers[0])
	fmt.Println("Angka kedua:", numbers[1])
	fmt.Println("Angka ketiga:", numbers[2])
	fmt.Println("Angka keempat:", numbers[3])
	fmt.Println("Angka kelima:", numbers[4])

	// jika declare array dengan jumlah elemen yang tidak diketahui, bisa menggunakan ... untuk menghitung jumlah elemen secara otomatis
	colors := [...]string{"Merah", "Hijau", "Biru"}
	fmt.Println("Warna pertama:", colors[0])
	fmt.Println("Warna kedua:", colors[1])
	fmt.Println("Warna ketiga:", colors[2])
	fmt.Println(colors)

	// jika array kita declare 3 tapi diisi hanya 2 maka elemen yang tidak diisi akan bernilai default yaitu 0 untuk tipe data int, "" untuk tipe data string, dan false untuk tipe data bool
	var fruits [3]string
	fruits[0] = "Apel"
	fruits[1] = "Jeruk"
	fmt.Println("Buah pertama:", fruits[0])
	fmt.Println("Buah kedua:", fruits[1])
	fmt.Println("Buah ketiga (default):", fruits[2])
	fmt.Println(fruits)

	var angkaLain [5]int
	angkaLain[0] = 1
	angkaLain[1] = 2
	angkaLain[2] = 3
	fmt.Println("Angka lain pertama:", angkaLain[0])
	fmt.Println("Angka lain kedua:", angkaLain[1])
	fmt.Println("Angka lain ketiga:", angkaLain[2])
	fmt.Println("Angka lain keempat (default):", angkaLain[3])
	fmt.Println("Angka lain kelima (default):", angkaLain[4])
	fmt.Println(angkaLain)

	// Menggunakan array dengan tipe data yang berbeda
	var mixed [4]interface{}
	mixed[0] = "Hello"
	mixed[1] = 123
	mixed[2] = 3.14
	mixed[3] = true

	fmt.Println("Mixed pertama:", mixed[0])
	fmt.Println("Mixed kedua:", mixed[1])
	fmt.Println("Mixed ketiga:", mixed[2])
	fmt.Println("Mixed keempat:", mixed[3])

	// print len
	fmt.Println("Panjang array names:", len(names))
	fmt.Println("Panjang array numbers:", len(numbers))
	fmt.Println("Panjang array colors:", len(colors))
	fmt.Println("Panjang array fruits:", len(fruits))
	fmt.Println("Panjang array angkaLain:", len(angkaLain))
	fmt.Println("Panjang array mixed:", len(mixed))
}
