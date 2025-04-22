package main

import "fmt"

type customString string

func (c *customString) printString() {
	deReferencedString := *c
	fmt.Println(deReferencedString)
}

func main() {
	var name customString = "Karthik Gowda"
	name.printString()
}
