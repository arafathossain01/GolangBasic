package main // every go program is a part of package. main is a package.

import (
	"custom.com/mathlib" // import custom package.
	"fmt"                //import the fmt from the main package.
)

const PI = 3.1416 // Global variable
func main() {
	fmt.Println("Hello World")

	/*---------Variable----------*/

	// Variable declaration with type
	var num int = 30
	fmt.Println(num)

	// Variable declaration without type
	var num2 = 53
	fmt.Println(num2)

	// Short variable declaration
	age := 32

	// Reassigning value
	age = 40
	name := "Arafat"
	fmt.Println(age)
	fmt.Println(name)

	// Multiple variable declaration
	var fname, id = "Hasan", 231
	fmt.Println(fname, id)

	// Multiple variables using var block
	var (
		num3 int    = 30
		str  string = "Hello"
	)

	fmt.Println(num3)
	fmt.Println(str)

	// Constant
	const p = 30
	fmt.Println(p)

	// Zero values
	var a string
	var b int
	var c bool
	fmt.Println(a) // ""
	fmt.Println(b) // 0
	fmt.Println(c) // false

	// variable shadowing
	t := 49

	if true {
		fmt.Println("-------Variable Shadowing-------")
		t := 69 // shadowing
		fmt.Println(t)
	}
	fmt.Println(t)

	/*---------if else & switch case----------*/
	// ifelse()
	// SwitchCase()

	/*------- Package Scope -------*/
	fmt.Println("----------Custom Pckage----------")
	mathlib.Add(num, num2)

	/*---------Function----------*/
	fmt.Println("----------Function----------")
	// standard function
	add(30, 45)

	// returnd function
	sum := add2(30, 20)
	fmt.Println("Sum (Returned function):", sum)
	addition, multiplication := getNumbers(5, 6)
	fmt.Println(addition, multiplication)

	//anonymous function or IIFE function
	func(x int, y int) {
		z := x + y
		fmt.Println("Anonymous function: ", z)
	}(10, 20)

	// function expression
	mul(10, 20)
}
