package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "hogehoge"

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("i = %T\n",i)
}
