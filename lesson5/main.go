package main

import "fmt"

// 算術演算

func main() {
	fmt.Println(1+2)
	fmt.Println(5-1)
	fmt.Println(5*4)
	fmt.Println(60/3)
	fmt.Println(10%3)

	fmt.Println("hoge"+"fuga")

	n := 0
	n+= 4
	fmt.Println(n)
	n++
	fmt.Println(n)
	n--
	fmt.Println(n)
	s := "hoge"
	s += "fuga"
	fmt.Println(s)
}