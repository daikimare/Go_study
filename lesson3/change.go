package main

import (
	"fmt"
	"strconv"
)

// 型変換

func main() {
	var i int = 1
	fl64 := float64(i)

	fmt.Println(fl64)
	fmt.Printf("fl64 : %T\n", fl64)
	fmt.Printf("i : %T\n", i)

	i2 := int(fl64)
	fmt.Printf("i2 : %T\n", i2)

	fl := 10.5
	i3 := int(fl)
	fmt.Printf("fl : %T\n", fl)
	fmt.Println(fl)
	fmt.Printf("i3 : %T\n", i3)
	fmt.Println(i3)

	// strig 
	var s string = "100"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	ci, _ := strconv.Atoi((s))
	fmt.Println(ci)
	fmt.Printf("ci = %T\n", ci)

	var ci2 int = 200
	s2 := strconv.Itoa(ci2)
	fmt.Println(s2)
	fmt.Printf("%T\n", s2)

	var h string = "Hello world"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Println(h2)
}