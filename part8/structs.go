package main

import (
	"fmt"
	"reflect"
)

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s *subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Rate:", s.rate)
	fmt.Println("Active?:", s.active)
}

func defaultSubscriber(name string) *subscriber {
	return &subscriber{
		name:   name,
		rate:   4.99,
		active: true,
	}
}

func applyDiscount(s *subscriber) {
	s.rate = 3.99
}

func main() {
	printInfo(defaultSubscriber("Abylay"))
	aigali := defaultSubscriber("Aigali")
	aigali.rate = 10
	printInfo(aigali)
	applyDiscount(aigali)
	printInfo(aigali)
	fmt.Println(reflect.TypeOf(aigali))
	fmt.Println(*aigali)
}
