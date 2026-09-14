package main

import "fmt"

// Structure for create user define variable | read only never update again
type User struct {
	Name string // member variable or property
	Age  int
}

// recevier function
func (usr User) printDetails() {
	fmt.Println("Name: ", usr.Name)
	fmt.Println("Age: ", usr.Age)
	fmt.Println("First call done.")
}

// recevier function with parameter
func (usr User) print(a int) {
	fmt.Println("Name: ", usr.Name)
	fmt.Println("Age: ", usr.Age)
	fmt.Println("Number: ", a)
}
func main() {
	var User1 User

	User1 = User{ // instanse or object
		Name: "Arafat",
		Age:  21,
	}

	User1.printDetails()

	User2 := User{
		Name: "Kakoli",
		Age:  20,
	}

	User2.print(20)
}
