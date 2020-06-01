package main

import "fmt"

func deferAndPanic() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC Now")
}
