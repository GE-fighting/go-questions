package main

import "fmt"

// func main.go() {
//
//		s := [3]int{1, 2, 3}
//		a := s[:0]
//		b := s[1:2]
//		c := s[1:2:cap(s)]
//		fmt.Printf("a's len is %d,cap is %d \n", len(a), cap(a))
//		fmt.Printf("b's len is %d,cap is %d \n", len(b), cap(b))
//		fmt.Printf("c's len is %d,cap is %d \n", len(c), cap(c))
//	}

//func main.go() {
//	var m map[string]int //A
//	m = make(map[string]int)
//	m["a"] = 1
//	if v, _ := m["b"]; v != 0 { //B
//		fmt.Println(v)
//	}
//}

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func main() {
	c := Work{3}
	var a A = c
	var b B = c
	fmt.Println(a.ShowA())
	fmt.Println(b.ShowB())
}
