package main

import "fmt"

// function

func main() {
	i := Plus(1,2)
	fmt.Println(i)

	i2, _ := Div(9,3)
	// i2, i3 := Div(9,3)
	fmt.Println(i2)

	i4 := Double(1000)
	fmt.Println(i4)

	Noreturn()
}

func Noreturn() {
	fmt.Println("No return")
	return
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func Div(x,y int) (int, int) {
	q := x/y
	r := x%y
	return q, r
}

func Plus(x, y int) int {
	return x + y
}