package main

import "fmt"

const Max = 3

func main()  {
	a := []int {10,100,200}

	var i int

	for i = 0; i < Max; i++ {
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}
}