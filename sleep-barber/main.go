package main

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// variables
var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	// seeding random generator
	rand.NewSource(time.Now().UnixNano())

	// print welcome message
	color.Yellow("the sleeping barber problem")
	color.Yellow("---------------------------")
	// create channels
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	// create datastructure that reps the Barber shop
	// barber shop
	shop := BarberShop{
		ShopCapacity: seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan: clientChan,
		BarbersDoneChan: doneChan,
		Open: true,
	}

	color.Green("The shop is open for the day")

	// add barbers
	shop.addBarber("Frankie")
	// start the barbershop as a go routine
	shopClosing := make(chan bool)
	closed := make(chan bool)
	go func() {
		<- time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	} ()

	// add clients
	i := 1
	go func(){
		for {
			// get a random number with average arrival rate
			randomMilliseconds := rand.Int() % (2 * arrivalRate)
			select {
				// waiting to see if the shop is closed
			case <- shopClosing:
				return
			case <- time.After(time.Millisecond * time.Duration(randomMilliseconds)):
				// add a client after random milliseconds
			}
		}
	}()
	// block until the barbershop is closed
	time.Sleep(5 * time.Second)
}