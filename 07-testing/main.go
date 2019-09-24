package main

import (
	"fmt"
)

func greeter(name string) (greet string) {
	greet = "Hello " + name
	return greet
}

func main() {
	fmt.Println("Testing")
	fmt.Println(greeter("mikael"))
}
