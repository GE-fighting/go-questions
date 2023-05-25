package main

import "fmt"

// type S struct {
// }
//
// func f(x interface{}) {
// }
//
// func g(x *interface{}) {
// }
//
//	func main.go() {
//		s := S{}
//		p := &s
//		f(s) //A
//		g(s) //B
//		f(p) //C
//		g(p) //D
//	}
type S struct {
	m string
}

func f() *S {
	return &S{
		"foo",
	} //A
}

func main() {
	p := f()         //B
	fmt.Println(p.m) //print "foo"
}
