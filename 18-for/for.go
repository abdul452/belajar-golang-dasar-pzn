package main

func main() {
	counter := 1

	// for loop
	for counter <= 5 {
		println("Counter:", counter)
		counter++
	}

	// for loop dengan short statement
	for i := 1; i <= 5; i++ {
		println("i:", i)
	}

	// for loop dengan range
	names := []string{"Alice", "Bob", "Charlie"}
	for index, name := range names {
		println("Index:", index, "Name:", name)
	}

	// for loop dengan range untuk map
	person := map[string]string{
		"name":    "Alice",
		"country": "USA",
	}
	for key, value := range person {
		println("Key:", key, "Value:", value)
	}

	// for loop tanpa kondisi (infinite loop)
	// for {
	//     // code
	// }
}
