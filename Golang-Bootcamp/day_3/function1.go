package main

import "fmt"

func multiply(a, b int) int { //function declaration
	return a * b
}

func add(a, b int) int {
	return a + b
}

func main() {
	result := multiply(5, 10) //function calling
	fmt.Printf("multiplication: %d", result)

	//Call by Value
	x := 10
	y := 20
	result1 := add(x, y)

	fmt.Printf("\nSum=%d", result1)
}
