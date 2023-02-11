package main

import "fmt"

// kumpulan data dengan tipe sama
// perlu menentukan jumlah data arraynya
// daya tampung array tidak bisa bertambah setelah array dibuat

func main() {
	// first way
	var names [3]string
	names[0] = "Djalal"
	names[1] = "Kurnia"
	names[2] = "Yusuf"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// second way. langsung bikin array dengan isi datanya
	var values = [3]int{
		1, 2, 3,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values)) // check panjang array
}
