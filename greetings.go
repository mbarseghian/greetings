package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.
	// message := fmt.Sprintf("Hi, %v. Welcome!", name)

	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}

/*
In this code, you:

- Add a randomFormat function that returns a randomly selected format for a greeting message. Note that randomFormat starts with a lowercase letter, making it accessible
only to code in its own package (in other words, it's not exported).
- In randomFormat, declare a formats slice with three message formats. When declaring a slice, you omit its size in the brackets, like this: []string.
This tells Go that the size of the array underlying the slice can be dynamically changed.
- Use the math/rand package to generate a random number for selecting an item from the slice.
- Add an init function to seed the rand package with the current time. Go executes init functions automatically at program startup, after global variables have been initialized.
For more about init functions, see Effective Go.
- In Hello, call the randomFormat function to get a format for the message you'll return, then use the format and name value together to create the message.
- Return the message (or an error) as you did before.
*/
