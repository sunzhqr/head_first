package main

import "fmt"

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk")
}

type Whistle string

func (w Whistle) MakeSound() error {
	fmt.Println("Tweet")
	return nil
}

func (w Whistle) Bla() {
	fmt.Println("Blaaaa")
}

type NoiseMaker interface {
	MakeSound()
}

func play(n NoiseMaker) {
	n.MakeSound()
}
func main() {
	var toy NoiseMaker
	toy = Whistle("Whistle")
	fmt.Println(toy)
	toy.MakeSound()
	toy = Horn("Horn")
	fmt.Println(toy)
	toy.MakeSound()
	play(Whistle("bye"))
}
