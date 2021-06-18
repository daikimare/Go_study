package main

import "fmt"

// function
// 引数に関数を取る場合

func main()  {
	CallFunction(func() {
		fmt.Println("this is function")
	})
}

func CallFunction(f func()) {
	f()
}