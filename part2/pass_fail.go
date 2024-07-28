package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('5')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.TrimSpace(input))
}
