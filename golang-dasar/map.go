package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Djalal Kurnia",
		"address": "Bogor",
	}

	person["Job"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book = make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Djalal"
	book["wrong"] = "Salah"

	fmt.Println(book)

	delete(book, "wrong")

	fmt.Println(book)
}
