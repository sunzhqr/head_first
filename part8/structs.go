package main

import "fmt"

func main() {
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}
	fmt.Printf("%#v\n", myStruct)
	myStruct.number = 3.14159
	myStruct.word = "Go"
	myStruct.toggle = true
	fmt.Printf("%#v\n", myStruct)
}
