package main

import "fmt"

// kirim function sebagai paramter
// filter func(string) string (ini sebagai function as parameter)

func functionAsParameterFirstExample(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println(nameFiltered)
}

// ini parameter function
type Filter func(string) string

func functionWithTypeDeclarations(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println(nameFiltered)
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

	functionAsParameterFirstExample("Anjing", spamFilter)
	functionWithTypeDeclarations("DJALAL", spamFilter)
}
