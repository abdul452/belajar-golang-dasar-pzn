package main

import "fmt"

func main() {
	counter := 0

	increment := func() {
		fmt.Println("Counter", counter)
		counter++
	}
	increment()
	increment()

	fmt.Println("Counter Final", counter)
}
