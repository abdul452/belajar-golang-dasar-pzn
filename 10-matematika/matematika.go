package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var c = 30
	var d = 40
	var e = 50

	var hasil = a + b - c + d*e
	fmt.Println("Hasil penjumlahan:", hasil)

	var i = 10
	i++ // i = i + 1
	fmt.Println("Nilai i setelah increment:", i)

	i-- // i = i - 1
	fmt.Println("Nilai i setelah decrement:", i)

	var sisa = 10 % 3 // sisa pembagian 10 dengan 3
	fmt.Println("Sisa pembagian 10 dengan 3:", sisa)

	i += 5 // i = i + 5
	fmt.Println("Nilai i setelah penambahan 5:", i)

	i -= 2 // i = i - 2
	fmt.Println("Nilai i setelah pengurangan 2:", i)

	i *= 3 // i = i * 3
	fmt.Println("Nilai i setelah perkalian dengan 3:", i)

	i /= 2 // i = i / 2
	fmt.Println("Nilai i setelah pembagian dengan 2:", i)
}
