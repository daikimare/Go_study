package main

import "fmt"

// 型
// byte (unit8)型

func main() {
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	// change string
	fmt.Println(string(byteA))

	// change byte from string
	c := []byte("HI")
fmt.Println(c)

	fmt.Println(string(c))
}