package main

import "exported-unexported/helpers"

func main() {
	helpers.Greet()

	person := helpers.Person{}
	person.Invokegreet()
}
