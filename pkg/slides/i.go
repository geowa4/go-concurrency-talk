package slides

import (
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// I More complex use of context with channels
func I() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	go func() {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Press Enter if the meeting is moved up!\n\n")
		_, _ = reader.ReadString('\n')
		fmt.Println("Canceling!")
		cancel()
	}()

	fmt.Println("Asking Alice for the team's progress to bring to the shareholder's meeting")
	statusReportChannel := aliceI(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("\nStatus report was not delivered!")
	case sr := <-statusReportChannel:
		fmt.Printf("\n%s\n", sr)
		fmt.Println("Give Alice a raise!")
	}
}

func aliceI(ctx context.Context) <-chan string {
	reportChannel := make(chan string)
	statusReport := "Team progress:\n"
	go func() {
		defer close(reportChannel)

		deadline, ok := ctx.Deadline()
		if ok {
			fmt.Println("Alice feels the pressure of the deadline set to", deadline)
		}
		timeLeft := deadline.Sub(time.Now())

		reportsNeeded := 2

		bobsContext, cancelBob := context.WithDeadline(ctx, time.Now().Add(timeLeft/3*2))
		defer cancelBob()
		bobsReportChannel := workerI(bobsContext)

		carolsContext, cancelCarol := context.WithDeadline(ctx, time.Now().Add(timeLeft/4*3))
		defer cancelCarol()
		carolsReportChannel := workerI(carolsContext)

		for reportsNeeded > 0 {
			fmt.Println("Waiting for a progress report")
			select {
			case <-ctx.Done():
				fmt.Println("The data for the report did not come in quickly enough!")
				cancelBob()
				cancelCarol()
				return
			case bobsProgress := <-bobsReportChannel:
				reportsNeeded--
				fmt.Println("Received progress update from Bob")
				statusReport += fmt.Sprintf("Bob: %d%%\n", bobsProgress)
			case carolsProgress := <-carolsReportChannel:
				reportsNeeded--
				fmt.Println("Received progress update from Carol")
				statusReport += fmt.Sprintf("Carol: %d%%\n", carolsProgress)
			}
		}
		reportChannel <- statusReport
	}()
	return reportChannel
}

func workerI(ctx context.Context) <-chan int {
	percentComplete := make(chan int)
	go func() {
		select {
		case <-ctx.Done():
			defer close(percentComplete)
			fmt.Println("Worker failed to submit progress in time")
		case <-time.After(time.Duration(rand.Intn(10000)) * time.Millisecond):
			percentComplete <- rand.Intn(101)
			fmt.Println("Worker submitted progress")
		}
	}()
	return percentComplete
}
