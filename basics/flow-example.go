package basics

import "fmt"

func findPrimes(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	if number > 1 {
		return true
	} else {
		return false
	}
}

func UsePrimes(end int) {
	fmt.Printf("Prime numbers less than %v: ", end)

	for number := 1; number < end; number++ {
		if findPrimes(number) {
			fmt.Printf("%v ", number)
		}
	}

	fmt.Printf("\n")
}
