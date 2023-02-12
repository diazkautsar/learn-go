package main

import "fmt"

func getFullNameAgain(firstName string, lastName string) string {
	var returnValue string = firstName + " " + lastName

	return returnValue
}

func main() {
	firstName := "Djalal"
	lastName := "Kurnia"

	var result string = getFullNameAgain(firstName, lastName)
	fmt.Println(result)
}
