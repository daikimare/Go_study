package main

import (
	"fmt"
)

func main() {
	var s string = "hogefuga"
	fmt.Println(s)
	fmt.Printf("%T\n",s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	fmt.Println(`hoge
	fuga
		piyo`)
	fmt.Println("\"")
	fmt.Println(`"`)
	fmt.Println(s[0])
	fmt.Println(string(s[0]))
}