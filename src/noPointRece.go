package main

import (
	"fmt"
)

// 定义点结构
type Point struct {
	X int
	Y int
}

// 非指针接收器的加方法
func (p Point) Add(other Point) Point {
	// 成员值与参数相加后返回新的结构
	return Point{p.X + other.X, p.Y + other.Y}
}
func main() {
	// 初始化点
	p1 := Point{1, 1}
	p2 := Point{2, 2}
	// 与另外一个点相加
	result := p1.Add(p2)
	// 输出结果
	fmt.Println(result)
}

//在计算机中，小对象由于值复制时的速度较快，所以适合使用非指针接收器。大对象因为复制性能较低，适合使用指针接收器，在接收器和参数间传递时不进行复制，只是传递指针。
