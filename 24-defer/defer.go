package main

func logging() {
	println("logging")
}

func runApp(value int) {
	defer logging()
	// jadi jika terjadi error di bawah ini, maka fungsi logging tetap akan dijalankan
	println("run application")
	result := 10 / value
	println("result", result)
}

func main() {
	runApp(5)
}
