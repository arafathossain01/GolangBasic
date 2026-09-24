package main

import "fmt"

func print(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], "  ")
	}
	fmt.Println()
}

func input(arr []int, target int) int {
	fmt.Print("Enter the array element: ")

	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Print("Enter target value: ")
	fmt.Scan(&target)

	return target
}

func linearSearch(arr []int, target int) int {

	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1

}
func main() {
	var (
		size   int
		target int
	)

	fmt.Print("Enter array size: ")
	fmt.Scan(&size)

	arr := make([]int, size)

	re := input(arr, target)
	print(arr)

	result := linearSearch(arr, re)

	if result != -1 {
		fmt.Println("Found ", re, " at position: ", result+1)
	} else {
		fmt.Println("Element not found.")
	}

}
