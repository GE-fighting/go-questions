package main

import "fmt"

//type Int int
//
//func (i Int) PrintInt() {
//	fmt.Println(i)
//}
//
//func main.go() {
//	var i Int = 1
//	i.PrintInt()
//}

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo People = &Student{}
	think := "speak"
	fmt.Println(peo.Speak(think))
}
