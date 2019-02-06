package main

import "fmt"

//sequence,sequence

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("The Outer Loop: %d\n", i)
		for j := 0; j < 4; j++ {
			fmt.Printf("\t\tThe Inner Loop: %d\n", j)
		}
	}

}
