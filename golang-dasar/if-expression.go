package main

import "fmt"

func main() {
	name := "Kurnia"

	if name == "Djalal" {
		fmt.Println("Hello Djalal")
	} else if name == "Kurnia" {
		fmt.Println("Hello Kurnia")
	} else {
		fmt.Println("Hallo untuk tidak keduanya")
	}

	// if we short statement
	if length := len(name); length > 3 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sesuai")
	}
}
