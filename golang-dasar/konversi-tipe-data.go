package main

import "fmt"

func main() {
	var nilai32 int32 = 100000

	var nilai64 int64 = int64(nilai32)

	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	// mengambil index data di string akan jadi byte
	// di konversi kembali dengan function string()
	var name = "Djalal"
	var e = name[0] // e itu adalah byte / uint8
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(eString)

}
