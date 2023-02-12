package main

import "fmt"

// loop in goolang, baru satu for loops

func main() {
	// first way
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	// for dengan statement
	var sliceExample = []string{"Djalal", "Kurnia"}
	for i := 0; i < len(sliceExample); i++ {
		fmt.Println(sliceExample[i])
	}

	// dengan cara lain
	// for keyNya, valueNamnya := ranges namaVariableMapSliceArray
	// kalo pak eunderscore berarti engga di pake

	for _, exp := range sliceExample {
		fmt.Println(exp)
	}
}
