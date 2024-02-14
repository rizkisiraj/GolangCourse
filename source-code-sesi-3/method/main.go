package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Introduce(message string) string {
	return fmt.Sprintf("%s My name is %s and I'm %d years old", message, p.Name, p.Age)
}

func (p Person) ChangeName1() {
	p.Name = "Fiyu"
}

func (p *Person) ChangeName2() {
	p.Name = "Fiyu"
}

func main() {
	person := Person{Name: "Fitri", Age: 20}
	fmt.Println(person.Introduce("Helloooo!"))

	person.ChangeName1()
	fmt.Println("Change name with ChangeName1 method", person.Name)

	person.ChangeName2()
	fmt.Println("Change name with ChangeName2 method", person.Name)
}
