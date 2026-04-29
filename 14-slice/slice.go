package main

import (
	"fmt"
)

func main() {
	names := []string{"Abdul", "Budi", "Cici", "Dedi", "Eka", "Gina"}

	slices1 := names[4:6] // Mengambil elemen dari index 0 hingga 1 (tidak termasuk index 2)
	fmt.Println("Slicing 1:", slices1)

	slices2 := names[2:6] // Mengambil elemen dari index 1 hingga 2 (tidak termasuk index 3)
	fmt.Println("Slicing 2:", slices2)

	slices3 := names[:2] // Mengambil elemen dari index 0 hingga 1 (tidak termasuk index 2)
	fmt.Println("Slicing 3:", slices3)

	slices4 := names[1:] // Mengambil elemen dari index 1 hingga akhir
	fmt.Println("Slicing 4:", slices4)

	slices5 := names[:] // Mengambil semua elemen
	fmt.Println("Slicing 5:", slices5)

	// Slicing dengan array
	numbers := [5]int{10, 20, 30, 40, 50}
	slices6 := numbers[1:4] // Mengambil elemen dari index 1 hingga 3 (tidak termasuk index 4)
	fmt.Println("Slicing dengan array:", slices6)

	// Slicing dengan tipe data yang berbeda
	mixed := []interface{}{"Hello", 123, 3.14, true}
	slices7 := mixed[0:2] // Mengambil elemen dari index 0 hingga 1 (tidak termasuk index 2)
	fmt.Println("Slicing dengan tipe data berbeda:", slices7)

	// Slicing dengan boolean
	slices8 := []bool{true, false, true}
	fmt.Println("Slicing dengan boolean:", slices8)

	// append slice
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:7]       // Mengambil elemen dari index 5 hingga 6 (tidak termasuk index 7)
	fmt.Println(daySlice1)       // Output: [Sabtu Minggu]
	daySlice1[0] = "Sabtu Baru"  // Mengubah elemen pada slice
	daySlice1[1] = "Minggu Baru" // Mengubah elemen pada slice
	fmt.Println(daySlice1)       // Output: [Sabtu Baru Minggu Baru]
	fmt.Println(days)            // Output: [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]

	// Menggunakan append untuk menambahkan elemen ke slice
	daySlice2 := append(daySlice1, "Libur") // Menambahkan elemen "Libur" ke slice
	daySlice2[0] = "Sabtu Terbaru"          // Mengubah elemen pada slice
	fmt.Println(daySlice1)                  // Output: [Sabtu Baru
	fmt.Println(daySlice2)                  // Output: [Sabtu Terbaru Minggu Baru Libur]
	fmt.Println(days)                       // Output: [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]

	var newSlice []string = make([]string, 2, 5) // Membuat slice dengan panjang 2 dan kapasitas 5
	newSlice[0] = "Hello"
	newSlice[1] = "World"
	// newSlice[2] = "!" // Akan menyebabkan panic karena panjang slice hanya 2, meskipun kapasitasnya 5
	fmt.Println("New Slice:", newSlice)
	fmt.Println("Panjang New Slice:", len(newSlice))
	fmt.Println("Kapasitas New Slice:", cap(newSlice))

	// Menambahkan elemen ke newSlice
	newSlice = append(newSlice, "Golang")
	fmt.Println("New Slice setelah append:", newSlice)
	fmt.Println("Panjang New Slice setelah append:", len(newSlice))
	fmt.Println("Kapasitas New Slice setelah append:", cap(newSlice))

	newSlice = append(newSlice, "test")
	newSlice = append(newSlice, "test2")
	newSlice = append(newSlice, "test3")
	fmt.Println("New Slice setelah append lagi:", newSlice)
	fmt.Println("Panjang New Slice setelah append lagi:", len(newSlice))
	fmt.Println("Kapasitas New Slice setelah append lagi:", cap(newSlice))

	// copy slice
	sourceSlice := []string{"A", "B", "C", "D", "E"}
	destinationSlice := make([]string, len(sourceSlice))
	copy(destinationSlice, sourceSlice)
	fmt.Println("Source Slice:", sourceSlice)
	fmt.Println("Destination Slice:", destinationSlice)

	// Mengubah elemen pada sourceSlice
	sourceSlice[0] = "X"
	fmt.Println("Source Slice setelah diubah:", sourceSlice)
	fmt.Println("Destination Slice setelah source diubah:", destinationSlice) // Destination tetap tidak berubah karena copy membuat salinan baru

	// beda array dengan slice
	// array memiliki ukuran tetap, sedangkan slice memiliki ukuran dinamis
	// array disimpan di stack, sedangkan slice disimpan di heap
	// array tidak bisa diubah ukurannya, sedangkan slice bisa diubah ukurannya dengan append
	// array tidak bisa digunakan sebagai parameter fungsi, sedangkan slice bisa digunakan sebagai parameter fungsi
	iniArray := [...]string{"A", "B", "C"}
	iniSlice := []string{"A", "B", "C"}

	fmt.Println("Ini Array:", iniArray)
	fmt.Println("Ini Slice:", iniSlice)
}
