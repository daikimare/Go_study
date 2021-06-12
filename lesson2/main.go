package main

import (
	"fmt"
)

// 変数
/* Go is static variable langage
	 */

var global int = 1000

func main()  {
	// 明示的な定義
	// var 変数名 型 = 値
	/* this syntax is used both in function to out of function */

	var hoge int = 100
	fmt.Println(hoge)

	var greet string = "Hello world"
	fmt.Println(greet)

	var t,f bool = true, false
	fmt.Println(t,f)

	var (
		num int = 200
		str string = "Hello Go"
	)
	fmt.Println(num,str)

	var fuga int
	var piyo string
	fmt.Println(fuga, piyo)

	fuga = 12345
	piyo = "Nobody's Home"
	fmt.Println(fuga, piyo)

	piyo = "Take What U Want"
	fmt.Println(piyo)

	// 暗黙的な定義
	// 変数名 := 値 
	/* not dynamic var (differece js's var) */
	/* this syntax is used in function. can not use out of function */
	num2 := 300
	fmt.Println(num2)

	num2 = 450
	fmt.Println(num2)

	// num2 = "hoge"
	// fmt.Println(num2)

	fmt.Println(global)
}