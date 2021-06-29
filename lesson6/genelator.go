package main

import "fmt"

// function
// genelator

func main() {
	ints := integer()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
}

func integer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}