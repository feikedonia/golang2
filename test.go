package main

import (
	"fmt"
)

var name int

func main() {
	fmt.Println("This is a test\n what is you name?")
	fmt.Scanln(&name)
	sub()
}

func sub() {
	fmt.Println("The program shall end\n", name)
}
