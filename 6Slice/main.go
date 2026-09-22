package main

import "fmt"

func variadic (number ...int){
	fmt.Println(number)
	fmt.Println(len(number))
	fmt.Println(cap(number))
}
func main() {
	arr := [5]string{"This", "is", "go", "interview", "question"}
	fmt.Println(arr)

	// slice from array
	fmt.Println("===== First Slice =====")
	slc := arr[1:4] // [is go interview]
	fmt.Println(slc)
	fmt.Println(len(slc))
	fmt.Println(cap(slc))

	// slice from slice
	fmt.Println("===== Second slice =====")
	slc1 := slc[1:3]
	fmt.Println(slc1)
	fmt.Println(len(slc1))
	fmt.Println(cap(slc1))

	//slice literal
	fmt.Println("===== Third Slice =====")
	s := []int{1, 2, 3, 4}
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	// using make function
	fmt.Println("===== Fourth Slice =====")
	slc2 := make([]int, 4) // make(type, size)
	fmt.Println(slc2)
	fmt.Println(len(slc2))
	fmt.Println(cap(slc2))

	slc3 := make([]int, 4, 6) // make(type, size, capacity)
	slc3[3] = 56              // len: 4, cap: 5
	fmt.Println(slc3)
	fmt.Println(len(slc3))
	fmt.Println(cap(slc3))

	// empty slice or nil slice
	fmt.Println("===== Empty Slice =====")
	var si []int
	fmt.Println(si)
	si = append(si, 3,4,5,6,7,6)
	fmt.Println(si)

	// variadic function
	fmt.Println("===== Variadic Function =====")
	variadic(1,2,3,4,5,6)

}
