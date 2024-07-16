package slides

import (
	"fmt"
	"time"
)

func A() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
