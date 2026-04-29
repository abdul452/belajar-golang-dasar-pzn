package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32) // convert nilai32 ke int64
	var nilai16 int16 = int16(nilai32) // convert nilai32 ke int16, akan terjadi overflow karena nilai32 lebih besar dari batas maksimal int16 yaitu 32767

	fmt.Println("Nilai 32 =", nilai32)
	fmt.Println("Nilai 64 =", nilai64)
	fmt.Println("Nilai 16 =", nilai16)

	var name = "Abdul salim"
	var a = name[0]         // nilai name[0] menghasilkan byte, karena string di Go disimpan sebagai byte array
	var aString = string(a) // convert byte ke string, karena name[0] menghasilkan byte, maka kita perlu mengkonversinya ke string agar bisa dibaca sebagai karakter

	fmt.Println("name =", name)
	fmt.Println("a =", a)
	fmt.Println("aString =", aString)

}
