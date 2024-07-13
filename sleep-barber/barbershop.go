package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

// methods for BarberShop
func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients", barber)

		for {
			// if there are no clients the barber goes to sleep
			if len(shop.ClientsChan) == 0 {
				color.Yellow("There is nothing to do %s takes a nap", barber)
				isSleeping = true
			}
			// we only check if the shop is open with the bool that tells us whether something was written to the client channel
			// Here shopOpen is the ok of the ClientsChan channel
			client, shopOpen := <-shop.ClientsChan
			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up", client, barber)
					isSleeping = false
				}
				// cut hair
				shop.cutHair(barber, client)
			} else {
				// shop is closed send barber home
				shop.sendBarberHome(barber)
				return
			}
		}
	}()
}


func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s is finished cutting %s's hair", barber, client)
}

func (shop *BarberShop) sendBarberHome(barber string) {
	color.Red("%s is going home", barber)
	shop.BarbersDoneChan <- true
}


func (shop *BarberShop) closeShopForDay() {
	color.Cyan("Closing shop for day.")
	close(shop.ClientsChan)
	shop.Open = false
	// This will block this function from completing till
	// the barbers are done
	for a := 1; a <= shop.NumberOfBarbers; a++ {
		// this value is set in the sendBarberHome function
		<- shop.BarbersDoneChan
	}

	close(shop.BarbersDoneChan)

	color.Green("The barbershop is now closed for the day")
	color.Green("----------------------------------------")

}


func (shop *BarberShop) addClient(client string) {
	// prints out a message
	color.Green("*** %s arrives", client)

	if shop.Open {
		select {
			
		}
	} else {
		color.Red("The shop is already closed, so %s leaves!", client)
	}
}
