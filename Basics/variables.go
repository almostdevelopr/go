package main

import "fmt"

// Global variable
// always declare with var (keyword)
var Balance = 100

// Package variable accessed across packages
var balanceLeft int = 123

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

	// Local variable
	// can be declared using short hand
	num := 69
	fmt.Println("Value of num:", num)

	// Package variables print
	fmt.Println("New balance:", balanceLeft)

	// Global variable's value changed
	Balance = 345
	fmt.Println("New Balance:", Balance)

	fmt.Printf("Value of Balance: %v, Type of Balance: %T", Balance, Balance)
}
