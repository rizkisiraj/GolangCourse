package main

import (
	"fmt"
	"strings"
)

func main() {

	// first way of initialize map
	var person1 map[string]string // declaration
	person1 = map[string]string{} // initialization

	person1["name"] = "Jennie"
	person1["age"] = "27"
	person1["address"] = "Seoul, Korea"

	fmt.Println("name:", person1["name"])
	fmt.Println("age:", person1["age"])
	fmt.Println("address:", person1["address"])

	// second way of initialize map
	var person2 = map[string]string{
		"name":    "Rose",
		"age":     "26",
		"address": "Seoul, Korea",
	}
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println("name:", person2["name"])
	fmt.Println("age:", person2["age"])
	fmt.Println("address:", person2["address"])

	// looping with map
	fmt.Println(strings.Repeat("-", 20))
	for key, value := range person1 {
		fmt.Println(key, ":", value)
	}

	// deleting value
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println("Before deleting:", person1)
	delete(person1, "age")
	fmt.Println("After deleting:", person1)

	// detecting value
	var person3 = map[string]string{
		"name":    "Jisoo",
		"age":     "27",
		"address": "Seoul, Korea",
	}
	value, exist := person3["age"]

	fmt.Println(strings.Repeat("-", 20))
	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value doesn't exist")
	}

	delete(person3, "age")
	updatedValue, updatedExist := person3["age"]

	if updatedExist {
		fmt.Println(updatedValue)
	} else {
		fmt.Println("Value has been deleted")
	}

	// combining slice with map
	var people = []map[string]string{
		{"name": "Jisoo", "age": "27"},
		{"name": "Jennie", "age": "27"},
		{"name": "Rose", "age": "26"},
		{"name": "Lisa", "age": "26"},
	}

	fmt.Println(strings.Repeat("-", 20))
	for key, person := range people {
		fmt.Printf("Index: %d, name: %s, age: %s\n", key, person["name"], person["age"])
	}
}
