package main

import (
	"fmt"
)

const sectionHeader string = `
#############################
#         07-Testing        #
#############################`

func greeter(name string) (greet string) {
	greet = "Hello " + name
	return greet
}

func main() {
	fmt.Println(sectionHeader)
	fmt.Println(greeter("mikael"))
}
