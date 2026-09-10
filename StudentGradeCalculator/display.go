package main

import "fmt"

func display(fName string, lName string, ID int, total float32, avg float32, grade string, status string) {
	fmt.Println("-------   RESULT   -------")
	fmt.Println("Name		: ", fName +" "+ lName)
	fmt.Println("ID	: ",ID)
	fmt.Println("Total	: ",total)
	fmt.Printf("Average	: %.2f\n",avg)
	fmt.Println("Grade	: ",grade)	
	fmt.Println("Status	: ",status)
	fmt.Println("-------     END     -------")

}
