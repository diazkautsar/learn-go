package main

import "fmt"

func getFullName(firstName string, lastName string) (string, string) {
	return firstName, lastName
}

func main() {
	firstName, lastName := getFullName("Djalal", "Kurnia")

	fmt.Println(firstName, lastName)

	// to ignore multiple return using _
}
