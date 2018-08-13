package main

import "fmt"

func main() {
	var numbers = make([]int,3,5)

	printSlice1(numbers)
}

func printSlice1(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}