package internal

import "fmt"

// hanya func init() yang bisa dijalankan otomatis ketika package ini diimport, jadi kita bisa gunakan untuk inisialisasi package
func init() {
	fmt.Println("init internal")
}
