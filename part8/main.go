package main

import "fmt"

type SomeStruct struct {
	someInt int
}

func main() {
	var s SomeStruct
	s.someInt = 42
	fmt.Println(s)
	fmt.Println(s.someInt)
	var point *SomeStruct = &s
	fmt.Println(point)
	fmt.Println(*point)
	fmt.Println((*point).someInt)
	fmt.Println(point.someInt)

	point.someInt = 5
	fmt.Println(s.someInt)
	fmt.Println(point.someInt)

}
