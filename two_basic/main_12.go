package main

import "fmt"

func main() {

	// pointer untuk mengambil alamat dari suatu variable

	// menggunakkan dan (&)
	// lalu untuk mengambil value dari alamat yang di panggil maka dengan menambahkan tanda bintang
	var poin = 100
	var poinAddrass *int = &poin

	fmt.Println(poinAddrass)
	fmt.Println("================")
	fmt.Println(*poinAddrass)

	fmt.Println("================")
	*poinAddrass = 200
	fmt.Println(poin)

}
