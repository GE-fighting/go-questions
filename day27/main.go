package main

import "fmt"

//type Direction int
//
//const (
//	North Direction = iota
//	East
//	South
//	West
//)
//
//func (d Direction) String() string {
//
//	return [...]string{"North", "East", "South", "West"}[d]
//}
//
//func main() {
//	fmt.Println(South)
//}

type Math struct {
	x, y int
}

var m = map[string]*Math{
	"foo": &Math{2, 3},
}

func main() {
	m["foo"].x = 4
	fmt.Println(m["foo"].x)
}
