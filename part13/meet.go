package main

import "fmt"

func greeting(channel chan string) {
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	channel <- "Hello!"
}

func main() {
	myChannel := make(chan string)
	go greeting(myChannel)
	myChannel <- "hello"
	myChannel <- "hellower"
	fmt.Println(<-myChannel)
}
