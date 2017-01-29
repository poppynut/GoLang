package main

import (
	"fmt"
)

func main() {
	a := 10
	b := "Golang"
	c := "poppysmoke"
	d := 4.17
	e := true
	var g string //declaring variable with identifier  "b" as type string
	g = "cowboy" //

	//Declare and assign at the same time. It is called Initialization
	//var g string = "cowboy"

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", g)
}
