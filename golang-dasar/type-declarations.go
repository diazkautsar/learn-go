package main

import "fmt"

// membuat alias untuk type data yang sudah ada
// NoKTP jadi type data string, dan bisa digunakan sebagai type data baru

func main() {
	type NoKTP string
	type Married bool

	var noKTP NoKTP = "KMZWAY87AA"
	var marriedStatus Married = true

	fmt.Println(noKTP)
	fmt.Println(marriedStatus)
}
