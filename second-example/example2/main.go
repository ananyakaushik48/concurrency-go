// This example has a race condition
package main

import ( 
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	wg.Done()
	msg = s
}

func main() {
	msg = "Hello world"
	
	// add two items for two decrements of the waitgroup
	wg.Add(2)
	go updateMessage("Hello universe")
	go updateMessage("hello cosmos")
	// wait for decrement of waitgroup
	wg.Wait()

	fmt.Println(msg)
}







// This example has a race condition
// package main

// import ( 
// 	"fmt"
// 	"sync"
// )

// var msg string
// var wg sync.WaitGroup

// func updateMessage(s string) {
// 	wg.Done()
// 	msg = s
// }

// func main() {
// 	msg = "Hello world"
	
// 	// add two items for two decrements of the waitgroup
// 	wg.Add(2)
// 	go updateMessage("Hello universe")
// 	go updateMessage("hello cosmos")
// 	// wait for decrement of waitgroup
// 	wg.Wait()

// 	fmt.Println(msg)
// }
