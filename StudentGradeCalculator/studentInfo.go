package main

import "fmt"

func getStudentInfo()(string, string, int){
	var fName string 
	var lName string 
	var ID int

	fmt.Print("Enter your first name: ")
	fmt.Scan(&fName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lName)
	fmt.Print("Enter your ID: ")
	fmt.Scan(&ID)

	return fName , lName, ID
	
}