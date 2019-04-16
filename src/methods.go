package main

import (
	"fmt"
	"math"
)

type Vertex struct { //定义结构体
	X, Y float64
}

//方法
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//函数
func Abs(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func main() {
	//方法
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
	//函数
	u := Abs(3, 4)
	fmt.Println(u)
}
