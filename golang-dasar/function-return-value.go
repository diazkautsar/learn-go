package main

import "fmt"

func completeName(firstName string, lastName string) string {
	return firstName + " " + lastName
}

func main() {
	var firstName string = "Djalal"
	var lastName string = "Kurnia"

	var completeName string = completeName(firstName, lastName)

	fmt.Println(completeName)
}
