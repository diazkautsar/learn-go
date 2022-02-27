package main

import "fmt"

func main() {
	// tipe data sama
	// menentukan jumlah data yang bisa ditampung oleh si array
	// daya tampung array tidak bisa berubah setelah di-insiasi

	var names [4]string

	names[0] = "Djalal"
	names[1] = "Kurnia"
	names[2] = "Diaz"
	names[3] = "Kautsar"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names[3])

	var values = [3]int{
		10, 9, 8,
	}

	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))
}
