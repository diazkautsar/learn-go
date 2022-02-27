package main

import "fmt"

func main() {
	var name string

	name = "Diaz Kautsar"
	fmt.Println(name)

	name = "Djalal Kurnia"
	fmt.Println(name)

	// tidak wajib masukan type data jika langsung memasukan value
	var friendName = "Hengki"
	fmt.Println(friendName)

	// kata kunci var tidak wajib. jika menggunakan := dan langsung menginisiasikan valuenya
	// digunakan hanya untk deklarasi awal
	// jika merubah selanjutnya tidak perlu := lagi
	country := "Indonesia"
	fmt.Println(country)

	// another option. if we declared more variable
	var (
		firstName = "Djalal"
		lastName  = "Kurnia"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}

// if we declared variable and never used. code will error
