package main

import "fmt"

type personStruct struct {
	name, country string
	age           int
	isAdmin       bool
}

func removeFiveYears2(person *personStruct) {
	person.age = person.age - 5
}

func structs() {

	var marianne personStruct
	p := new(personStruct)
	p2 := personStruct{"Marianne", "Peru", 22, false}
	removeFiveYears2(&p2)
	fmt.Println("Hi", marianne, p, *p, p2, p2.name)
}

func mainStructs() {
	var marianne personStruct
	p := new(personStruct)
	p2 := personStruct{"Marianne", "Peru", 22, false}
	//removeFiveYears(&p2)
	p2.removeFiveYears()
	fmt.Println("Hi", marianne, p, *p, p2, p2.name)
	p2.speakBitch()
}
