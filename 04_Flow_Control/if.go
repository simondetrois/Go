package main

import "fmt"

func If() {
	/*
		Like with for-Loops, if conditions aren't required to be surrounded with (), but {} is required
	*/
	if 1 < 2 {
		fmt.Print("Wow")
	}

	/*
		Like for loops, if statements can start with a statement before executing the condition.
		Vars declared in that statement are only in scope until the end of if
	*/
	if test := 1; test < 4 {
		fmt.Print(test)
	}

}
