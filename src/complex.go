package main

import "fmt"

func main() {
	x := complex(1, 2)
	y := complex(3, 5)
	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))
}
