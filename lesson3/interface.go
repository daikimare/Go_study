package main

import "fmt"

// interface

func main() {
	var x interface{}
	fmt.Println(x)
	
	x = 1
	fmt.Println(x)
	x = 3.14
	fmt.Println(x)
	x = "hoge"
	fmt.Println(x)
	x = [3]int{1,3,5}
	fmt.Println(x)
}