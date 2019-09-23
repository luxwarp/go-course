package main

import "fmt"

// function declaration.
func add(num1, num2 float32) float32 {
	return num1 + num2
}

// this is the initial function that always gets run.
func main() {
	fmt.Println("Functions")
	fmt.Println("1.3 + 1 =", add(1.3, 1))
	fmt.Println("10 - 2.3 =", sub(10, 2.3))
}

// you can also declare functions after main().
func sub(num1, num2 float32) float32 {
	return num1 - num2
}
