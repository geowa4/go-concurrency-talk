package slides

import (
	"fmt"
	"math/rand"
	"time"
)

func C() {
	c := make(chan string)
	workerCount := 10
	for i := 0; i < workerCount; i++ {
		go CDoSomething(c, i)
	}
	for s := range c {
		fmt.Println(s)
	}
	// Does this code work?
}

func CDoSomething(c chan string, i int) {
	duration := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(duration)
	c <- fmt.Sprintf("%d slept for %d", i, duration)
}
