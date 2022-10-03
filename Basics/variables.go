package main

import "fmt"

func main() {
	/*
		Method 1:
		1) Declaration
		2) Initialization
	*/
	var a int
	a = 12

	fmt.Println(a)

	/*
		Method 2:
		-> Declaration with/without datatype and Initialization
	*/
	var b int = 45
	var c = 43
	fmt.Println(b)
	fmt.Println(c)

	/*
		Method 3:
		-> Declaration without var (keyword) && datatype and Initialization
	*/
	num := 69
	fmt.Println("Value of num:", num)
}
