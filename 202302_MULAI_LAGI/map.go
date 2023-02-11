package main

import "fmt"

// map kaya object di javascript

// person := map[typeDataKeynya]typeDataNilainya {}

func main() {
	person := map[string]string{
		"name":    "Djalal Kurnia",
		"address": "Lebak",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	person["job"] = "Programmer"

	fmt.Println(person)

	book := make(map[string]string)

	book["title"] = "Buah Jatuh Jauh Anaknya"
	book["authos"] = "Dadang Koneal"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}
