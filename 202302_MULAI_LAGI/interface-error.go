package main

// CONTOH MEMBUAT ERROR INTERFACE

import (
	"errors"
	"fmt"
)

func Pembagian(value int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("PEMBAGI DENGAN NOL TIDAK SAH")
	} else {
		return value / pembagi, nil
	}
}

func main() {
	hasil, err := Pembagian(1, 0)
	if err == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
