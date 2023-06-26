package main

import (
	"fmt"
	"time"	
)

func main(){

i := 2
fmt.Println("write ", i, " as ")
switch i {
case 1:
	fmt.Println("one")
case 2:
	fmt.Println("two")
case 3:
	fmt.Println("three")
}

switch time.Now().Weekday(){
case time.Saturday, time.Sunday:
	fmt.Println("it's the weekend")
default:
	fmt.Println("it's a weekday")
}

t := time.Now()
switch{
case t.Hour() < 12:
	fmt.Println("it's before noon")
default:
	fmt.Println("it's after noon")
}

whatAmi := func(i interface{}) {
	switch t := i.(type) {
	case bool:
		fmt.Println("I'm a bool")
	case int:
		fmt.Println("I'm an int")
	default:
		fmt.Printf("Don't know type %T\n",t)
	}
}

whatAmi(true)
whatAmi(1)
whatAmi("hey")



}