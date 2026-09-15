package main

import "fmt"

var arr2 = [3]string{"I" , "Love", "You"}

func arrPrint(numbers [3]int){
	fmt.Println("From function:",numbers)
}
func main() {
	var arrName [3]int
	arrName[0] = 3
	arrName[1] = 23
	fmt.Println(arrName)
	arrPrint(arrName)

	arr1 := [2]int{2, 1}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr2[1])

}
