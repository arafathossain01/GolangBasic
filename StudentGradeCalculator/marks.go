package main
import "fmt"

func getMarks ()(float32, float32, float32, float32){
	var bangla , english, math, physics float32
	

	fmt.Print("Enter your Bangla marks: ")
	fmt.Scan(&bangla)
	fmt.Print("Enter your English marks: ")
	fmt.Scan(&english)
	fmt.Print("Enter your Math marks: ")
	fmt.Scan(&math)
	fmt.Print("Enter your Physics marks: ")
	fmt.Scan(&physics)

	return bangla, english, math, physics
}