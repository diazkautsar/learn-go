package main

import "fmt"

func getFullNameAgain2() (string, string) { // bisa lebih dari satu returnnya
	return "Djalal", "Kurnia"
}

func main() {
	firstName, lastName := getFullNameAgain2()

	fmt.Println(firstName, lastName)

	// Jika tidak peduli return kedua
	firstName2, _ := getFullNameAgain2()
	fmt.Println(firstName2)
}
