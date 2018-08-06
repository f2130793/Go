package main

import (
	"os"
	"fmt"
	"strings"
)

func main()  {
	//s, sep := "", ""
	//for _, arg := range os.Args[1:] {
	//	s += sep + arg
	//	sep = " "
	//}
	//fmt.Println(s)

	//一种更高效的方式
	fmt.Println(strings.Join(os.Args[1:], ","))
}