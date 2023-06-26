package main

import "fmt"

func main() {

	var name1 = "eko"
	var name2 = "eko"

	var result = name1 == name2
	println(result)

	println("=====================")
	println("array")
	println("=====================")

	var names [3]int
	for i:=0; i <= 2; i++ {
		names[i] = i
		println(names[i])
	}

	//atau seperti ini
	println("=====================")
	// var nn = [3]int {
	// 	11,
	// 	22,
	// 	33,
	// }
	//fmt.Println(names)
	
}
