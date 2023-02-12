package main

import "fmt"

// variabel argument

func sumAll(numbers ...int) int { // harus diujung kanan untuk variabel argumentnya
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

// misal kita kirim lebih daari satu data, maka akan di tambung menjadi sebuah variabel numbers. dan menjadi slice

func main() {
	total := sumAll(10, 20, 30)
	fmt.Println(total)

	// jika kita punya slice. dan akan dikirim ke variadic function, maka tambahkan titik tiga di variable slicenya
	numbers := []int{10, 50, 100}
	totalInSlice := sumAll(numbers...)
	fmt.Println(totalInSlice)
}
