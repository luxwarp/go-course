package main

import (
	"fmt"
)

const sectionHeader string = `
#############################
#       06-Multi file       #
#############################`

func main() {
	fmt.Println(sectionHeader)
	fmt.Println(greet("Mikael"))
}
