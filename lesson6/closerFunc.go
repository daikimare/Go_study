package main

import "fmt"

// function
// closer

func main() {
	f := Later()
	fmt.Println(f("Hello"))
	fmt.Println(f("My"))
	fmt.Println(f("name"))
	fmt.Println(f("is"))
	fmt.Println(f("go lang"))
}

func Later() func(string) string {
	var store string
	return func (next string) string {
		s := store
		store = next
		return s
	}
}