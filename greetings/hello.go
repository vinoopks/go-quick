package greetings

import "fmt"

//Hello returns the greetings for the named person

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. I am Vinoop! Nice to meet you!", name)
	return message
}
