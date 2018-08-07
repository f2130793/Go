package main

var x, y int
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

const  aa= "hello"

var c, d = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"


func main() {
	g, h := 123, "hello"
	k := 50

	println(x, y, a, b, c, d, e, f, g, h, k, aa)
}
