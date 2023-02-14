package main

import "fmt"

// panic => digunakan untuk menghentikan program
// panic => biasanya digunakan ketika terjadi error pada saat program kita berjalan
// panic => pada saat panic function di panngil, defer function akan tetap di eksekusi

func endApp() {
	message := recover() // ini untuk recover. untuk ketika panic terjadi, dia akan tangkap data panic, dan aplikasi akan tetap berjalan
	fmt.Println("ERROR WITH MESSAGE ", message)
	fmt.Println("END APP BRO")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR BRO")
	}
	fmt.Println("APLIKASI AP")
}

func main() {
	runApp(false)
	//runApp(true)
}
