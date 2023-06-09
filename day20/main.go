package main

import "fmt"

//	func f() {
//		defer fmt.Println("D")
//		fmt.Println("F")
//	}
//
//	func main.go() {
//		f()
//		fmt.Println("M")
//	}
type Person struct {
	age int
}

func main() {
	person := &Person{28}

	// 1.
	defer fmt.Println(person.age)

	// 2.
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	// 3.
	defer func() {
		fmt.Println(person.age)
	}()

	person = &Person{29}
}
