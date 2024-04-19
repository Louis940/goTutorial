package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// set properties of the predefined logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// a slice of names
	names := []string{"Louis", "Abi", "Ciaran", "Lola", "Phoebe"}

	// Get a greeting message for each name
	messages, err := greetings.Hellos(names)
	// If an error was returned , print it to the console and
	// exit the program
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned, print the returned message to the console.
	fmt.Println(messages)
}
