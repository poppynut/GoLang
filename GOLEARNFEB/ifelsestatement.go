package main

import "fmt"

func main() {
	x := 43
	if x == 42 {
		fmt.Println("Value of X is : ", x)
	} else if x == 43 {
		fmt.Println("Value of x is:", x)
	} else if x == 44 {
		fmt.Println("Value of x is :", x)
	} else {
		fmt.Println("Value of x is not 42, 43, 44")
	}
}
