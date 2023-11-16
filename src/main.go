package main

import (
	"fmt"
)

// c chan <- // Dato de entrada
// c <- chan // dato de salida

func say(text string, c chan <- string) {
	c <- text
}

func main() {

	c := make(chan string, 1)

	fmt.Println("Hello")

	go say("bye", c)

	fmt.Println(<- c)
}
