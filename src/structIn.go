package main

import (
	"fmt"
)

type BasicColor struct {
	R, G, B float32
}
type Color struct {
	//Basic BasicColor
	BasicColor //结构体内嵌
	Alpha      float32
}

func main() {
	var c Color
	c.R = 1
	c.G = 1
	c.B = 0
	c.Alpha = 1
	fmt.Printf("%+v", c)
}
