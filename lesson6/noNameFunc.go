package main

import "fmt"

// function
// no name function

func main()  {

	// act any work
	f := func (x, y int) int {
		return x+y
	}
	i := f(1,2)
	fmt.Println(i)

	// act same timing
	i2 := func (x,y int) int {
		return x+y
	}(1,4)

	fmt.Println(i2)
}