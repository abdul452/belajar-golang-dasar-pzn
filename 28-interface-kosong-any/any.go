package main

import "fmt"

func Ups() any { // ini any adalah tipe data kosong yang bisa menampung semua tipe data, sama dengan interface{}
	// bisa return apa saja, karena tipe data any bisa menampung semua tipe data
	// return "Ups" // string
	// return 123 // int
	// return []string{"Ups", "Ups", "Ups"} // slice
	// return map[string]string{ // map
	// 	"name": "Ups",
	// 	"job":  "Programmer",
	// }
	// return struct { // struct
	// 	Name string
	// 	Job  string
	// }{
	// 	Name: "Ups",
	// 	Job:  "Programmer",
	// }
	// return func() { // function
	// 	fmt.Println("Ups")
	// }
	return true // boolean
}

// bisa juga seperti ini
func Ups2() interface{} { // ini sama dengan any, karena interface{} adalah tipe data kosong yang bisa menampung semua tipe data
	// return "Ups"
	return 123
}

func main() {

	var data any = Ups()
	fmt.Println(data)
	var data2 interface{} = Ups2()
	fmt.Println(data2)

}
