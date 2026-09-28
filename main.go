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

// Divide returns the quotient of two integers or an error if dividing by zero
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(Greet("Go"))
}