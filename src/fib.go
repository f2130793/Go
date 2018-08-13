//斐波那契数列第n个数
//package main
//
//import "fmt"
//
//func fib(n int) int {
//	x, y := 0 ,1
//	for i := 0; i < n; i++ {
//		x, y = y, x+y
//	}
//	return x
//}
//
//func main() {
//	ret := fib(10)
//	fmt.Println(ret)
//}



package main

import "fmt"

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}
