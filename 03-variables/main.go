package main

import "fmt"

// Global variables, accessable from all packages in main.
var name string = "Mikael Luxwarp Carlsson"
var age int = 29

func main() {
	fmt.Println("Variables")
	fmt.Printf("Name value %s\n", name)
	fmt.Printf("Name type: %T\n", name)
	fmt.Printf("Age value: %d\n", age)
	fmt.Printf("Age type: %T\n", age)

	// short declaration variable on accessable within this function.
	color := "blue"
	fmt.Println("Color:", color)
}
