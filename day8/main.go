package main

//func hello() []string {
//	return nil
//}
//
//func main() {
//	h := hello
//	fmt.Printf("%T\n", h)
//	if h == nil {
//		fmt.Println("nil")
//	} else {
//		fmt.Println("not nil")
//	}
//}

func GetValue() interface{} {
	return 1
}

func main() {
	i := GetValue()
	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
}
