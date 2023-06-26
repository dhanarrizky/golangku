package main

import "fmt"

func main(){

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ",v1)

	v3 := m["k3"]
	fmt.Println("v3: ",v3)

	fmt.Println("len: ", len(m))

	delete(m, "k2")
	fmt.Println("map: ", m)

	_, prs := m["k2"]
	fmt.Println("prs: ", prs)


	n:=map[string]int{"foo": 1,"bar":2}
	fmt.Println("map: ", n)




	fmt.Println("================")
	fmt.Println("range")
	fmt.Println("============")


	nums := []int{2,3,4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ",sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}


	kvs := map[string]string{"a":"apple","b":"banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k,v)
	}

	for k := range kvs {
		fmt.Println("key: ",k)
	}

	for i, c := range "go"{
		fmt.Println(i,c)
	}
}