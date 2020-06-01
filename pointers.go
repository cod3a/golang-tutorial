package main

import "fmt"

func add10(nPtr *int) {
	*nPtr = *nPtr + 10
}

func pointers() {
	// pointers
	x := 51
	xPtr := new(int)
	add10(&x)
	add10(xPtr)
	fmt.Println(x, *xPtr)
}
