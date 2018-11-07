package main

import "fmt"

func main()  {
	s := []int{0,1,2,3,4,5,6}

	//向左移动两个元素
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s)
}

func reverse (s []int) {
	for i, j :=0, len(s) - 1; i<j;i,j = i+1,j-1 {
		s[i],s[j] = s[j],s[i]
	}
}
