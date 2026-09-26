package main

import "fmt"

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}

// Greet returns a greeting message
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(Greet("Go"))
}