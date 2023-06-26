package main

import "fmt"

func main()  {

	fmt.Println("cara for pertama")

	for i:= 0; i <=2; i++ {
		fmt.Println(i)
	}


	fmt.Println("cara for kedua")

	i:= 0
	for i <=2 {
		fmt.Println(i)
	
		i++
	}

	fmt.Println("contoh penggunaan for loop lain")
	total := 0
	for i := 0; i <= 4; i ++ {
		total += i
	}
	println("hasilnya adalah : ", total)

}