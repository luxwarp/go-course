package main

import "fmt"

const sectionHeader string = `
#############################
#        03-Variables       #
#############################`

// Global variables, accessable from all packages in main.
var name string = "Mikael Luxwarp Carlsson"
var age int = 29

func main() {
	fmt.Println(sectionHeader)

	fmt.Printf("Name value %s\n", name)
	fmt.Printf("Name type: %T\n", name)
	fmt.Printf("Age value: %d\n", age)
	fmt.Printf("Age type: %T\n", age)

	// short declaration variable only accessable within this function.
	color := "blue"
	fmt.Println("Color:", color)
}
