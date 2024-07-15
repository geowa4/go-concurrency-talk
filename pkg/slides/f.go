package slides

import (
	"fmt"
	"math/rand"
	"time"
)

func F() {
	c := make(chan string)
	go multiplex(c)
	<-time.After(10 * time.Second)
	close(c)
}

func multiplex(fanIn chan string) {
	alice := sourceF("Alice")
	bob := sourceF("Bob")

	go func() {
		for {
			fanIn <- <-alice
		}
	}()
	go func() {
		for {
			fanIn <- <-bob
		}
	}()
	for msg := range fanIn {
		fmt.Println(msg)
	}
}

func sourceF(name string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			sleepDuration := time.Duration(rand.Intn(3000)) * time.Millisecond
			//fmt.Println(fmt.Sprintf("%s is sleeping for %d", name, sleepDuration))
			time.Sleep(sleepDuration)
			c <- fmt.Sprintf("%s says %d", name, rand.Intn(1000))
		}
	}()
	return c
}
