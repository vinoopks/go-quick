package greetings

import (
	"errors"
	"fmt"
)

//Hello returns the greetings for the named person

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. I am Vinoop! Nice to meet you!", name)
	return message, nil
}
