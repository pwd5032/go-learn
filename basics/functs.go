package basics

/*
Function Structure
func name(parameters) (results) {
	body-content
}

If you want something to be private, start its name with a lowercase letter
If you want something to be public, start its name with an upercase letter

Go is a 'pass by value' programming language
	i.e. a copy is created when passing a variable to a function instead of passing the variable itself
The value itself may be passed by using the "&" operator to get the address of the object
	the "*" operator then dereferences the address to get access to the object

when calling a function that returns multiple values, unwanted return values can be avoided using
an underscore on variable assignment
e.g. int1, _ := some_function_w_2_returns()
*/

// Version of the functions
var Version = "1.0"

func Simple_sum(int1 int, int2 int) (result int) {
	return int1 + int2
}

func UpdateName(name *string) {
	*name = "Pat"
}
