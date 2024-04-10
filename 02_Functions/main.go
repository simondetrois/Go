package main

import "fmt"

func main() {
	one()
	two(1, 2)
	three(1, 2, "three")
	number := four(1, 2)
	anotherNumber, andAString := five(number, number)
	fmt.Printf("Another number has the value of %d, and the string is %s", anotherNumber, andAString)
	six()
}

/*
A functions is declared using the func keyword
*/

func one() {}

/*
They can receive params like func name (name type)
*/

func two(numberOne int, numberTwo int) {}

/*
Types of multiple params with the same type can be inferred by just specifying the type of the last param
*/

func three(numberOne, numberTwo int, andAString string) {}

/*
Return values are specified after the params parentheses
*/

func four(numberOne int, numberTwo int) int {
	return numberOne + numberTwo
}

/*
A function can return multiple return values
*/

func five(numberOne int, numberTwo int) (int, string) {
	return numberOne + numberTwo, "added"
}

/*
If the return values are named, they treated as vars at the top of the function.
A return statement without args returns the named return values (naked return)
*/

func six() (returnMeOne string, returnMeTwo, returnMeThree int) {
	returnMeOne = "one"
	returnMeTwo = 2
	returnMeThree = 3
	return
}
