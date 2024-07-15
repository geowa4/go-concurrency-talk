package slides

import (
	"fmt"
	"math/rand"
	"time"
)

func B() {
	for i := 0; i < 10; i++ {
		go bDoSomething(i)
	}
	time.Sleep(10 * time.Second)
}

func bDoSomething(i int) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Println(i)
}
