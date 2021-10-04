package main

import "fmt"

func main()  {
	fmt.Println("Hello world")
	n := 123
	p := &n
	m := 10000
	p2 := &m

	fmt.Println("p value: %d, address: %p",*p,p)
	fmt.Println("p2 value: %d, address: %p",*p2,p2)
}
