package main

import "fmt"

func SwitchCase() {
	var score float32

	fmt.Print("Enter your score: ")
	fmt.Scan(&score)

	switch {
	case score > 100 || score < 0:
		fmt.Println("Invalid Score")

	case score >= 80:
		fmt.Println("Your score is GOOD")

	case score >= 70:
		fmt.Println("Your score is AVERAGE")

	case score >= 60:
		fmt.Println("Your score is PASS")

	default:
		fmt.Println("Your score is LOW")
	}
}
