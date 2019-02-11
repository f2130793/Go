package main

import "fmt"

func main() {
	var numbers = make([]int, 1, 5)
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(b[1:4]) // ['o','l','a']
	fmt.Println(b[:2])  //go
	fmt.Println(b[2:])  //lang
	fmt.Println(b[:])   //golang

	printSlice1(numbers)
}

func printSlice1(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
