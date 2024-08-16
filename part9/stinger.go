package main

import "fmt"

type Stinger string

func (s Stinger) greeting(greet Stinger) Stinger {
	return greet + s
}

func main() {
	deadpool := Stinger("Chimichango")
	fmt.Println(deadpool.greeting("HIIIIIIIIIIIIIIIII "))
	wolverine := Stinger("imbicile")
	fmt.Println(wolverine.greeting("Bye "))
}
