package main

import "fmt"

// pointer juga berlaku ke dalam sebuah function

type AddressAgain struct {
	City, Province, Country string
}

// fucntion dengan kondisi tidak pakai pointer (pass by value)
func ChangeAddressToIndonesiaNotPointer(address AddressAgain) {
	address.Country = "INDONESIA"
}

func ChangeAddressToIndonesiaPointer(address *AddressAgain) {
	address.Country = "INDONESIA"
}

func main() {
	// pass by value
	var address1 AddressAgain = AddressAgain{"LEBAK", "BANTEN", ""}
	ChangeAddressToIndonesiaNotPointer(address1)
	fmt.Println(address1)

	// pass by reference
	var address2 AddressAgain = AddressAgain{"LEBAK", "BANTEN", ""}
	ChangeAddressToIndonesiaPointer(&address2)
	fmt.Println(address2)
}
