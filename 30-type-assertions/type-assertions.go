package main

import "fmt"

func random() any {
	// return "OK"
	// return 42
	return true
}

func main() {
	result := random()

	// Type assertion
	value, ok := result.(string)
	if ok {
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Type assertion failed")
	}

	// cek type dengan switch
	switch v := result.(type) {
	case string:
		fmt.Println("Value is a string:", v)
	case int:
		fmt.Println("Value is an integer:", v)
	default:
		fmt.Println("Unknown type:", v)
	}
}
