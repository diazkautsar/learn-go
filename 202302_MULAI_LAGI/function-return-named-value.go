package main

import "fmt"

func getFullNameAgain3() (firstName, lastName string) { // ini return nya bisa di deklarasikan
	firstName = "Djalal"
	lastName = "Kurnia"

	return
}

func main() {
	a, b := getFullNameAgain3()

	fmt.Println(a)
	fmt.Println(b)
}
