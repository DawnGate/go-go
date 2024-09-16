package main

import (
	"fmt"
	"log"

	"example.com/greetings"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println(quote.Go())

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Kraken")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
