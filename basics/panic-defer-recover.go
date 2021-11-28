package basics

import "fmt"

func tempFunc() {
	fmt.Println("Test - please replace")
}

func PanicIfNegative() {
	for val := 0; ; {
		fmt.Print("Enter number or negative number to exit: ")
		fmt.Scanf("%d", &val)

		if val < 0 {
			panic("Uh oh - that's a negative ghost rider")
		} else if val == 0 {
			fmt.Println("0 is neither a negative or positive number.")
		} else {
			fmt.Println("You entered:", val)
		}
	}
}
