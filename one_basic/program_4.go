package main

import "fmt"

func main(){

	var a[5]int
	fmt.Println("emp: ", a)

	a[4]=100
	fmt.Println("set: ", a)
	fmt.Println("get: ",a[4])

	fmt.Println("len: ",len(a))

	b := [5]int{1,2,3,4,5}
	fmt.Println("dcl: ",b)

	var twoD [2][3]int
	for i :=0; i < 2;i++ {
		for j := 0 ;j <3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ",twoD)

	fmt.Println("==============================")
	fmt.Println("slices")
	fmt.Println("==============================")

	var s[]string
	fmt.Println("uninit: ",s,s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp: ",s ,"len: ",len(s), "cap: ", cap(s))




	fmt.Println("================================")
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	fmt.Println("len: ", len(s))

	s = append(s, "d")
	fmt.Println(s)
	s = append(s, "e","f")
	fmt.Println(s)
}