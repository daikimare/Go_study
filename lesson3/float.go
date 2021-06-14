package main

// 型
// 浮動小数点型

import (
	"fmt"
)

func main() {
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2
	fmt.Println(fl64+fl)
	fmt.Printf("%T, %T\n", fl64, fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n",fl32)

	fmt.Printf("%T\n", float64(fl32))

	hoge := 0
	sum  := 1.0 + hoge
	fuga := -1

	fmt.Println(sum)
	fmt.Printf("%T\n", float32(fuga))
}