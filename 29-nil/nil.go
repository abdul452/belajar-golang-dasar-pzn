package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		// return nil hanya bisa jika tipe data yang dikembalikan adalah pointer, map, slice, channel, interface, atau function
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("Abdul")
	if data == nil {
		println("data map masih kosong")
	} else {
		fmt.Println(data["name"])
	}
}
