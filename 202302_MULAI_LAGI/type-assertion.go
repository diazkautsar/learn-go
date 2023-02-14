package main

import "fmt"

// kemampuan merubah tipe data menjadi type data yang diinginkan
// fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong
// konversi harus dilakukan berdasarkan type data returnnya

func random() interface{} {
	return true
}

func main() {
	var result interface{} = random()
	//var resultString string = result.(string)
	//fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("INI STRING " + value)
	case int:
		fmt.Println("INI INTEGER ", value)
	default:
		fmt.Println("TIDAK TAUUU", value)
	}
}
