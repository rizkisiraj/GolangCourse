package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Member struct {
	position string
	person   Person
}

// change this function into main() and comment the other main() func to run this code
func embededStruct() {
	var member1 = Member{}

	member1.person.name = "Jennie"
	member1.person.age = 27
	member1.position = "Rapper"
	fmt.Printf("%+v", member1)
}
