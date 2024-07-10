package main

import "fmt"

func printSomething(s string){
    fmt.Println(s)
}

func main() {
go printSomething("this is first")

printSomething("this is second")
}
