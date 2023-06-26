package main

import(
	"fmt"
	"math"
)
const s string = "constant"

func main(){
	fmt.Println(s)

	const n = 50000000000


	const d =3e20/n
	fmt.Println(d)

	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

	fmt.Println("===============================")
	fmt.Println("for loop")
	fmt.Println("===============================")

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i+1
	}

	for j :=7;j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n :=0 ;n<=5;n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	fmt.Println("============================")
	fmt.Println("if else")
	fmt.Println("============================")

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9 ; num < 0 {
		fmt.Println(num , " is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}