package main

import (
	"fmt"
)

func main() {
	a := 10
	b := "Golang"
	c := "poppysmoke"
	d := "4.17"
	e := "true"
	var g string //declaring variable with identifier  "b" as type string
	g = "cowboy" //

	//Declare and assign at the same time. It is called Initialization
	//var g string = "cowboy"

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", g)
}
