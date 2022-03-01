package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "DJALAL"
	middleName = "NAMA TENGAH KURNIA"
	lastName = "KURNIA"

	return
}

func main() {
	firstName, miiddleName, lastName := getFullName()

	fmt.Println(firstName, miiddleName, lastName)
}
