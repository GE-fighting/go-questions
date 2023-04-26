package main

import (
	"fmt"
)

// 1.这道题考的是使用 append 向 slice 添加元素，第一段代码常见的错误是 [1 2 3]，需要注意
// func main() {
// 	s := make([]int, 5)
// 	s = append(s, 1, 2, 3)
// 	fmt.Println(s)
// }

func main() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s)
}

//2
// func funcMui(x,y int)(sum int,error){
// 	return x+y,nil
// }
