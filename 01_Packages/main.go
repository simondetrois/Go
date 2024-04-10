/*
All Go programs are made up on packages.
The main package is the starting point for a program.
Other Packages can be imported like:

	import "PACKAGE_NAME"

or when having multiple (factored) imports

import (

	"PACKAGE_1"
	"PACKAGE_2"

)
*/
package main

import (
	"fmt"
	"packages/exporter"
)

func main() {
	fmt.Println("Hello world")

	/*
		Only capitalized names can be exported from other packages.
	*/
	exporter.Public()
}
