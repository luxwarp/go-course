package main

import "fmt"

const sectionHeader string = `
#############################
#        02-Hello you       #
#############################`

func main() {
	fmt.Println(sectionHeader)

	myName := "Luxwarp"
	fmt.Printf("Hello, my name is %s .\n", myName)
	fmt.Print("WHat is your name? : ")
	var yourName string
	fmt.Scanln(&yourName)
	fmt.Println("Nice to meet you", yourName)
}
