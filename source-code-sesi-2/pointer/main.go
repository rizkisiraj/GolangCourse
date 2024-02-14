package main

import (
	"fmt"
	"strings"
)

func main() {

	// declaration
	var firstExample *int
	var secondExample *int
	_, _ = firstExample, secondExample

	// initialization (memory address)
	var firstNumber int = 4
	var secondNumber *int = &firstNumber
	fmt.Println()
	fmt.Println("firstNumber (value) : ", firstNumber)
	fmt.Println("firstNumber (memory address) : ", &firstNumber)
	fmt.Println("secondNumber (value) : ", *secondNumber)
	fmt.Println("secondNumber (memory address) : ", secondNumber)

	// changing value through pointer
	var firstGroup string = "BlackPink"
	var secondGroup *string = &firstGroup
	fmt.Println("\n", strings.Repeat("-", 20), "\n")
	fmt.Println("firstGroup (value) : ", firstGroup)
	fmt.Println("firstGroup (memory address) : ", &firstGroup)
	fmt.Println("secondGroup (value) : ", *secondGroup)
	fmt.Println("secondGroup (memory address) : ", secondGroup)
	fmt.Println("\n", strings.Repeat("#", 30), "\n")

	firstGroup = "NewJeans"
	*secondGroup = "MayKid"
	fmt.Println("firstGroup (value) : ", firstGroup)
	fmt.Println("firstGroup (memory address) : ", &firstGroup)
	fmt.Println("secondGroup (value) : ", *secondGroup)
	fmt.Println("secondGroup (memory address) : ", secondGroup)

	// pointer as parameter
	var age int = 10
	fmt.Println("\n", strings.Repeat("-", 20), "\n")
	fmt.Println("Before:", age)
	changeValue(&age)
	fmt.Println("After:", age)

}

func changeValue(number *int) {
	*number = 20
}
