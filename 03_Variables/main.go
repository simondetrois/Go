package main

import "fmt"

/*
The var statement declares a list of vars (like in functions) with the type at least.
The var statement can be declared at package or at function level.
*/
var a, b, c int

/*
like imports, variable declaration can be put into blocks
*/
var ( // Variables declared without an explicit initial value are given their zero value.
	varA1 bool
	varA2 bool = false // zero value for bools
	varB1 int
	varB2 int = 0 // zero value for numeric types
	varC1 string
	varC2 string = "" // zero value for strings
)

func main() {
	fmt.Println(varA1, varA2, varB1, varB2, varC1, varC2)

	var d int
	/*
		If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
	*/
	var e, f, g = "e", "f", false

	fmt.Println(a, b, c, d, e, f, g)

	/*
		Inside a function, the short assignment (:=) can be used, outside a function everything must start with a keyword like var or func
	*/
	var h string = "h"
	i, j := "j", 187
	fmt.Println(h, i, j)

	/*
		These are the basic types in Go:
		- bool
		- string
		- int  int8  int16  int32  int64
		  uint uint8 uint16 uint32 uint64 uintptr
		- byte // alias for uint8
		- rune // alias for int32 (=> Unicode Code Point)
		- float32 float64
		- 	complex64 complex128

		With Type conversion ( T([var]) ) you can switch between Types
	*/

	var k, l int = 1.0, 2
	m, n := int32(k), float64(l)

	fmt.Println(m, n)

	/*
		Constants can only be declare using const identifier [type] = value
	*/
	const o int = 1

	fmt.Print(o)
}
