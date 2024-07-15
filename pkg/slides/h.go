package slides

import (
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// H Introducing context
func H() {
	alice := sourceH()
	defer close(alice)
	bob := sourceH()
	defer close(bob)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	go func() {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Press Enter to cancel")
		_, _ = reader.ReadString('\n')
		fmt.Println("Canceling!")
		cancel()
	}()
	sinkH(ctx, alice, bob)
}

func sourceH() chan int {
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

func sinkH(ctx context.Context, alice, bob <-chan int) {
	end := false

	for !end {
		select {
		case <-ctx.Done():
			fmt.Printf("Thanks, bye! %v\n", ctx.Err())
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
