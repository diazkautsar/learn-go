package main

import "fmt"

func recursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * recursive(value-1)
	}
}

func main() {
	fmt.Println(recursive(5))
}
