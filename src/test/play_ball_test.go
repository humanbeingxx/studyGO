package test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

var courtWaitGroup sync.WaitGroup

//TestPlayBall test for no buffer channel
func TestPlayBall(t *testing.T) {
	court := make(chan int)
	courtWaitGroup.Add(2)

	go player("TOM", court)
	go player("JIM", court)

	court <- 1

	courtWaitGroup.Wait()
}

func player(name string, court chan int) {

	defer courtWaitGroup.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}

		if rand.Intn(100)%13 == 0 {
			fmt.Printf("Player %s missed\n", name)
			close(court)
			return
		}

		time.Sleep(500 * time.Millisecond)

		fmt.Printf("Player %s hit %d\n", name, ball)
		ball++

		court <- ball
	}
}
