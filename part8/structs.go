package main

import "fmt"

type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
}

func main() {
	showInfo(car{"Aventador", 455})
	fmt.Printf("% 	v\n", minimiumOrder("Hex bolts"))
}

func showInfo(c car) {
	fmt.Println("This is ", c.name, "with top speed of ", c.topSpeed)
}

func minimiumOrder(descripiton string) part {
	var p part
	p.description = descripiton
	p.count = 100
	return p
}
