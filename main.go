package main

import "fmt"

type customString string

func (c *customString) log() {
	deReferencedString := *c
	fmt.Println(deReferencedString)
}

func main() {
	var name customString = "Karthik Gowda"
	name.log()
}
