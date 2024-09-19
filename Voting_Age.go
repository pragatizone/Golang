package main

import "fmt"

func main() {
	var age int

	// Prompt the user for their age
	fmt.Println("Enter your age:")
	fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println("You are eligible for voting.")
	} else {
		fmt.Println("You are not eligible for voting.")
	}
}
