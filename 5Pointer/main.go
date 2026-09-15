package main

import "fmt"

func print(numbers *[4]int){
	fmt.Println(numbers)
}
func main() {

	arr := [4]int{1,2,3,4}
	print(&arr)

	// x := 10
	// addr := &x

	// fmt.Println(addr) // address of x
	// fmt.Println(*addr) //value of addr

	// *addr = 30 // update the value of x
}
