package main

import "fmt"

// adalah interface yang bisa dipasang type data apapun atau return type data apapun

func ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "ups"
	}
}

func main() {
	paramSatu := ups(1)
	fmt.Println(paramSatu)

	paramDua := ups(2)
	fmt.Println(paramDua)
}
