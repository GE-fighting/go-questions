package main

import (
	"fmt"
	"time"
)

//func main() {
//	v := []int{1, 2, 3}
//	for i := range v {
//		v = append(v, i)
//	}
//}

func main() {

	var m = [...]int{1, 2, 3}

	for i, v := range m {
		go func(x, y int) {
			fmt.Println(x, y)
		}(i, v)
	}

	time.Sleep(time.Second * 3)
}
