package main

import "fmt"

// panic => for error handling to stop app if error
// in panic but defer tetap kepanggil

// recover to get data error in panic
// recover set in defer function

func endApp() {
	message := recover()

	if message != nil {
		fmt.Println("Data panic: ", message)
	}

	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(true)
	fmt.Println("Djalal")
}
