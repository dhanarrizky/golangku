package main

import "fmt"

// empty variable
// var school = ""
var school string

// introduce var in golang language

var firstName, LastName = "Dhanar", "Rizky"

var fullName = firstName + " " + LastName

func main(){
	
	schoolTwo := "private"

	x, y := "Dhanar", "Rizky"

	z := x + " " + y

	fmt.Println("two both declaration above have same function for declaration empty var : " + school)
	fmt.Println("my full name" + fullName)
	fmt.Println("var short name : " + schoolTwo)
	fmt.Println("print z : " + z)


	}
