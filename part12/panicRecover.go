package main

import "fmt"

func calmDown() {
	r := recover()
	if err, ok := r.(error); ok {
		fmt.Println(err.Error())
	}
	fmt.Println("dsdsds")
}

func freakOut() {
	defer calmDown()
	panic("oh no")
}

func main() {
	defer calmDown()
	err := fmt.Errorf("WTF")
	panic(err)
}
