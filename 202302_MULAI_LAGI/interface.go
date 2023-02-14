package main

import "fmt"

// interface berisi dengan beberapa method

type HasName interface {
	GetName() string
}

func sayHelloAgain(hasName HasName) { // bida dipakai sebuah parameter juga untuk interface
	fmt.Println("HELLO ", hasName.GetName())
}

type Person struct {
	Name string
}

// karena person punya function GetName, maka dia sudah otomatis jadi punya interface GetName
func (person Person) GetName() string {
	return person.Name
}

func main() {
	var djalal Person
	djalal.Name = "DJALAL"

	sayHelloAgain(djalal)
}
