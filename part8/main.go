package main

import (
	"fmt"
	"head_first/magazine"
)

//
//type SomeStruct struct {
//	someInt int
//}

func main() {
	subscriber := magazine.Subscriber{
		Name: "Sanzhar",
	}
	subscriber.Street = "Dostyk 12"
	subscriber.PostalCode = "01"
	fmt.Println(subscriber)

	employee := magazine.Employee{
		"Alem",
		500000,
		magazine.Address{
			"Magilik el",
			"Astana",
			"KZ",
			"001",
		},
	}
	fmt.Println(employee)
}
