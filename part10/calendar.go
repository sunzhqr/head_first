package main

import (
	"fmt"
	"head_first/calendar"
	"log"
)

func main() {
	event := calendar.Event{}
	err := event.SetTitle("My birthday!")
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(4)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title())
	fmt.Println(event.Month())
}
