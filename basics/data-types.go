package basics

import (
	"fmt"
	"strconv"
)

/*
Variables must be used if declared
	Exception might be is used in an underlying package (this)
Constants don't have to be used

Types
Basic: numbers, strings, and booleans
Aggregate: arrays and structs
Reference: pointers, slices, maps, functions, and channels
Interface: interface

Escape Characters
\n new line
\r carriage return
\t tabs
\' single quote
\" double quote
\\ backslash
*/

// Example Variable Initialization
var firstName, lastName string
var age int

var (
	dogName string = "Tonks"
	dogLegs int    = 4
)

var height float32 = 60.1
var tonksCuteFlag bool = true

// Type conversion with implicit variable type assignment
var bigHeight = float64(height)

// Constant Declaration
const RanchHouseName = "Hilltop"

func PrintDataTypes() {
	// In-function variable declaration
	myFirstName, myLastName := "Patrick", "Doyle"
	fmt.Printf("basics.Print_data_types output: %v, %v, %v\n", age, myFirstName, myLastName)

	// int to string conversion
	fmt.Println("This used to be an int: ", strconv.Itoa(dogLegs))
}
