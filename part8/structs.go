package main

import (
	"fmt"
)

type subsciber struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s subsciber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Rate:", s.rate)
	fmt.Println("Active?:", s.active)
}

func defaultSubscriber(name string) subsciber {
	return subsciber{
		name:   name,
		rate:   4.99,
		active: true,
	}
}

func main() {
	printInfo(defaultSubscriber("Abylay"))
	aigali := defaultSubscriber("Aigali")
	aigali.rate = 10
	printInfo(aigali)
}
