package main

import (
	"fmt"
)

func main() {
	a := 0
	if a==2 {
		fmt.Println("two")
	}else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I don't know this value")
	}

	x := 0
	if x := 2; true {
		fmt.Println(x)
	}
	fmt.Println(x)
}
