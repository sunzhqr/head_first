package main

import (
	"fmt"
	"head_first/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	coffeepot := CoffeePot("Samsung")
	fmt.Print(coffeepot, "\n")
	fmt.Println(coffeepot)
	fmt.Printf("%s", coffeepot)
}

type ComedyError string

func (c ComedyError) Error() string {
	return string(c)
}
func TryOut(player Player) {
	player.Play("Untouchable")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not a TapeRecorder!")
	}
	var err error
	err = ComedyError("Babl")
	fmt.Println(err.Error())
}

type CoffeePot string

func (c CoffeePot) String() string {
	return string(c) + " coffee pot"
}
