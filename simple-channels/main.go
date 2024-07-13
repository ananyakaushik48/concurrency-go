package main

import (
	"fmt"
	"strings"
)

func shout(ping chan string, pong chan string){
	for {
		s := <-ping

		pong <-fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)
	
	fmt.Println("Type something and press ENTER (enter Q to quit)")
	for {
		// print a prompt
		fmt.Print("-> ")

		// get userInput
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}


		ping <- userInput

		response := <- pong
		fmt.Println("response:", response)
	}
	fmt.Println("all completed closing")
	close(ping)
	close(pong)
	// always close channels
}