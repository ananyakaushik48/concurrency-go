package main

import (
	"fmt"
	"time"
)

func listenToCHan(ch chan int) {
	for {
		i := <-ch
		fmt.Println("Got", i, "from channel")

		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan int, 10)

	go listenToCHan(ch)

	for i := 0; i <= 100; i++ {
		fmt.Println("something", i, "to channel ...")
		ch <- i
		fmt.Println("sent", i, "to channel")
	}

	fmt.Println("Done!")
	close(ch)
}
