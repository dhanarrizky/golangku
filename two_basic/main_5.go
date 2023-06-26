package main

import (
	"fmt"
	"strconv"
)

func main(){

	//convert value
	price := 100

	//"Itoa" for Int to String
	//"Atoi" for Sting to Interger

	// untuk fungsi mengeluarkan eror
	// i, err := strconv.Atoa(price)
	
	// jika ada eror tapi di kosongi maka kita bisa memberikan nilai kosong dengan underscore
	// i, _ := strconv.Atoa(price)


	fmt.Println("from int to strong : " + strconv.Itoa(price) )

	}
