package main

import "fmt"

// kirim function sebagai paramter

func functionAsParameterFirstExample(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello " + nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "...."
	} else {
		return name
	}
}

func main() {
	//spamFilterFunction := spamFilter

	//result := functionAsParameterFirstExample("DJALAL", spamFilterFunction("DK"))
	//fmt.Println(result)
}
