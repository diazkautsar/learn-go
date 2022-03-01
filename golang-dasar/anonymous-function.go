package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked ", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}

func main() {
	blacklist := func(s string) bool {
		return s == "Admin"
	}

	registerUser("Admin", blacklist)
}
