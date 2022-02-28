package main

import "fmt"

// break => menghentikan perulangan
// continue => dilewat, masuk lagi ke perulangan

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Perulangan ke - ", i)
	}
}
