package main

import "fmt"

//	func main() {
//		fmt.Println([...]int{1} == [2]int{1})
//		fmt.Println([]int{1} == []int{1})
//	}
var p *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	//use p
	fmt.Println(*p)
}

func main() {
	var err error
	p, err = foo()
	if err != nil {
		fmt.Println(err)
		return
	}
	bar()
	fmt.Println(*p)
}
