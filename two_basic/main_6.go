package main

import (
	"fmt"
)

func multiPlay( num1 int, num2 int) int {
	return num1 * num2
}

// function
// int di belakang function digunakkan untuk mendeklaraskian apa yang akan di kembalikan
func printText() int{
	fmt.Println("this for excample")
	// semua func harus mengembalikan nilai, or should to return one value or more
	return 10
}

func main(){

	fmt.Println("this function have result to : ", printText())
	fmt.Println("function multi play : ", multiPlay(3,2))

	}
