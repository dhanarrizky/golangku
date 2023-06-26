package main

import "fmt"


func sum(nums ...int){
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
}

func intSeq() func() int {
	i := 0
	return func() int {
		i ++
		return i
	}
}


func main(){

	sum(1,2)
	sum(1,2,3)


	nums :=[]int{1,2,3,4}
	sum(nums...)




	fmt.Println("=========================")
	fmt.Println("closures")
	fmt.Println("=========================")

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())


	newInts := intSeq()
	fmt.Println(newInts())
}