package main

import (
	"fmt"
)

var y = 100
var z int
var k string
var m bool
var l string = "Shaken, but not Stirred"
var r string = `James Said "Shaken, not Stirred"`

func main() {
	x := 42
	fmt.Println("Hello, playground")
	fmt.Println(x)
	var y = 43
	fmt.Println("Value of y is ", y)
	fee()
	fmt.Println(z)
	fmt.Printf("%T\n", k)
	fmt.Printf("%T\n", m)
	fmt.Printf("%T\n", z)
	fmt.Println(r)

}

func fee() {

	fmt.Println("Again y :", y)
}
