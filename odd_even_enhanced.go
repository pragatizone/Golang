package main

import "fmt"

func main() {
	var num int
	fmt.Println("Please enter a positive integer..")
	fmt.Scanln(&num)
	if num%2 == 0 {
		fmt.Println("Number ", num, "is an even number")
	} else {
		fmt.Println("Number ", num, "is an odd number")
	}
}
