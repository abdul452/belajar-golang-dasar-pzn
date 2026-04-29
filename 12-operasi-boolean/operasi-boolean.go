package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var nilaiAbsen = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusNilaiAbsen bool = nilaiAbsen > 80

	var lulus bool = lulusNilaiAkhir && lulusNilaiAbsen
	fmt.Println("Apakah lulus?", lulus)

	var lulusSalahSatu bool = lulusNilaiAkhir || lulusNilaiAbsen
	fmt.Println("Apakah lulus salah satu?", lulusSalahSatu)

	var tidakLulus bool = !lulus
	fmt.Println("Apakah tidak lulus?", tidakLulus)
}
