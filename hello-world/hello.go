package main

import (
	"fmt"

	"rsc.io/quote"
	//this is for learning how to call an external module. ignore
	"github.com/vinoopks/go-quick/greetings"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	//from greetings module
	fmt.Println(greetings.Hello("Peter"))
}
