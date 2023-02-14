package main

import "fmt"

// inisiasi untuk beberapa type data, jika awalnya kosong, bisa dibuat nil
// interface, function, map, slice, pointer, channel

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	var person map[string]string = NewMap("DJALAL")
	if person == nil {
		fmt.Println("DATA KOSONG")
	} else {
		fmt.Println(person)
	}
}
