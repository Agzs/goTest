// fibutil project main.go
package main

import (
	"fibutil/fib"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	var n int
	fmt.Scanln(&n)
	for i := 0; i <= n; i++ {
		//t := int64(i)
		fmt.Println("Fibm(", i, ") = ", fib.Fibm(int64(i)))
	}
	//fmt.Println("Fibm(5) = ", fib.Fibm(5))
}
