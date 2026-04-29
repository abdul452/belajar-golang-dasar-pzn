package main

import "fmt"

func main() {
	type NoKtp string

	var noKtpSaya NoKtp = "1234567890"
	var contoh string = "1111111111"
	var contohKtp = NoKtp(contoh)

	fmt.Println(contohKtp)
	fmt.Println(noKtpSaya)

}
