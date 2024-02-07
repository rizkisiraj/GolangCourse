package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Print("Enter your number: ")
	_, err := fmt.Scan(&n)

	if err != nil {
		fmt.Print(err.Error())
		return
	}

	for i := 1; i <= n; i++ {
		sqrtNumber := math.Sqrt(float64(i))
		cubeNumber := math.Cbrt(float64(i))

		if sqrtNumber == float64(int(sqrtNumber)) && cubeNumber == float64(int(cubeNumber)) {
			fmt.Println("SquareCube")
		} else if cubeNumber == float64(int(cubeNumber)) {
			fmt.Println("Cube")
		} else if sqrtNumber == float64(int(sqrtNumber)) {
			fmt.Println("Square")
		} else {
			fmt.Println(i)
		}
	}
}
