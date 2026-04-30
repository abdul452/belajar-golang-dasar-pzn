package helper

// version juga private
var version = "1.0.0"

// Application ini public karena huruf pertama besar, jadi bisa dipanggil dari luar package
var Application = "golang"

// ini func private karna huruf pertama tidak besar, jadi tidak bisa dipanggil dari luar package
func sayGoodbye(name string) string {
	return "Goodbye " + name
}

// ini func public karna huruf pertama besar, jadi bisa dipanggil dari luar package
func SayHello(name string) string {
	return "Hello " + name
}
