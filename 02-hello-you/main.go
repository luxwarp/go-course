package main

import "fmt"

func main() {
	myName := "Luxwarp"
	fmt.Printf("Hello, my name is %s .\n", myName)
	fmt.Print("WHat is your name? : ")
	var yourName string
	fmt.Scanln(&yourName)
	fmt.Println("Nice to meet you", yourName)
}
