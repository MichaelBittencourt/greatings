package greatings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) { // function that start in upper case are functions only called externally
	// If no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets initial values for variables used in the function.

func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greatings messages. The returned
// message is selected at random

func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
