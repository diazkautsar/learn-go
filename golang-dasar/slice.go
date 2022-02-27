package main

import "fmt"

func main() {
	// slice potongan dari array.
	// mirip array, untuk slice ukuran bisa berubah

	// slice memiliki 3 data. pointer, length, capacity
	// pointer => penunjuk data perama di array pada slice
	// length => panjang slice
	// capacity => total data dalam length tersebut

	// merubah array mengubah slicenya juga
	// merubah slice akan mengubah arraynya juga

	// create slice
	var months = [...]string{ // ... untuk memasukan kita belum tau jumlah datanya berapa
		"January",
		"February",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"July",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "Diubah"
	fmt.Println(slice1)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Djalal")
	fmt.Println(slice3)

	// create slice dari awal
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Djalal"
	newSlice[1] = "Kurnia"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [2]int{1, 2}
	iniSlice := []int{1, 2}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
