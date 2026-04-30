package main

import "fmt"

type Man struct {
	Name string
}

// func (receiver) methodName() { ... }
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	man := Man{Name: "John"}
	man.Married()
	fmt.Println(man.Name)
}
