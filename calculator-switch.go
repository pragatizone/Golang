package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	var result string

	fmt.Println("Enter your first Number: ")
	fmt.Scanln(&num1)

	fmt.Println("Enter your Second Number: ")
	fmt.Scanln(&num2)

	fmt.Println("Enter operation(add,subtract,divide,multiply):")
	fmt.Scanln(&result)

	switch result {
	case "add":
		println("result:", num1+num2)
	case "subtract":
		println("result:", num1-num2)
	case "multiply":
		println("result:", num1*num2)
	case "divide":
		println("result:", num1/num2)
	default:
		fmt.Println("Error:Invalid Operation")
	}
}
