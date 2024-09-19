package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Println("Enter your name:")
	fmt.Scanln(&name)
	fmt.Println("Enter you age:")
	fmt.Scanln(&age)
	fmt.Println("Your name is:", name, "Age is: ", age)
}
