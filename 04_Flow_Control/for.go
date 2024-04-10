package main

import "fmt"

func ForControl() {
	/*
		The loop construct is the only looping construct in go
		for init; condition; post
	*/
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

	/*
		init and post are condition, therefore you can just leave the semicolons
	*/
	sum := 0
	for sum < 100 {
		sum += 1
	}

	/*
		endless loop
	*/
	for {
	}
}
