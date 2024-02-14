package main

import (
	"fmt"
	"strings"
)

type Employee struct {
	name     string
	age      int
	position string
}

func main() {
	// giving value to struct
	var employee Employee
	employee.name = "Fitri"
	employee.age = 20
	employee.position = "Backend"

	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.position)

	// initializing struct
	var employee1 = Employee{}
	employee1.name = "Ayu"
	employee1.age = 20
	employee1.position = "Fullstack"

	var employee2 = Employee{name: "Anggraini", age: 21, position: "Frontend"}
	fmt.Printf("Employee 1: %+v \n", employee1)
	fmt.Printf("Employee 2: %+v \n", employee2)

	// pointer to struct
	var employee3 *Employee = &employee2
	fmt.Printf("Employee 2 name: %+v \n", employee2.name)
	fmt.Printf("Employee 3 name: %+v \n", employee3.name)

	fmt.Println(strings.Repeat("-", 20))

	employee3.name = "Hani"
	fmt.Printf("Employee 2 name: %+v \n", employee2.name)
	fmt.Printf("Employee 3 name: %+v \n", employee3.name)
}
