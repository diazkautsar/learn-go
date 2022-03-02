package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (cutomer Customer) sayHello() {
	fmt.Println(cutomer.Name)
}

func main() {
	var diaz Customer
	diaz.Name = "Djalal Kurnia"
	diaz.sayHello()
}
