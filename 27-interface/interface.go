package main

import "fmt"

// interface adalah sebuah kontrak yang harus dipenuhi oleh struct yang mengimplementasikannya
type HasName interface {
	GetName() string // ini adalah method signature, tidak memiliki implementasi
}

// SayHello adalah function yang menerima parameter value yang bertipe HasName, sehingga value harus mengimplementasikan interface HasName
func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

// Person struct memiliki field Name
type Person struct {
	Name string
}

// Person mengimplementasikan HasName karena memiliki method GetName yang sesuai dengan signature di interface HasName
func (person Person) GetName() string {
	return person.Name
}

func main() {
	person := Person{Name: "Eko"}
	SayHello(person)
}
