//多重赋值，最大公约数
package main

import "fmt"

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	ret := gcd(161, 322)
	fmt.Println(ret)
}

