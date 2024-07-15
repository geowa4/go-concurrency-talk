package slides

import (
	"fmt"
	"math/rand"
	"time"
)

// G Introducing select
func G() {
	alice := sourceG()
	defer close(alice)
	bob := sourceG()
	defer close(bob)
	timeout := time.After(10 * time.Second)
	sinkG(timeout, alice, bob)
}

func sourceG() chan int {
	c := make(chan int)
	go func() {
		for {
			sleepDuration := time.Duration(rand.Intn(3000)) * time.Millisecond
			//fmt.Println(fmt.Sprintf("%s is sleeping for %d", name, sleepDuration))
			time.Sleep(sleepDuration)
			c <- rand.Intn(1000)
		}
	}()
	return c
}

func sinkG(timeout <-chan time.Time, alice, bob <-chan int) {
	end := false

	for !end {
		select {
		case t := <-timeout: // time.After(10 * time.Second)
			fmt.Printf("Thanks, bye! %v\n", t)
			end = true
		case v1 := <-alice:
			fmt.Println("Alice says", v1)
		case v2 := <-bob:
			fmt.Println("Bob says", v2)
		default:
			time.Sleep(1 * time.Second)
		}
	}
}
