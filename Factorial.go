package main

import "fmt"

func factorial(n int) int {
	if n == 0 || n == 1 {
		fmt.Println(n)
		return 1
	}
	return n * factorial(n-1)
}
func main() {
	var num int
	fmt.Println("Enter a Number")
	fmt.Scanln(&num)
	result := factorial(num)
	fmt.Printf("The factorial of %d is %d\n", num, result)
}
