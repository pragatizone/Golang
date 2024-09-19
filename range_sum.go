package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Enter the start of the range")
	fmt.Scanln(&a)
	fmt.Println("Enter the end of the range")
	fmt.Scanln(&b)

	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("The sum of numbers from %d to %d is: %d\n", a, b, sum)
}
