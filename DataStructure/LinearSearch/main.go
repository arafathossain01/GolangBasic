package main
import "fmt"

func main(){
	var size int
	fmt.Print("Enter size: ")
	fmt.Scan(&size)
	arr := make([]int, size)

	fmt.Print("Enter elements: ")
	for i:=0; i<size; i++{
		fmt.Scan(&arr[i])
	}
	fmt.Print("Enter search value: ")
	var target int
	fmt.Scan(&target)

	result := lSearch(arr, target)

	if result != -1{
		fmt.Println("Find the element: ",target, " at position: ",result+1)
	}else{
		fmt.Println("Element not found.")
	}
}