package main

import (
	"fmt"
	"os"

	"github.com/pwd5032/go-learn/basics"
)

/*
every main.go file must have a main() function
*/

func main() {
	userName := "Patrick"
	fmt.Printf("Hello %v,\n", userName)

	// Function from package example from functs.go
	fmt.Printf("Result: %v of Type: %T\n", basics.Simple_sum(1, 2), basics.Simple_sum(1, 1))

	// Execute data type example
	basics.PrintDataTypes()

	// Execute command line arg example
	if len(os.Args) > 1 {
		basics.Get_and_print_arg()
	}

	// Execute flow example
	basics.UsePrimes(40)

	// Execute switch example
	basics.FizzBuzz(30)

	// Execute panic defer recover example
	basics.PanicIfNegative()

	// Point & Address example from functs.go
	basics.UpdateName(&userName) // Function with point arg
	fmt.Printf("Goodbye %v\n", userName)
}
