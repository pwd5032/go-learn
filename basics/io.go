package basics

import (
	"fmt"
	"os"
	"strconv"
)

func Get_and_print_arg() {
	// Grab arguments when executing main

	number1, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("basics.Get_and_print_arg output: Input Arg: %v with type: %T\n", number1, number1)
}
