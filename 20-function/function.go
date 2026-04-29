package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

// func parameter
func sayHelloWithName(name string) {
	fmt.Println("Hello", name)
}

// func return value
func getHelloWithName(name string) string {
	return "Hello " + name
}

// func return int
func sum(a int, b int) int {
	return a + b
}

// func return multiple value
func calculate(a int, b int) (int, int) {
	return a + b, a * b
}

// func named return value
func calculateWithNamedReturn(a int, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	return sum, product
}

// func return slice
func getEvenNumbers(numbers []int) []int {
	var evenNumbers []int
	for _, number := range numbers {
		if number%2 == 0 {
			evenNumbers = append(evenNumbers, number)
		}
	}
	return evenNumbers
}

// func return interface{}
func getData(data interface{}) interface{} {
	return data
}

// func variadic parameter, keuntungan func variadic parameter adalah kita bisa mengirimkan banyak parameter dengan tipe data yang sama tanpa harus membuat slice terlebih dahulu
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// func value
func getGoodBye(name string) string {
	return "Good Bye " + name
}

// func as parameter
// untuk func as parameter, parameter bisa berupa alias
type Filter func(string) string

// contoh func as parameter dengan alias
func sayHelloWithFilterAlias(name string, filter Filter) {
	name = filter(name)
	fmt.Println("Hello", name)
}

// contoh func as parameter tanpa alias
func sayHelloWithFilter(name string, filter func(string) string) {
	name = filter(name)
	fmt.Println("Hello", name)
}

// func return func
func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	}
	return name
}

func main() {
	sayHello()
	sayHelloWithName("Alice")
	sayHelloWithName("Bob")
	fmt.Println(getHelloWithName("Charlie"))
	fmt.Println(sum(5, 5))
	result1, result2 := calculate(5, 5)
	fmt.Println(result1, result2)
	sum, product := calculateWithNamedReturn(5, 5)
	fmt.Println(sum, product)
	fmt.Println(getEvenNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(getData("Hello, World!"))
	fmt.Println(sumAll(1, 2, 3, 4, 5))
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(sumAll(numbers...)) // kita bisa mengirimkan slice dengan menggunakan ... setelah nama slice

	contoh := getGoodBye("Diana")
	misal := getGoodBye("Eve")
	fmt.Println(contoh)
	fmt.Println(misal)

	sayHelloWithFilter("Anjing", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Kucing", filter)
	sayHelloWithFilterAlias("Anjing", spamFilter)
}
