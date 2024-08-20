package main

import "fmt"

type Stinger string

func (s Stinger) greeting(greet Stinger) Stinger {
	return greet + s
}

func AcceptEverything(a interface{}) {
	fmt.Println(a.MakeSound())
}

func main() {
	deadpool := Stinger("Chimichango")
	fmt.Println(deadpool.greeting("HIIIIIIIIIIIIIIIII "))
	wolverine := Stinger("imbicile")
	fmt.Println(wolverine.greeting("Bye "))
	AcceptEverything(true)
}
