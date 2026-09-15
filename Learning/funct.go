package main

import "fmt"

//standerd function
func add(num1 int, num2 int) {
	sum := num1 + num2
	fmt.Println("The sum (standard function): ", sum)
}

// return function
func add2(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2

	return sum, mul
}

// inti function
func init() {
	fmt.Println("This is init function. user can't call it.")
}

// function expression
var mul = func(a int, b int) {
	z := a * b
	fmt.Println("Multiplication from function expression: ", z)
}

// higher order function
func operation(a int, b int, add func(p int, q int)) {
	add(a, b) // add func(p int, q int) => call back function
}

func operation1() func(x int, y int) {
	return add1
}
