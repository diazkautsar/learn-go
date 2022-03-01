package main

import "fmt"

// kemampuan untuk mengambil data dari function

func main() {
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
}
