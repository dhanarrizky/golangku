package main

import (
	"fmt"

)

func getBio(age int, name string, status string ) (bio string, ageIn10 int) {
	
	ageIn10 = age + 10
	bio = name + " is " + status 

	// hanya menggunakkan return karena sudah di deklarasikan di atas apa yang akan di keluarkan/ akan di kembalikan
	return
}
var a ="ini rahasia"

// function tanpa return
func tanpaReturn()  {
	fmt.Println(a)
}


func main() {

	// nilai dari return fungsi dimasukkan kedalam suatu variable baru
	besicInfo, ageInfo := getBio(21, "dhanar", "rizky")

	fmt.Println(besicInfo)
	fmt.Println("mempunyai umur : ", ageInfo)
	fmt.Println("==========================")
	tanpaReturn()

	}
