package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blackList := func(name string) bool { // its call anonymos function because without name
		return name == "ANJING"
	}

	registerUser("DJALAL", blackList)
	registerUser("ANJING", blackList)
	registerUser("DJALAL", func(name string) bool {
		return name == "ANJING"
	})
	registerUser("ANJING", func(name string) bool {
		return name == "ANJING"
	})
}
