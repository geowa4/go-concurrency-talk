package slides

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func D() {
	c := make(chan string)
	wg := &sync.WaitGroup{}
	workerCount := 10
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go DDoSomething(c, wg, i)
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	for s := range c {
		fmt.Println(s)
	}
}

func DDoSomething(c chan string, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	duration := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(duration)
	c <- fmt.Sprintf("%d slept for %d", i, duration)
}
