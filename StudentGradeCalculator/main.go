package main

import "fmt"

func main() {
	fmt.Println("------Take Student info------")

	firstName, lastName, ID := getStudentInfo()

	bangla, english, math, physics := getMarks()

	if !isValidMarks(bangla) ||
		!isValidMarks(english) ||
		!isValidMarks(math) ||
		!isValidMarks(physics) {
		fmt.Println("Input invalid. Marks must be between 0 to 100")
		return
	}

	total := totalMarks(bangla, english, math, physics)
	average := averageMarks(total, 4)
	grade := grade(average)
	status := passFail(bangla, english, math, physics)
	
	display(firstName, lastName, ID, total, average, grade, status)
}
