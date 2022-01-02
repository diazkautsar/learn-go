package main

import (
	"fmt"

	"rsc.io/quote"

	"example.com/greetings"

	"log"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println(quote.Go())

	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"Djalal", "Dudung", "Dalmajid"} // slice name

	// message, err := greetings.Hello("") // dont have value
	// message, err := greetings.Hello("Djalal") // have value
	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
