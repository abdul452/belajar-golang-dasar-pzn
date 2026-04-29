package main

// ini komentar

/*	ini komentar juga
	tapi lebih dari satu baris
	banyak baris
	atau paragraf
*/
import "fmt"

func endApp() {
	fmt.Println("Aplikasi Selesai")
	message := recover()
	if message != nil {
		fmt.Println("Error Terjadi:", message)
	}
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(true)
	fmt.Println("Aplikasi Tetap Berjalan")
}
