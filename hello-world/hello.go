package main

import (
	"fmt" //This is a package in the standard Go libarary

	"rsc.io/quote" //added as per tutorial in golang.org.
	//this is for learning how to call an external module. ignore
	"github.com/vinoopks/go-quick/greetings"
)

//This is the entry point for the executable program
func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	//from greetings module
	greeting, error := greetings.Hello("Peter")
	if error != nil {
		fmt.Println(greeting)

	}
}
func operateTwoVariables() {
	var x int //This is how we declare variables in Go.
	x = 3     //Variable assignment
	//short hand declaration and initialization of the varible.
	y := 4
	sum, prod := learnMultipleReturn(x, y) //function can return multiple values
	fmt.Println("sum: ", sum, "product: ", prod)

}

func learnMultipleReturn(x, y int) (sum, prod int) { //x and y are the varibles  -- sum and prod are the return items
	return x + y, x * y //return two values
}
