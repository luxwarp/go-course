package main

import "fmt"

const sectionHeader string = `
#############################
#        05-Functions       #
#############################`

// function declaration.
// It's structure is: func (keyword) name(parameter, parameter type) return type.

// add takes two numbers and returns the summary.
func add(num1, num2 float32) float32 {
	return num1 + num2
}

func main() {
	fmt.Println(sectionHeader)
	fmt.Println("1.3 + 1 =", add(1.3, 1))
	fmt.Println("10 - 2.3 =", sub(10, 2.3))
}

// you can also declare functions after main().

// sub takes two numbers and returns the difference.
func sub(num1, num2 float32) float32 {
	return num1 - num2
}
