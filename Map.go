package main

import (
	"fmt"
)

func main() {
	// Input arrays
	environment := []string{"dev", "qc", "uat", "prod"}
	ec2_instance := []string{"t3.micro", "t3.small", "t3.medium", "t3.large"}

	// Create a map to store the environment and EC2 instance mapping
	envToInstance := make(map[string]string)

	// Populate the map with data from both arrays
	for i := 0; i < len(environment); i++ {
		envToInstance[environment[i]] = ec2_instance[i]
	}

	// Output the result
	fmt.Println("Environment to EC2 Instance mapping:")
	for env, instance := range envToInstance {
		fmt.Printf("%s -> %s\n", env, instance)
	}
}
