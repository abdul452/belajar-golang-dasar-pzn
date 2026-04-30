package main

import (
	"belajar-golang-dasar/36-package-initialization/database"
	// panggil package internal, hanya untuk menjalankan func init() nya saja, jadi kita gunakan blank identifier
	_ "belajar-golang-dasar/36-package-initialization/internal"
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
