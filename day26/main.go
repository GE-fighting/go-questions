package main

import "fmt"

//const (
//	a = iota
//	b = iota
//)
//const (
//	name = "name"
//	c    = iota
//	d    = iota
//)
//
//func main.go() {
//	fmt.Println(a)
//	fmt.Println(b)
//	fmt.Println(c)
//	fmt.Println(d)
//}

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func main() {

	var s *Student
	if s == nil {
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is not nil")
	}
	var p People = s
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}
}
