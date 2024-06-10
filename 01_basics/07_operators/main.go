package main

import "fmt"

func main() {
	// Penerapan operator matematika
	var value = (((2+6)%3)*4 - 2) / 3
	fmt.Printf("Result is %d \n", value)

	// Penerapan operator perbandingan
	var isEqual = (value == 2)
	fmt.Printf("nilai %d (%t) \n", value, isEqual)

	// Penerapan operator logika
	var left = true
	var right = false

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)
}
