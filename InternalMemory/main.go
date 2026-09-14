package main

import "fmt"

const a = 100
var p = 100

func outer() func(){
	money := 100
	age := 30

	fmt.Println("Age =",age)

	show := func(){ // clouser
		money = money + a + p
		fmt.Println(money)
	}

	return show
}

func call(){
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
}

func main(){
	call()
}

func init(){
	fmt.Println("=== Blank ===")
}






















// var a = 10

// func add(x int, y int) {
// 	z := x + y
// 	fmt.Println(z)
// }
// func main() {
// 	add(5, 6)
// 	add(7, 8)
// }

// func init() {
// 	fmt.Print("Init function")
// }
