package main

import "unsafe"

const (
	A = "abc"
	B = len(A)
	C = unsafe.Sizeof(A)
)

//iota 特殊常量 可以认为是一个可以被编译器修改的常量
const (
	D = iota // 0
	E        //1
	F        //2
	H = "ha" //独立值，iota += 1
	I        //"ha"
	J = 100  //100
	K        //100
	M = iota //7,恢复计数
	L        //8
)

// << 二进制位一并左移iota位
const (
	O = 1 << iota
	P = 3 << iota
	Q
	R
)

func main() {
	println(D, E, F, H, I, J, K, M, L, O, P, Q, R)
}
