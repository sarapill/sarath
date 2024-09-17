package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Enter first number: ")
	fmt.Scanln(&a)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)

	sum := a + b
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, sum)
}
