package main

import "fmt"

func (p *personStruct) removeFiveYears() {
	p.age = p.age - 5
}

func (p *personStruct) speakBitch() {
	fmt.Println("Hey Bitch, my name is", p.name)
}

func (p *personStruct) sayHello() {
	fmt.Println("Hello Bitch")
}

type androidStruct struct {
	Person personStruct
	Model  string
}

type androidStruct2 struct {
	personStruct
	Model string
}

func mainMethods() {
	var marianne personStruct
	p := new(personStruct)
	p2 := personStruct{"Marianne", "Peru", 22, false}
	//removeFiveYears(&p2)
	p2.removeFiveYears()
	fmt.Println("Hi", marianne, p, *p, p2, p2.name)
	p2.speakBitch()

	and := new(androidStruct)
	and.Person.sayHello()

	and2 := new(androidStruct2)
	and2.personStruct.sayHello()
	and2.sayHello()
}
