package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // keluar dari loop ketika i sama dengan 5
		}
		fmt.Println("i:", i)
	}

	// continue statement
	for j := 0; j < 10; j++ {
		if j%2 == 0 {
			continue // lewati iterasi ini jika j adalah angka genap
		}
		fmt.Println("j:", j)
	}

	// break dengan label
OuterLoop:
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if x == 1 && y == 1 {
				break OuterLoop // keluar dari loop luar ketika x dan y sama dengan 1
			}
			fmt.Printf("x: %d, y: %d\n", x, y)
		}
	}
}
