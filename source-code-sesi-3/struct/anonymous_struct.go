package main

import "fmt"

// change this function into main() and comment the other main() func to run this code
func anonymousStruct() {

	// anonymous struct tanpa pengisian field
	var member2 = struct {
		person   Person
		position string
	}{}
	member2.person = Person{name: "Jisoo", age: 27}
	member2.position = "Visual, vocal"

	// anonymous struct dengan pengisian field
	var member3 = struct {
		person   Person
		position string
	}{
		person:   Person{name: "Rose", age: 26},
		position: "Main vocal",
	}

	fmt.Printf("Member 2: %+v\n", member2)
	fmt.Printf("Member 3: %+v\n", member3)

}
