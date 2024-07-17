package slides

import (
	"fmt"
	"math/rand"
	"time"
)

func E() {
	request, reply := requestReply()
	for i := 0; i < 5; i++ {
		sleepDuration := time.Duration(rand.Intn(5000)) * time.Millisecond
		fmt.Println("Sleeping for", sleepDuration)
		time.Sleep(sleepDuration)

		request <- rand.Intn(1000)
		fmt.Println("Received random value of", <-reply)
		fmt.Println()
	}
}

func requestReply() (chan<- int, <-chan int) { // not the different syntax for the return types
	request := make(chan int)
	reply := make(chan int)
	go func() {
		for {
			maxRand := <-request
			fmt.Println("Generating random number below", maxRand)
			reply <- rand.Intn(maxRand)
		}
	}()
	return request, reply
}
