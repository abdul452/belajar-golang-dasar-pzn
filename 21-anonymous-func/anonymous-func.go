package main

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		println("You are blocked")
	} else {
		println("Welcome", name)
	}
}

func main() {
	// menggunakan anonymous function dengan assign ke variable
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("anjing", blacklist)

	// langsung menggunakan anonymous function tanpa assign ke variable
	registerUser("admin", func(name string) bool {
		return name == "anjing"
	})
}
