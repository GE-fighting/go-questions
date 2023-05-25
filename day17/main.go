package main

import "fmt"

func f() (int, int) {
	return 1, 1
}

//func main.go() {
//	var x int
//	//x, _ := f()
//	x, _ = f()
//	fmt.Println(x)
//	x, y := f()
//	//x, y = f()
//	fmt.Println(y)
//
//}

//	func increaseA() int {
//		var i int
//		defer func() {
//			i++
//		}()
//		return i
//	}
//
//	func increaseB() (r int) {
//		defer func() {
//			r++
//		}()
//		return r
//	}
//
//	func main.go() {
//		fmt.Println(increaseA())
//		fmt.Println(increaseB())
//	}
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
	var a A = Work{3}
	s := a.(Work)
	fmt.Printf("%T\n", s)
	fmt.Println(s.ShowA())
	fmt.Println(s.ShowB())
}
