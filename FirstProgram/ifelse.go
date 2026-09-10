package main

import "fmt"

func ifelse() {
	var marks float32
	fmt.Print("Enter your marks: ")
	fmt.Scan(&marks) // user input

	if marks > 100 || marks < 0 {
		fmt.Println("Invalid input.")
	} else if marks >= 80 {
		fmt.Println("YOu got A+")
	} else if marks >= 70 {
		fmt.Println("YOu got A")
	} else if marks >= 60 {
		fmt.Println("You got A-")
	} else if marks >= 50 {
		fmt.Println("You got B")
	} else if marks >= 40 {
		fmt.Println("You got C")
	} else if marks >= 33 {
		fmt.Println("You got D")
	} else {
		fmt.Println("You failed")
	}
}
