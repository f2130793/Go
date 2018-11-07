package main

import "fmt"

func main() {
	var a  = [3]int{2,4,6}  //or a := [3]int{2,4,6}
	var b = [...]int{1,2,3}
	fmt.Printf("%T\n", b)
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%d %d\n", i,v)
	}

	for _, v := range a {
		fmt.Printf("%d\n",v)
	}
}
