package main

import "fmt"

// dalam struct method juga berlaku pass by value maupun pas by reference

// berikut contoh untuk pass by value

type Man struct {
	Name string
}

func (man Man) MarriedPassByValue() {
	man.Name = "Mr. " + man.Name
}

func (man *Man) MarriedPassByReference() {
	man.Name = "Mr. " + man.Name
}

func main() {
	var laki Man = Man{
		Name: "DJALAL",
	}
	// ketika kirim ke function hanya kirim pass by value
	// tapi tidak menjadi berubah, karena yang dikirim duplikasinya
	laki.MarriedPassByValue()
	fmt.Println(laki)

	// INI CODE Dengan pointer
	var laki2 Man = Man{
		Name: "DJALAL",
	}
	laki2.MarriedPassByReference()
	fmt.Println(laki2)
}
