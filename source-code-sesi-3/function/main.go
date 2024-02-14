package main

import (
	"fmt"
	"strings"
)

func main() {
	greet("Fitri", 20)

	secondGreet("Fitri", "Bekasi")

	printMessage := thirdGreet("Hallooo", "Fitri")
	fmt.Println(printMessage)

	// result func returning multiple values
	multipleResult, dividedResult := calculate(5)
	fmt.Println("Multiple :", multipleResult)
	fmt.Println("Divided :", dividedResult)

	// result func predefined return values
	newMultipleResult, newDividedResult := newCalculate(10)
	fmt.Println("New Multiple :", newMultipleResult)
	fmt.Println("New Divided :", newDividedResult)

	// result func variadic 1
	studentLists := printStudent("Asrie", "Nadia", "Neneng", "Risya", "Daniela")
	fmt.Println(studentLists)

	// result func variadic 2
	numberLists := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := sumNumber(numberLists...)
	fmt.Println(result)

	// result func variadic 3
	printFavGroup("Fitri", "BlackPink", "ITZY", "NewJeans", "Red Velvet", "Aespa", "Girls Generation")
	// girlGroups := []string{"BlackPink", "ITZY", "NewJeans", "Red Velvet", "Aespa", "Girls Generation"}
	// printFavGroup("Fitri", girlGroups...)
}

func greet(name string, age int) {
	fmt.Println("Hello there, My name is", name, "and I'm", age, "years old")
}

// function (parameter)
func secondGreet(name, address string) {
	fmt.Println("Hello there, My name is", name)
	fmt.Println("I live in", address)
}

// function (return)
func thirdGreet(message string, names string) string {
	// result := fmt.Sprintf("%s %s", message, names)
	var result string = fmt.Sprintf("%s %s", message, names)
	return result
}

// function (returning multiple values)
func calculate(number float64) (float64, float64) {
	multiple := number * number
	divided := number / number

	return multiple, divided
}

// function (predefined return values)
func newCalculate(number float64) (multiple float64, divided float64) {
	multiple = number * number
	divided = number / number

	return
}

// function (variadic function #1)
func printStudent(names ...string) []string {
	var result []string

	for _, value := range names {
		result = append(result, value)
	}

	return result
}

// function (variadic function #2)
func sumNumber(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

// function (variadic function #2)
func printFavGroup(name string, favGroups ...string) {
	mergeFavGroups := strings.Join(favGroups, ", ")

	fmt.Println("Hello there, I'm", name)
	fmt.Println("My favorite groups are", mergeFavGroups)
}
