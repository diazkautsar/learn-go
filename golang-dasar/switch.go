package main

import "fmt"

func main() {
	name := "Bambang"

	switch name {
	case "Djalal":
		fmt.Println("Hallo Djalal")
	default:
		fmt.Println("Nama tidak terdaftar")
	}

	switch length := len(name); length > 3 {
	case true:
		fmt.Println("Ini adalah case true")
	case false:
		fmt.Println("Ini adalah case false")
	}

	var age int = 50
	switch {
	case age > 50:
		fmt.Println("Tua juga anda")
	case age < 50:
		fmt.Println("Hmmm masih muda")
	default:
		fmt.Println("hmm bingung juga umurnya pas")
	}
}
