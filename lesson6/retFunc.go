package main

import "fmt"

// function

func ReturnFunc() func() {
	return func ()  {
		fmt.Println("this is function")
	}
}

func main() {
	f := ReturnFunc()
	f()
}