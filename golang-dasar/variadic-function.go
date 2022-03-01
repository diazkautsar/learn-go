package main

import "fmt"

// variadic function disimpan di belakang
// untuk jadi slice atau array

func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {
	total := sumAll(10, 100, 200, 300, 400)

	fmt.Println(total)

	slice := []int{10, 20}
	totalSlice := sumAll(slice...)

	fmt.Println(totalSlice)
}
