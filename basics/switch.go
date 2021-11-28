package basics

import "fmt"

// Example function for switch statements
func FizzBuzz(max int) {
	fmt.Printf("Switch: \n")
	for i := 1; i <= max; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Printf("FizzBuzz ")
		case i%3 == 0:
			fmt.Printf("Fizz ")
		case i%5 == 0:
			fmt.Printf("Buzz ")
		default:
			fmt.Printf("%v ", i)
		}
	}
	fmt.Println()
}
