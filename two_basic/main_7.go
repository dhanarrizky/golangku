package main

import (
	"fmt"
	"strconv"

)

func multiReturn(age int, name string, status string ) (string, string) {
	ageNow := strconv.Itoa(age) 
	return  " is " + name + " status " + status ,
	 "have age " + ageNow
}

func getBio(age int, name string, status string ) string {
	ageNow := strconv.Itoa(age) 
	return name + "is" + "status" + status + "have age " + ageNow
}



func main(){

	fmt.Println(getBio(12, "dhanar", "rizky"))
	fmt.Println(getBio(1, "krisnadhy", "ikky"))
	fmt.Println(getBio(3, "hey", "you"))

	fmt.Println("=============================================")

	basicInfo , ageInfo := multiReturn(12, "dhanar", "rizky")
	fmt.Println(basicInfo, ageInfo)
	

	}
