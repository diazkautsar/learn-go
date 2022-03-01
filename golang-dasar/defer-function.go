package main

import "fmt"

// defer => untuk memanggil function, setelah satu block function selesai di eksekusi
// defer => tidak peduli error atau success pada function tersebut

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging()
	fmt.Println("Run application")
}

func main() {
	runApplication()
}
