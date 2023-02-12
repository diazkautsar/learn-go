package main

import "fmt"

// function bisa disempen sebagai value dari sebuah variabel

func getGoodBye(name string) string {
	return "Good bye " + name
}

func main() {
	sayGoodBye := getGoodBye
	fmt.Println(sayGoodBye("DJALAL"))
}
