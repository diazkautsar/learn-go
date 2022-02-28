package main

import "fmt"

func main() {
	var counter int = 1

	for counter <= 10 {
		fmt.Println("Counter ", counter)
		counter++
	}

	for i := 1; i < 5; i++ {
		fmt.Println("Another case", i)
	}

	names := []string{"Diaz", "Djalal", "Kautsar"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// for range
	for index, name := range names {
		fmt.Println("index = ", index, "value = ", name)
	}

	person := make(map[string]string)
	person["Name"] = "Djalal Kurnia"
	person["Job"] = "Software Enggineer"

	for key, value := range person {
		fmt.Println(key, value)
	}
}
