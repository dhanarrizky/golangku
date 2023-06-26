package main

import "fmt"

func main(){
	var total int

	a := 4
	b := 5

	total = a + b
	
	fmt.Println("total : ", total)

	b++
	total = a + b
	fmt.Println("total : ", total)

	}
