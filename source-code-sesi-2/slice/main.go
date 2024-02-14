package main

import (
	"fmt"
	"strings"
)

func main() {

	// slice (make function)
	var animals = make([]string, 3)
	animals[0] = "cat"
	animals[1] = "bird"
	animals[2] = "rabbit"

	// slice (append function)
	animals = append(animals, "kucing", "burung", "kelinci")
	fmt.Printf("%#v\n", animals)

	// slice (append function with ellipsis)
	animals1 := []string{"cat", "bird", "rabbit"}
	animals2 := []string{"horse", "duck", "butterfly"}

	animals1 = append(animals1, animals2...)
	fmt.Printf("%#v\n", animals1)

	// slice (copy function)
	copyFunc := copy(animals1, animals2)
	fmt.Println("Animals 1 =>", animals1)
	fmt.Println("Animals 2 =>", animals2)
	fmt.Println("Copied elements =>", copyFunc)

	// slice (slicing)
	fruits1 := []string{"apple", "mango", "orange", "banana", "grape"}

	var fruits2 = fruits1[1:4]
	fmt.Printf("%#v\n", fruits2)

	var fruits3 = fruits1[0:]
	fmt.Printf("%#v\n", fruits3)

	var fruits4 = fruits1[:3]
	fmt.Printf("%#v\n", fruits4)

	var fruits5 = fruits1[:] // equal to fruits1[:len(fruits1)]
	fmt.Printf("%#v\n", fruits5)

	// slice (combining slicing & append)
	colors := []string{"red", "green", "blue", "purple", "white"}
	colors = append(colors[:3], "pink")
	fmt.Printf("%#v\n", colors)

	// slice (backing array)
	var lang1 = []string{"python", "ruby", "java", "go", "javascript"}
	var lang2 = lang1[2:4]
	lang2[0] = "rust"
	fmt.Println("Lang 1 =>", lang1)
	fmt.Println("Lang 2 =>", lang2)

	// slice (cap function)
	var newFruits1 = []string{"apple", "mango", "durian", "banana"}
	fmt.Println("Fruits1 cap:", cap(newFruits1))
	fmt.Println("Fruits1 len:", len(newFruits1))
	fmt.Println(strings.Repeat("-", 20))

	var newFruits2 = newFruits1[0:3]
	newFruits2[0] = "Jeruk"
	newFruits1[2] = "Jasuke"
	fmt.Println("Fruits1: (value)", newFruits1)
	fmt.Println("Fruits2: (value)", newFruits2)
	fmt.Println("Fruits2 cap:", cap(newFruits2))
	fmt.Println("Fruits2 len:", len(newFruits2))
	fmt.Println(strings.Repeat("-", 20))

	var newFruits3 = newFruits1[1:]
	fmt.Println("Fruits3 cap:", cap(newFruits3))
	fmt.Println("Fruits3 len:", len(newFruits3))
	fmt.Println(strings.Repeat("-", 20))

	// slice (create a new backing array)
	cars := []string{"Ford", "Honda", "Mazda", "Audi", "VW"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)
	cars[0] = "Nissan"
	fmt.Println("cars:", cars)
	fmt.Println("newCars", newCars)
}
