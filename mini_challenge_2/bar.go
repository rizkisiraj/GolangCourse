package main

import "fmt"

func main() {
	var stringnya string = "selamat datang"
	var person1 map[string]int
	person1 = map[string]int{}

	for key := range stringnya {
		fmt.Printf("%c\n", stringnya[key])
		value, exist := person1[string(stringnya[key])]
		if exist {
			person1[string(stringnya[key])] = value + 1
		} else {
			person1[string(stringnya[key])] = 1
		}
	}
	fmt.Println(person1)
}
