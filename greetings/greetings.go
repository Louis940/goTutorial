package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting
func Hello(name string) (string, error) {
	//if no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	// returns a greeting that embeds the name in the message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
