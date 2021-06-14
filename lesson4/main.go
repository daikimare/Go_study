package main

import "fmt"

// 定数

const Pi = 3.14
const (
	URL = "hogehogehoge"
	Sitename = "fugafugafuga"
)
const (
	A = 1
	B
	C
	D = "piyo"
	E
	F
)
const Big = 9223372036854775807
const Big2 = Big - 1
const (
	c0 = iota
	c1
	c2
)

func main() {
	fmt.Println(Pi)
	fmt.Printf("%T\n", Pi)

	fmt.Println(URL, Sitename)
	fmt.Println(A,B,C,D,E,F)

	fmt.Println(Big)
	fmt.Println(Big2)
	fmt.Println(Big - 1)

	fmt.Println(c0,c1,c2)
}