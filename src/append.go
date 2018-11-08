package main

import "fmt"

func main() {
	//byte 等同于int8，常用来处理ascii字符
	//rune 等同于int32,常用来处理unicode或utf-8字符
	var runes []rune
	for _, r := range "Hello,世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
}
