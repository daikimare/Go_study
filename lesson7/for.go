package main

import (
	"fmt"
)

func main() {
	// i := 0
	// for {
	// 	i++
	// 	if i == 10 {
	// 		break
	// 	}
	// 	fmt.Println("hoge")
	// }

	// for i := 0; i < 10; i++ {
	// 	if i==10 {
	// 		continue
	// 	} else if i==6 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }
	
	// arr := [3]int{1,2,3}
	// for i:=0; i< len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// for _, v := range arr {
	// 	fmt.Println(v)
	// }

	/* slice */
	// sl := []string{"hoge","fuga","piyo"}
	// for i,v:=range sl {
	// 	fmt.Println(i,v)
	// }

	/* map */
	m := map[string]int{"hoge":100, "fuga":200}
	for k := range m {
		fmt.Println(k)
	}

}
