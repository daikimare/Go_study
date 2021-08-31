package main

import (
	"fmt"
)

func main() {
	// n := 1
	// switch n {
	// case 1,2:
		// fmt.Println("1 or 2")
	// case 3,4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("I can not understand")
	// }

	// switch n:=2; n {
	// case 1,2:
	// 	fmt.Println("1 or 2")
	// case 3,4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("I can not understand")
	// }

	// n := 6
	// switch {
	// case n > 0 && n < 4:
	// 	fmt.Println("0 < n < 4")
	// case n > 3 && n <7:
	// 	fmt.Println("3 < n < 7")
	// }

	/* type asertion */
	anything("hoge")
	anything(1)

	var x interface{} = 3
	i := x.(int)
	fmt.Println(i + 2)
}

func anything(a interface{}) {
	fmt.Println(a)
}
