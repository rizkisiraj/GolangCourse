package main

import "fmt"

// change this function into main() and comment the other main() func to run this code
func sliceToStruct() {

	// slice to struct
	var people = []Person{
		{name: "Jisoo", age: 27},
		{name: "Jennie", age: 27},
		{name: "Rose", age: 26},
		{name: "Lisa", age: 26},
	}

	for _, value := range people {
		fmt.Printf("%+v\n", value)
	}

	// slice to anonymous struct
	var member = []struct {
		person   Person
		position string
	}{
		{person: Person{name: "Jisoo", age: 27}, position: "Visual, vocal"},
		{person: Person{name: "Jennie", age: 27}, position: "Rapper, vocal"},
		{person: Person{name: "Rose", age: 26}, position: "Main vocal"},
		{person: Person{name: "Lisa", age: 26}, position: "Rapper, main dancer"},
	}

	for _, value := range member {
		fmt.Printf("%+v\n", value)
	}

}
