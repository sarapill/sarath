package main

import "fmt"

func subtract(a int, b int) int {
	return a - b
}

func main1() {
	var num1, num2 int
	fmt.Println("Enter two numbers to subtract:")
	fmt.Scanln(&num1, &num2)

	result := subtract(num1, num2)
	fmt.Printf("The result of subtracting %d from %d is: %d\n", num2, num1, result)
}
