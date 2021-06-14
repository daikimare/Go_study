package main

// 型
// 整数型

import (
	"fmt"
)

func main() {
	var i,i2 int = 100, 50

	fmt.Println(i+i2)
	fmt.Printf("%T\n",i2)
	/* %Tは型を憑依するためのパラメータ */

	fmt.Printf("%T\n", int32(i2))
}