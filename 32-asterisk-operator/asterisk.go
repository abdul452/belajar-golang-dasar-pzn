package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	var address1 Address = Address{"Bandung", "Jawa Barat", "Indonesia"}
	// bikin memory baru address1
	var address2 *Address = &address1 // pass by reference, address2 menyimpan alamat dari address1
	// var address2 adalah nama lain untuk akses ke address1
	// jadi jika ingin mengubah address1, bisa menggunakan address2 atau langsung address1
	address2.City = "Jakarta"
	fmt.Println(address1)
	fmt.Println(address2)

	// kita mau pindahkan address2 ke memory baru, jadi address2 tidak lagi menunjuk ke address1
	// address2 = &Address{"Surabaya", "Jawa Timur", "Indonesia"}
	// asterisk
	*address2 = Address{"Surabaya", "Jawa Timur", "Indonesia"}
	// sekarang address2 menunjuk ke memory yang sama dengan address1, tapi value nya sudah berubah
	fmt.Println(address1)
	fmt.Println(address2)

	// operator new
	var address3 *Address = new(Address) // new akan membuat memory baru untuk address3
	// address3 adalah pointer yang menunjuk ke memory baru yang dibuat oleh new
	// jadi address3 menyimpan alamat dari memory baru yang dibuat oleh new
	// kita bisa mengakses value dari memory baru tersebut menggunakan operator asterisk
	address3.City = "Padang" // bisa juga langsung mengubah value dari memory baru tersebut menggunakan operator dot
	fmt.Println(address3)
}
