package main

import (
	"fmt"
)

func main() {
	/*
	* No implicit type conversion
	* Everything needs to be explicitly type casted
	 */
	var a int32 = 100
	var b int64 = int64(a)
	fmt.Println(b)

	var num int = 65
	var d string = string(rune(num))
	fmt.Println(d)
	fmt.Printf("%v, %T\n", d, d)
}
