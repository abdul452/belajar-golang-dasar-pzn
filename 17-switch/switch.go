package main

func main() {
	// switch statement
	// switch expression {
	// case value1:
	//     // code
	// case value2:
	//     // code
	// default:
	//     // code
	// }
	name := "Abdul"

	switch name {
	case "Abdul":
		println("Nama adalah Abdul")
	case "Budi":
		println("Nama adalah Budi")
	default:
		println("Nama bukan Abdul atau Budi")
	}

	switch length := len(name); {
	case length > 5:
		println("Nama terlalu panjang")
	default:
		println("Nama cukup pendek")
	}

	// switch short statement
	switch length := len(name); length > 5 {
	case true:
		println("Nama terlalu panjang")
	case false:
		println("Nama cukup pendek")
	}

	// switch tanpa expression
	switch {
	case len(name) > 5:
		println("Nama terlalu panjang")
	default:
		println("Nama cukup pendek")
	}
}
