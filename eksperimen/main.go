package main

import "fmt"

func sum(x []int, c chan string) {
	sum := 0
	for _, y := range x {
		sum += y
	}
	c <- fmt.Sprintf("siraj, %d", sum)

}

func okeGas() {
	fmt.Println("oke gas")
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan string)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y)
}
