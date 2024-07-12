package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	rightFork int // locks essentially one access per unit time
	leftFork  int
}

// Philosophers is a list of all philosophers
var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Seneca", leftFork: 2, rightFork: 3},
	{name: "Marcus", leftFork: 3, rightFork: 4},
}

var hunger = 3                // how many times does a philosopher eat
var eatTime = 1 * time.Second // time taken to eat
var thinkTime = 3 * time.Second
var sleepTime = 3 * time.Second

// finding the order in which they finished eating
var orderMutex sync.Mutex // a mutex to keep track of who is done eating
var orderFinished []string // the order in which they finished eating


func main() {
	// eatTime = 0 * time.Second
	// sleepTime = 0 * time.Second
	// thinkTime = 0 * time.Second

	fmt.Println("Dining philosophers problem")
	fmt.Println("---------------------------")
	fmt.Println("the table is empty")

	// added sleep time
	time.Sleep(sleepTime)

	dine()

	fmt.Println("The table is empty")
}

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks is a map of all forks
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// start the meal
	for i := 0; i < len(philosophers); i++ {
		// fire off go routine for the current philosopher
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	// seat the philosopher at the table
	fmt.Printf("\t%s is seated at the table.\n", philosopher.name)
	seated.Done()

	seated.Wait()
	// eat three times
	for i := hunger; i > 0; i-- {
		// get a lock on both forks
		// this checks if the 
		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork\n", philosopher.name)
		}

		fmt.Printf("\t%s has both forks and is eating\n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking\n", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("\t%s put down the forks\n", philosopher.name)
	}

	fmt.Println(philosopher.name, "is satisfied")
	fmt.Println(philosopher.name, "left the table")


	// keeping track of the order in which they finished eating
	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher.name)
	orderMutex.Unlock()
}
