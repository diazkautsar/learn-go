package main

import "fmt"

// defer => dipanggil setelah function beneran selesai
// defer => tidak peduli error atau tidak

func logging() {
	fmt.Println("LOGGING BRO")
}

func runApplication(value int) {
	defer logging() // tapi jika disimpen disini sebagai defer. dia akan kepanggil (baik error ataupun tidak)
	fmt.Println("RUN APPLICATION")
	result := 10 / value
	fmt.Println("RESULT ", result)
	//logging() // dengan kondisi runApplication kirim value 0 dan logging disimpan disini, karena di line 14 error. jadi ini tidak akan kepanggil
}

func main() {
	runApplication(5)
}
